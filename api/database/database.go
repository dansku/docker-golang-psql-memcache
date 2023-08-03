package database

import (
	"database/sql"
	"os"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bradfitz/gomemcache/memcache" // memcached
	"github.com/joho/godotenv"                // package used to read the .env file
	_ "github.com/lib/pq"                     // postgres golang driver
)

// Global variables
var (
	DB *sql.DB
	MC *memcache.Client
)

// Start database
func Start() {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	// Postgres ---------------------------------------------------

	DB, err = openPostgresConnection(os.Getenv("POSTGRES_CONNECTION"))
	handleSQLError(err, "Error opening the database - ETL")
	log.Info().Msg("Database Successfully connected!!!")

	// ------------------------------------------------------------

	MC = memcache.New(os.Getenv("MEMCACHE_CONNECTION"))

	err = MC.Ping()
	handleSQLError(err, "Error opening the memcached")
	log.Info().Msg("Memcached Successfully connected!!!")
}

func openPostgresConnection(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(2000)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
func handleSQLError(err error, msg string) {
	if err != nil {
		log.Error().Err(err).Msg(msg)

		if err.Error() == "sql: no rows in result set" {
			return
		}

		if err.Error() == "pq: sorry, too many clients already" {
			os.Exit(1)
		}
	}
}
