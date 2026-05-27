package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode string
	Timezone string
}


func (server *Server) Initialize(appConfig AppConfig, dBConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName + " API")

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",dBConfig.Host, dBConfig.User, dBConfig.Password, dBConfig.DBName, dBConfig.Port, dBConfig.SSLMode, dBConfig.Timezone)	
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database:" + err.Error())
	}

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Server is running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func Run() {
	server := Server{}
	appConfig := AppConfig{}
	dBConfig := DBConfig{}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoShop")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9999")

	dBConfig.Host = getEnv("DB_HOST", "localhost")
	dBConfig.User = getEnv("DB_USER", "root")
	dBConfig.Password = getEnv("DB_PASSWORD", "root")
	dBConfig.DBName = getEnv("DB_NAME", "db_goshop")
	dBConfig.Port = getEnv("DB_PORT", "5432")
	dBConfig.SSLMode = getEnv("DB_SSLMODE", "disable")
	dBConfig.Timezone = getEnv("DB_TIMEZONE", "Asia/Jakarta")

	server.Initialize(appConfig, dBConfig)
	server.Run(":" + appConfig.AppPort)
}
