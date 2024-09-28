package psql

import (
	"fmt"

	"github.com/spf13/viper"
)

type dbConfig struct {
	URL    string
	Driver string
}

type Config struct {
	db dbConfig
}

func NewConfig() Config {
	config := Config{
		db: dbConfig{
			URL:    newDBParams("db.logs_storage").ParseURL(),
			Driver: viper.GetString("db.logs_storage.driver"),
		},
	}

	return config
}

type dbParams struct {
	configPath string

	host     string
	port     string
	dbname   string
	user     string
	password string
}

func newDBParams(configPath string) dbParams {
	return dbParams{
		configPath: configPath,

		host:     viper.GetString(fmt.Sprintf("%s.host", configPath)),
		port:     viper.GetString(fmt.Sprintf("%s.port", configPath)),
		dbname:   viper.GetString(fmt.Sprintf("%s.dbname", configPath)),
		user:     viper.GetString(fmt.Sprintf("%s.user", configPath)),
		password: viper.GetString(fmt.Sprintf("%s.password", configPath)),
	}
}

func (db dbParams) ParseURL() string {
	template := viper.GetString(fmt.Sprintf("%s.urlTemplate", db.configPath))

	return fmt.Sprintf(template, db.host, db.port, db.dbname, db.user, db.password)
}
