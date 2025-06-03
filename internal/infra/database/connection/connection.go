package connection

import (
	"log"
	"strconv"

	"github.com/Gabriel-Schiestl/api-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConfig *config.DbConfig
var Db *gorm.DB

func SetupConfig(host, user, password, port, name string) {
	dbPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Error converting DB_PORT to int")
	}

	DbConfig = config.NewDbConfig(host, user, password, name, dbPort)

	Db, err = gorm.Open(postgres.Open(DbConfig.ToString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	sqlDb, err := Db.DB()
	if err != nil {
		log.Fatalf("Error getting DB connection: %v", err)
	}

	defer sqlDb.Close()
}