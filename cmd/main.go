package main

import (
	"context"
	"gRPC_Users/internal/queue"
	"gRPC_Users/internal/repository"
	"gRPC_Users/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
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

	db, err := repository.NewPostgresDB(repository.Config{
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
	client, err := repository.NewRedisDB(repository.ConfigRedis{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: "",
		DB:       dbCount,
	})
	if err != nil {
		log.Fatalf("failed to initialize db: #{err.Error()}")
	}

	con, err := repository.NewClickHouseDB(repository.ConfigClickHouse{
		Host: viper.GetString("clickHouse.host"),
		Port: viper.GetString("clickhouse.port"),
	})
	if err != nil {
		log.Fatalf("failed to initialize ClickHouseDB: #{err.Error()}")
	}

	repos := repository.NewRepository(db)
	reposRedis := repository.NewRedisRepository(client)
	services := service.NewService(repos, reposRedis)
	// доделать handlers :=

	reposClickHouse := repository.NewClickHouseRepository(con)
	consumer := queue.NewConsumer(reposClickHouse)
	go consumer.LogCreateUsr(context.Background())

	listener, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatalf("could not listen on port")
	}

	grpcServer := grpc.NewServer()

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
