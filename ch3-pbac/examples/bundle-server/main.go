package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"jimmyray.io/opa-bundle-api/pkg/model"
	"jimmyray.io/opa-bundle-api/pkg/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultConfigFile string = "server-config.json"
)

var (
	flags map[string]string
)

func initServer() {
	flags = make(map[string]string)
	flagConfigFile := "Server config file"

	var serverConfigFile string

	flag.StringVar(&serverConfigFile, "config", defaultConfigFile, flagConfigFile)
	flag.Parse()

	err := model.LoadConfig(serverConfigFile)
	if err != nil {
		utils.Logger.Error().Err(err).Msg("error loading server config")
		panic(err)
	} else {
		utils.Logger.Info().Msg("server successfully configured")
	}

	model.ServiceInfo.NAME = model.SC.Metadata.Name
	model.ServiceInfo.ID = model.GetServiceId()

	var level zerolog.Level
	switch model.SC.Init.LogLevel {
	case "debug":
		level = zerolog.DebugLevel
	case "error":
		level = zerolog.ErrorLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "warn":
		level = zerolog.WarnLevel
	default:
		level = zerolog.InfoLevel
	}

	// Init logger
	utils.InitLogger(utils.LogOptions{
		OutputFile: model.SC.Init.LogFile,
		Level:      level,
	})

	utils.Logger.Info().Msg("Service started successfully.")
	utils.Logger.Info().Msgf("Flags: %+v, Args: %+v", flags, os.Args)

	model.IC = model.InfoController{
		ServiceInfo: model.ServiceInfo,
	}

	if model.SC.Bundles.Enable {
		err = model.BuildBundle()
		if err != nil {
			utils.Logger.Error().Err(err).Msg("build bundles failure")
		} else {
			utils.Logger.Debug().Msg("bundles processed")
		}
	} else {
		utils.Logger.Info().Msg("server started with no bundles")
	}

	// Registered bundles
	reg, _ := model.RegBundles.Json()
	utils.Logger.Debug().Msgf("Registered bundles: %s", string(reg))
}

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return r
}

func main() {
	initServer()
	utils.Logger.Info().Msgf("Listening on socket %s:%s", model.SC.Network.ServerAddress, model.SC.Network.ServerPort)
	router := Router()

	// Middleware
	if model.SC.AuthZ.Enable {
		err := model.EnableAuth()
		if err != nil {
			utils.Logger.Error().Err(err).Msg("could not enabled authz")
			panic(err)
		}
		utils.Logger.Debug().Msg("AuthZ middleware enabled")
		router.Use(model.AuthZMiddleware)
	} else {
		utils.Logger.Debug().Msg("AuthZ middleware disabled")
	}
	
	if utils.Logger.GetLevel().String() == zerolog.DebugLevel.String() {
		utils.Logger.Debug().Msg("Request logging middleware enabled")
		router.Use(model.LoggingMiddleware)
	}

	if model.SC.Init.EnableEtag {
		utils.Logger.Debug().Msg("ETag middleware enabled")
		router.Use(model.EtagMiddleware)
	} else {
		utils.Logger.Debug().Msg("ETag middleware disabled")
	}

	if !model.SC.Init.AllowDirList {
		utils.Logger.Debug().Msg("Directory listing prevention middleware enabled")
		router.Use(model.NoListMiddleware)
	} else {
		utils.Logger.Debug().Msg("Directory listing prevention middleware disabled")
	}

	// Handlers
	router.HandleFunc(model.SC.Network.Uris.Health, model.IC.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc(model.SC.Network.Uris.Info, model.IC.GetServiceInfo).Methods(http.MethodGet)
	router.HandleFunc(model.SC.Network.Uris.Api+"/bundles", model.C.GetBundles).Methods(http.MethodGet)
	router.HandleFunc(model.SC.Network.Uris.Api+"/server/config", model.C.GetServerConfig).Methods(http.MethodGet)
	fs := http.FileServer(http.Dir(model.SC.Bundles.BundleOutDir))
	//fs := http.FileServer(http.Dir(model.SC.Bundles.BundleOutDir))
	router.PathPrefix(model.SC.Bundles.BundleUri).Handler(http.StripPrefix(model.SC.Bundles.BundleUri, fs))

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// TLS?
		if model.SC.Network.TLS.Enable {
			fmt.Println(http.ListenAndServeTLS(model.SC.Network.ServerAddress+":"+model.SC.Network.ServerPort, model.SC.Network.TLS.Cert, model.SC.Network.TLS.Key, router))
		} else {
			fmt.Println(http.ListenAndServe(model.SC.Network.ServerAddress+":"+model.SC.Network.ServerPort, router))
		}
	}()

	<-done
	utils.Logger.Info().Msg("server stopping...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		if model.SC.Bundles.PurgeBundles {
			// Purge bundles and registry
			model.PurgeBundles()
		}
		cancel()
	}()

	utils.Logger.Info().Msg("Server Exited Gracefully")
}
