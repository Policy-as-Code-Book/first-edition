package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"sync"
	"time"
)

const (
	DefaultLogFile string = "server.log"
)

type LogOptions struct {
	OutputFile string
	Level      zerolog.Level
}

var once sync.Once

var Logger zerolog.Logger

func InitLogger(lo LogOptions) {

	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		fileLogger := &lumberjack.Logger{
			Filename:   lo.OutputFile,
			MaxSize:    5, //
			MaxBackups: 10,
			MaxAge:     14,
			Compress:   true,
		}

		output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)

		//var gitRevision string
		//
		//buildInfo, ok := debug.ReadBuildInfo()
		//if ok {
		//	for _, v := range buildInfo.Settings {
		//		if v.Key == "vcs.revision" {
		//			gitRevision = v.Value
		//			break
		//		}
		//	}
		//}

		Logger = zerolog.New(output).
			Level(lo.Level).
			With().
			Timestamp().
			//Str("git_revision", gitRevision).
			//Str("go_version", buildInfo.GoVersion).
			Logger()
	})
}
