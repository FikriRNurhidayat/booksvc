package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetLocalStorageDirectory() string {
	return viper.GetString("local_storage_directory")
}

func GetPublicURL() string {
	url := viper.GetString("public_url")
	if url != "" {
		return url
	}

	port := viper.GetInt("port")

	return fmt.Sprintf("http://localhost:%d", port)
}
