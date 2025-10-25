package main

import (
	"log"
	"social-network/internal/db"
	"social-network/internal/env"
	"social-network/internal/store"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
    log.Fatal("Error loading .env file")
  }

   cfg :=  config{
		addr: env.GetString("ADDR" ,":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR","postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS",25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS",25),
			maxIdleTime: env.GetString( "DB_MAX_IDLE_TIME","15m"),
		},
	}
	db, err := db.New(
	cfg.db.addr,
	cfg.db.maxOpenConns,
	cfg.db.maxIdleConns,
	cfg.db.maxIdleTime,
)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("Database Connection pool established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
 }	
 mux := app.mount()

 log.Fatal(app.run(mux))
}