package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "viper-pzn-go", config.GetString("app.name"))
	assert.Equal(t, "Luthfi Izzuddin Hanif", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestYAML(t *testing.T) {
	config := viper.New()
	//config.SetConfigName("config")
	//config.SetConfigType("yaml")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	// read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "viper-pzn-go", config.GetString("app.name"))
	assert.Equal(t, "Luthfi Izzuddin Hanif", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestENVFile(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	// read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "viper-pzn-go", config.GetString("APP_NAME"))
	assert.Equal(t, "Luthfi Izzuddin Hanif", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
}

func TestENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv() // baca environment var dari OS

	// read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "viper-pzn-go", config.GetString("APP_NAME"))
	assert.Equal(t, "Luthfi Izzuddin Hanif", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))

	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
