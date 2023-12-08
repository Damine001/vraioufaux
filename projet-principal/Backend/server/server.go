package server

import (
	"flag"
	"fmt"
	"vertoufaux/api"
	"vertoufaux/database"

	"log"
	"os"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[WARNING]", err)
	}
}

// database connection
func DBConnection() (*gorm.DB, error) {

	// create database logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)

	// database url
	url := fmt.Sprintf("host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable")

	return gorm.Open(postgres.Open(url), &gorm.Config{Logger: newLogger})
}

// run database
func RunServer() {

	// database connection
	db, err := DBConnection()
	if err != nil {
		panic(fmt.Sprintf("[WARNING] database connection: %v", err))
	}

	// initialize casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("[WARNING] failed to initialize casbin adapter: %v", err))
	}

	// load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("[WARNING] failed to create casbin enforcer: %v", err))
	}

	// check database migration
	// define a bool flag
	database_flag := flag.Bool("database", false, "Bool variable to create database")
	flag.Parse()

	// just create database and quit
	if *database_flag {
		// auto migrate tables & create root user
		database.AutoMigrateDatabase(db, enforcer)
		return
	}

	// declare api routes
	router := gin.New()

	// create api routes group
	router_api := router.Group("/api")
	{
		// call API routes by adding /api as a prefix
		api.RoutesApis(router_api, db, enforcer)
	}

	// run the server
	router.Run(os.Getenv("APP_PORT"))
}
