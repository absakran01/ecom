package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/absakran01/ecom/config"
	"github.com/absakran01/ecom/db"
	"github.com/absakran01/ecom/service/user"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	log.Println(`<-#-#-#-#-#-#-#-#-#---GO ECOM API---#-#-#-#-#-#-#-#-#->`)
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error{
	// setup db
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
	log.Fatal(err)
	}
	db.Driver()

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	store := user.NewStore(s.db)
	userhandler := user.NewHandler(store)
	userhandler.RegisterRoutes(subrouter)
	log.Println("listening on", s.addr)


	return http.ListenAndServe(s.addr, router)
}