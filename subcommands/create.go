package subcommands

import (
	"docker-hosting-cli/logs"
	"docker-hosting-cli/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Create(projectName string) {
	url := fmt.Sprintf("%s/api/images", apiUrl)

	logs.Info("prepare from data")
	multipartWriter, requestBody := utils.CreateFormData(projectName)

	logs.Info("send post request to %s", url)
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		logs.Error("failed on creating POST request: %s", err.Error())
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("failed on sending POST request: %s", err.Error())
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

	config := strings.Join([]string{
		fmt.Sprintf("id: %d", response.Image.Id),
		fmt.Sprintf("name: %s", response.Image.Name),
	}, "\n")

	logs.Info("writing to config file")
	content := []byte(config)
	err = os.WriteFile(configFile, content, 0644)
	if err != nil {
		logs.Error("failed on writing to config file %s: %s", configFile, err.Error())
	}

	logs.Info("done")
}
