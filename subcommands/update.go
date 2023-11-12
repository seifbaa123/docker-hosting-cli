package subcommands

import (
	"docker-hosting-cli/logs"
	"docker-hosting-cli/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func Update() {
	logs.Info("reading config file")
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		logs.Error("failed on reading config file: %s", err.Error())
	}

	logs.Info("parsing config file")
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		logs.Error("failed on parsing config file: %s", err.Error())
	}

	url := fmt.Sprintf("%s/api/images/%d", apiUrl, config.Id)

	logs.Info("prepare from data")
	multipartWriter, requestBody := utils.CreateFormData(config.Name)

	logs.Info("send put request to %s", url)
	req, err := http.NewRequest("PUT", url, &requestBody)
	if err != nil {
		logs.Error("failed on creating PUT request: %s", err.Error())
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("failed on sending PUT request: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logs.Error("Request failed with status code: %d", resp.StatusCode)
	}

	logs.Info("reading response body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error("failed on reading response body: %s", err.Error())
	}

	logs.Info("decoding JSON")
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		logs.Error("failed on decoding JSON: %s", err.Error())
	}

	if response.Message != "success" {
		logs.Error(response.Message)
	}

	logs.Info("done")
}
