package core

import (
	"context"
	"fmt"
	"gitee.com/favefan/doconsole/initialize"
	"gitee.com/favefan/doconsole/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", initialize.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    initialize.ServerSetting.ReadTimeout,
		WriteTimeout:   initialize.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// s.ListenAndServe()
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
