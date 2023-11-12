package utils

import (
	"docker-hosting-cli/config"
	"docker-hosting-cli/logs"
	"docker-hosting-cli/types"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadProjectConfig() types.ProjectConfig {
	logs.Info("reading config file")
	yamlFile, err := os.ReadFile(config.ConfigFile)
	if err != nil {
		logs.Error("failed on reading config file: %s", err.Error())
	}

	logs.Info("parsing config file")
	var projectConfig types.ProjectConfig
	err = yaml.Unmarshal(yamlFile, &projectConfig)
	if err != nil {
		logs.Error("failed on parsing config file: %s", err.Error())
	}

	return projectConfig
}
