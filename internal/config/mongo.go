package config

import "github.com/spf13/viper"

func GetMongoConnectionURL() string {
	return viper.GetString("mongo_connection_url")
}

func GetMongoDatabaseName() string {
	return viper.GetString("mongo_database_name")
}
