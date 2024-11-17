package api

import (
	"database/sql"
	"log"
	"net/http"
	"rest/service/user"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	address string
	db      *sql.DB
}

func NewApiServer(address string, db *sql.DB) *ApiServer {
	return &ApiServer{
		address: address,
		db:      db,
	}
}

func Run(s *ApiServer) error {

	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// user related routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Server is running on ", s.address)
	return http.ListenAndServe(s.address, router)
}
