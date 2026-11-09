package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config Error")
	}

	client, _, err := db.Connect(cfg)
	if err != nil {
		log.Fatal("DB Error")
	}

	defer func() {
		if err := db.DisConnect(client); err != nil {
			log.Println("Mongo DisConnect Error:", err)
		}
	}()

	router := server.NewRouter()

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatal("Server Failed")
	}
}
