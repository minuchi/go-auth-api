package lib

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabaseConfigString(databaseConfig DatabaseConfig) string {
	databaseConfigMap := map[string]string{
		"host":     databaseConfig.Host,
		"user":     databaseConfig.User,
		"password": databaseConfig.Password,
		"dbname":   databaseConfig.DbName,
		"port":     databaseConfig.Port,
		"sslmode":  databaseConfig.Sslmode,
		"TimeZone": databaseConfig.Timezone,
	}

	var dns string
	for key, value := range databaseConfigMap {
		dns += fmt.Sprintf("%s=%s ", key, value)
	}
	return strings.TrimSpace(dns)
}

func ConnectToDB(databaseConfig DatabaseConfig) *gorm.DB {
	dsn := getDatabaseConfigString(databaseConfig)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
