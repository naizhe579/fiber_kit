package env

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

func InitEnvFile(filePath string) {
	log.Info().Msg("InitEnvFile: <" + filePath + "> 配置文件加载中...")
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatal().Err(err).Msg("InitEnvFile: <" + filePath + "> 配置文件加载失败")
	} else {
		log.Info().Msg("InitEnvFile: <" + filePath + "> 配置文件加载成功")
	}
}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
