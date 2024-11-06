package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"cms/v2/internal/api/operate"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
	db  *gorm.DB
	rdb *redis.Client
	operateAppClient operate.AppClient
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connDB(app)
	connRdb(app)
	connOperateAppClient(app)

	return app
}

func connDB(app *CmsApp) {
	dbUrl := os.Getenv("DB_URL")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbUrl)
	mysqlDB, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	app.db = mysqlDB
}

func connRdb(app *CmsApp) {
	redisDBStr := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDBStr)

	if err != nil {
		log.Fatalf("Error converting REDIS_DB to int: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		panic(err)
	}

	app.rdb = rdb
}

func connOperateAppClient(app *CmsApp) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)

	if err != nil {
		panic(err)
	}

	appClient := operate.NewAppClient(conn)
	app.operateAppClient = appClient
}