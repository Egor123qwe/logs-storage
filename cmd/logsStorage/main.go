package main

import (
	"context"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/viper"

	"github.com/Egor123qwe/logs-storage/internal/app"
	exit "github.com/Egor123qwe/logs-storage/internal/util/context"
	"github.com/Egor123qwe/logs-storage/internal/util/logger"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	logger.Init()
}

func main() {
	srv, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := exit.WithSignal(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := srv.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
