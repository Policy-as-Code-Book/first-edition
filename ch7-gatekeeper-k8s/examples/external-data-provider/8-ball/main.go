package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"jimmyray.io/magic-8-ball/pkg/model"
	"jimmyray.io/magic-8-ball/pkg/utils"
)

func main() {
	flagConfigFile := "kubectl config file path (string)"
	flag.StringVar(&model.ConfigFile, "file", "", flagConfigFile)
	flag.StringVar(&model.ConfigFile, "f", "", flagConfigFile)
	flag.Parse()

	fmt.Println("Config file path: ", model.ConfigFile)

	if model.ConfigFile == "" {
		panic("Config file path not specified")
	}

	e := model.AC.LoadConfig(model.ConfigFile)
	if e != nil {
		panic(fmt.Errorf("error ingesting config file: %w", e))
	}

	fmt.Println(fmt.Sprintf("config file (%s) ingested successfully", model.ConfigFile))
	b, err := model.AC.Yaml()
	if err != nil {
		panic(fmt.Errorf("error reading config: %w", e))
	}

	fmt.Println(string(b))

	if model.AC.Network.TLS.Enabled {

		key := model.AC.Network.TLS.Key
		cert := model.AC.Network.TLS.Cert
		tf := utils.FileExists(key)
		if tf {
			fmt.Println("TLS Key file exists: ", key)
		} else {
			fmt.Println("TLS Key file does not exist: ", key)
		}
		tf = utils.FileExists(cert)
		if tf {
			fmt.Println("TLS Cert file exists: ", cert)
		} else {
			fmt.Println("TLS Cert file does not exist: ", cert)
		}
	}

	fmt.Println("Server Address: " + model.AC.Network.ServerAddress)
	fmt.Println("Server Port: " + model.AC.Network.ServerPort)

	router := Router()
	router.HandleFunc("/healthz", model.C.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/answer", model.C.Answer).Methods(http.MethodGet)
	router.HandleFunc("/provide", model.C.Provide).Methods(http.MethodPost)

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if model.AC.Network.TLS.Enabled {
			fmt.Println(http.ListenAndServeTLS(model.AC.Network.ServerAddress+":"+model.AC.Network.ServerPort, model.AC.Network.TLS.Cert, model.AC.Network.TLS.Key, router))
		} else {
			fmt.Println(http.ListenAndServe(model.AC.Network.ServerAddress+":"+model.AC.Network.ServerPort, router))
		}
	}()

	<-done
	fmt.Println("server stopping...")

	//_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer func() {
	//	// cleanup
	//	cancel()
	//}()
	//
	//fmt.Println("server exited gracefully")
}

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return r
}
