package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"surelink-go/server"
	db "surelink-go/sqlc"
	"surelink-go/util"
)

func main() {
	globalConfig, err := util.LoadGlobalConfig(".")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	conn, err := sql.Open(globalConfig.DBDriver, globalConfig.DBSource)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}

	store := db.NewStore(conn)

	serverObj := server.NewServer(store)
	err = serverObj.Start(globalConfig.ServerAddress)
	if err != nil {
		log.Fatal("can't start the server", err)
	}

}
