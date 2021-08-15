package cmd

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Joker666/spotangels-demo/web"
	"github.com/Joker666/spotangels-demo/web/controller"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// srvCmd is the serve sub command to start the api server
var srvCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve serves the api server",
	RunE:  serve,
}

func init() {
	//
}

func serve(cmd *cobra.Command, args []string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connectDatabase()
	if err != nil {
		return err
	}
	log.Println(db.Name())

	regulationController := controller.NewRegulationController()

	port := ":" + os.Getenv("PORT")
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	web.NewServer(s, regulationController)

	log.Println("Starting server on", port)
	srv := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func connectDatabase() (*gorm.DB, error) {
	sqlDB, err := sql.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		return nil, err
	}

	i := 0
	for {
		pingErr := sqlDB.Ping()
		log.Println("Pinging DB")
		if pingErr != nil {
			log.Println(pingErr)
			if i > 10 {
				log.Println("Could not connect to database, exiting")
				os.Exit(1)
			}
		} else {
			log.Println("Database Connected")
			break
		}
		time.Sleep(time.Duration(1) * time.Second)
		i = i + 1
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
