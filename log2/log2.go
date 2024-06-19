package log2

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func InitLogConfig() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.NewConsoleWriter(
		func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = time.DateTime
		}))
	log.Info().Msg("InitLogConfig:日志初始化成功")
}
