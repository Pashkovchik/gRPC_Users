package main

import (
	"gRPC_Users/internal/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialising configs: %s", err.Error())
	}
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("error loading env variables: #{err.Error()}")
	}

	_, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: #{err.Error()}")
	}

	dbCount, _ := strconv.Atoi(viper.GetString("redis.dbcount"))
	_, err := repository.NewRedisDB(repository.ConfigRedis{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: "",
		DB:       dbCount,
	})
	if err != nil {
		log.Fatalf("failed to initialize db: #{err.Error()}")
	}

	_, err := repository.NewClickHouseDB(repository.ConfigClickHouse{
		Host: viper.GetString("clickHouse.host"),
		Port: viper.GetString("clickhouse.port"),
	})
	if err != nil {
		log.Fatalf("failed to initialize ClickHouseDB: #{err.Error()}")
	}

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
