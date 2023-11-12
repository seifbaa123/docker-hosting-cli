package subcommands

import (
	"docker-hosting-cli/config"
	"docker-hosting-cli/logs"
	"docker-hosting-cli/utils"
	"fmt"
	"net/http"
)

func Update() {
	projectConfig := utils.ReadProjectConfig()
	url := fmt.Sprintf("%s/api/images/%d", config.ApiUrl, projectConfig.Id)

	logs.Info("prepare from data")
	multipartWriter, requestBody := utils.CreateFormData(projectConfig.Name)

	logs.Info("send put request to %s", url)
	req, err := http.NewRequest("PUT", url, &requestBody)
	if err != nil {
		logs.Error("failed on creating PUT request: %s", err.Error())
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response := utils.ParseResponseBody(req)
	if response.Message != "success" {
		logs.Error(response.Message)
	}

	logs.Info("done")
}
