package main

import (
	"fmt"
	"net/http"
	"time"
	"y-u-y-a/template-go/config"
	"y-u-y-a/template-go/handler"
	"y-u-y-a/template-go/usecase"
	"y-u-y-a/template-go/utils"

	chi "github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var location = "Asia/Tokyo"

func init() {
	// タイムゾーンの設定
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
	utils.Clock = utils.NewRealTimeClock()
}

func main() {
	/*****************************
	ロギングの出力設定
	******************************/
	var logger *zap.Logger
	if config.AppEnv.IsProd() {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	/*****************************
	DB接続
	******************************/
	// db_url := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 	config.AppEnv.DbUser, config.AppEnv.DbPassword,
	// 	config.AppEnv.DbHost, config.AppEnv.DbPort,
	// 	config.AppEnv.DbName,
	// )
	// client, err := ent.Open("postgres", db_url)
	// if err != nil {
	// 	logger.Fatal("データベース接続に失敗しました", zap.Error(err))
	// }
	// defer client.Close()

	/*****************************
	ミドルウェアの設定
	******************************/
	r := chi.NewRouter()
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(120 * time.Second))
	r.Use(chiMiddleware.Logger)
	// CORSの設定
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Set-Cookie"},
		MaxAge:           3600,
	}))

	/*****************************
	ルーティングの設定
	******************************/
	userUsecase := usecase.NewUserUsecase(logger)
	publicHandler := handler.NewPublicHandler(logger, userUsecase)
	r.Route("/", func(r chi.Router) {
		r.Get("/user", publicHandler.Route().ServeHTTP)
	})

	/*****************************
	サーバー起動
	******************************/
	pp.Println("appConfig is", config.AppEnv)
	logger.Info(fmt.Sprintf("Start listening on %v", config.AppEnv.ServerPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.AppEnv.ServerPort), r); err != nil {
		logger.Fatal("サーバー起動に失敗しました", zap.Error(err))
	}
}
