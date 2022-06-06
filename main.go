package main

import (
	"dbworks/config"
	"dbworks/database"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	cfg, err := config.NewConfig().FromFileOrDefault("congif.json")
	if err != nil {
		log.Println(err.Error())
		return
	}

	db, err := database.Connect(cfg.ConnString())
	if err != nil {
		log.Println(err.Error())
		return
	}

	q := `CREATE TABLE IF NOT EXISTS test(id serial);`
	_, err = db.Exec(q)
	if err != nil {
		log.Println(err.Error())
		return
	}

}
