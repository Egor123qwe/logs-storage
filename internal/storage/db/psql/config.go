package psql

import (
	"fmt"

	"github.com/spf13/viper"
)

type DB struct {
	URL    string
	Driver string
}

type Config struct {
	logStorage DB
}

func NewConfig() Config {
	config := Config{
		logStorage: DB{
			URL:    parseDbURL("db.logs_storage"),
			Driver: viper.GetString("db.logs_storage.driver"),
		},
	}

	return config
}

func parseDbURL(dbConfigName string) string {
	host := viper.GetString(fmt.Sprintf("%s.host", dbConfigName))
	port := viper.GetString(fmt.Sprintf("%s.port", dbConfigName))
	dbname := viper.GetString(fmt.Sprintf("%s.dbname", dbConfigName))
	user := viper.GetString(fmt.Sprintf("%s.user", dbConfigName))
	password := viper.GetString(fmt.Sprintf("%s.password", dbConfigName))

	urlTemplate := "host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable"

	return fmt.Sprintf(urlTemplate, host, port, dbname, user, password)
}
