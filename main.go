package main

import (
	"Butterfly/config"
	db "Butterfly/db/sqlc"
	"Butterfly/handler"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

func main() {
	conf, err := config.LoadConfig("config/.")
	if err != nil {
		log.Fatalf("Can't load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbConn, err := sqlx.Connect(conf.Postgres.Driver, conf.Postgres.DatabaseSource())
	if err != nil {
		log.Fatalf("Can't connect database: %v", err)
	}

	store := db.NewStore(dbConn)

	router := handler.NewHandler(conf, store)

	srv := &http.Server{
		Addr:    conf.Http.Address(),
		Handler: router.Router,
	}

	shutdown(ctx, srv)
}

func shutdown(ctx context.Context, srv *http.Server) {
	// 启动HTTP服务器
	go func() {
		log.Print("Listening and serving HTTP on :" + strings.Split(srv.Addr, ":")[1])
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("Error starting server: ", err)
		}
	}()

	// 用于接收退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	/// 等待退出信号
	<-quit
	log.Print("Received exit signal, shutting down...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}

	log.Println("Server closed...")
}
