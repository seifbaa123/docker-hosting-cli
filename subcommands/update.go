package subcommands

import (
	"docker-hosting-cli/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Id   int    `yaml:"id"`
	Name string `yaml:"name"`
}

func Update() {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	url := fmt.Sprintf("%s/api/images/%d", apiUrl, config.Id)

	multipartWriter, requestBody := utils.CreateFormData(config.Name)
	req, err := http.NewRequest("PUT", url, &requestBody)
	if err != nil {
		log.Fatal("Error creating POST request:", err)
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending POST request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	fmt.Println(string(body))
}
