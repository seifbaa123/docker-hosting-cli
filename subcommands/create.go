package subcommands

import (
	"docker-hosting-cli/config"
	"docker-hosting-cli/logs"
	"docker-hosting-cli/utils"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Create(projectName string) {
	url := fmt.Sprintf("%s/api/images", config.ApiUrl)

	logs.Info("prepare from data")
	multipartWriter, requestBody := utils.CreateFormData(projectName)

	logs.Info("send post request to %s", url)
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		logs.Error("failed on creating POST request: %s", err.Error())
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response := utils.ParseResponseBody(req)
	projectConfig := strings.Join([]string{
		fmt.Sprintf("id: %d", response.Image.Id),
		fmt.Sprintf("name: %s", response.Image.Name),
	}, "\n")

	logs.Info("writing to config file")
	content := []byte(projectConfig)
	err = os.WriteFile(config.ConfigFile, content, 0644)
	if err != nil {
		logs.Error("failed on writing to config file %s: %s", config.ConfigFile, err.Error())
	}

	logs.Info("done")
}
