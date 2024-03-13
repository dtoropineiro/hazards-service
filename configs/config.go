package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("Error al leer el archivo de configuración: %s", err)
	}
}
