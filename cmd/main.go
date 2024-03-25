package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/edgarSucre/jw/adapters/db"
	"github.com/edgarSucre/jw/adapters/web"
	"github.com/edgarSucre/jw/features/auth"
	"github.com/edgarSucre/jw/features/user"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// jupyterClient := japi.NewClient(
	// 	"http://localhost:9091/hub/api",
	// 	"8e7f67a80f7c23694786d353257c33a1f79d35777fec3c96f9f5d625bbf202a5",
	// )

	if err := loadEnv(); err != nil {
		panic(err)
	}

	dbConn, err := sql.Open("sqlite3", "db/jupyter.db")
	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewSqliteRepository(dbConn)

	sm := web.NewSessionManager()
	navigator := web.Navigator{}
	renderer := web.ViewRenderer{}

	authUC := auth.NewUseCase(repository)
	userUC := user.NewUseCase(repository)

	authHandler := auth.NewHandler(navigator, renderer, sm, authUC)
	userHandler := user.NewHandler(renderer, userUC)

	server := web.NewServer(*userHandler, *authHandler)
	http.ListenAndServe(":8050", server.Handler)

}

func loadEnv() error {
	if os.Getenv("ENVIRONMENT") == "ci" {
		return nil
	}

	return godotenv.Load(".env")
}
