package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type CheatPath struct {
	Name     string
	Path     string
	Tags     []string
	Readonly bool
}

var Editor string

func ReadConfig() (CheatPath, error) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath("$HOME/dotfiles/go-cheatsheet/")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	// Set the type of the configuration file
	viper.SetConfigType("toml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var paths []CheatPath

	if err := viper.UnmarshalKey("cheatpaths", &paths); err != nil {
		return CheatPath{}, err
	}

	Editor = viper.GetString("main.editor")

  // TODO: Handle sending all Send all 
  return paths[0], nil
}
