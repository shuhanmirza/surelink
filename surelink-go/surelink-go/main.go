package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"surelink-go/goserver"
	gedis "surelink-go/redisStore"
	db "surelink-go/sqlc"
	"surelink-go/util"
)

func main() {

	initialTests()

	globalConfig, err := util.LoadGlobalConfig(".")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	conn, err := sql.Open(globalConfig.DBDriver, globalConfig.DBSource)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}
	store := db.NewStore(conn)

	redisStore := gedis.NewRedisStore(globalConfig.RedisUrl)

	serverObj := goserver.NewServer(store, redisStore)
	err = serverObj.Start(globalConfig.ServerAddress)
	if err != nil {
		log.Fatal("can't start the server", err)
	}

}

func initialTests() {
	if _, err := os.Stat(util.FONT_COMIC_PATH); err != nil {
		panic(err.Error())
	}
}
