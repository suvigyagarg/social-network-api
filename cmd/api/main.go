package main

import (
	"log"
	"social-network/internal/env"
	"social-network/internal/store"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
    log.Fatal("Error loading .env file")
  }

     store := store.NewStorage(nil)
   cfg :=  config{
		addr: env.GetString("ADDR" ,":8080"),
	}
	app := &application{
	config: cfg,
	store: store,
 }	
 mux := app.mount()

 log.Fatal(app.run(mux))
}