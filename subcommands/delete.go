package subcommands

import (
	"docker-hosting-cli/config"
	"docker-hosting-cli/logs"
	"docker-hosting-cli/utils"
	"fmt"
	"net/http"
	"os"
)

func Delete() {
	projectConfig := utils.ReadProjectConfig()
	url := fmt.Sprintf("%s/api/images/%d", config.ApiUrl, projectConfig.Id)

	logs.Info("send delete request to %s", url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		logs.Error("failed on creating DELETE request: %s", err.Error())
	}

	response := utils.ParseResponseBody(req)
	if response.Message != "success" {
		logs.Error(response.Message)
	}

	err = os.Remove(config.ConfigFile)
	if err != nil {
		logs.Error("failed on deleting config file: %s", err.Error())
	}

	logs.Info("done")
}
