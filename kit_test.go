package test

import (
	"github.com/naizhe579/fiber_kit/db"
	"github.com/naizhe579/fiber_kit/env"
	"github.com/naizhe579/fiber_kit/log2"
	"testing"
)

func TestKit(t *testing.T) {
	log2.InitLogConfig()
	env.InitEnvFile("./kit.env")
	dsn := env.GetEnvValue("DSN")
	db.InitDB(dsn)
}
