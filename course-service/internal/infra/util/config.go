package util

import (
	"log"

	"github.com/spf13/viper"
)

func SetupConfig(filePath, fileName, fileType string) {
	// initialize Viper
	viper.SetConfigName(fileName) // Name of the config file without extension
	viper.SetConfigType(fileType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(filePath) // Optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error reading config: %s\n", err.Error())
	}
}
