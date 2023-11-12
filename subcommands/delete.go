package subcommands

import (
	"docker-hosting-cli/logs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func Delete() {
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

	logs.Info("send delete request to %s", url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		logs.Error("failed on creating DELETE request: %s", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("failed on sending DELETE request: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logs.Error("request failed with status code: %d", resp.StatusCode)
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

	err = os.Remove(configFile)
	if err != nil {
		logs.Error("failed on deleting config file: %s", err.Error())
	}

	logs.Info("done")
}
