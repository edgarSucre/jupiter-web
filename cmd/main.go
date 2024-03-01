package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edgarSucre/jw/adapters/db"
	"github.com/edgarSucre/jw/adapters/web"
	"github.com/edgarSucre/jw/features/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	// jupyterClient := japi.NewClient(
	// 	"http://localhost:9091/hub/api",
	// 	"8e7f67a80f7c23694786d353257c33a1f79d35777fec3c96f9f5d625bbf202a5",
	// )

	dbConn, err := sql.Open("sqlite3", "db/jupyter.db")
	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewSqliteRepository(dbConn)

	userHandler := user.NewHandler(repository)
	server := web.NewServer(*userHandler)
	http.ListenAndServe(":8050", server.Handler)

}
