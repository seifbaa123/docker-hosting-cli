package subcommands

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func Delete() {
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

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal("Error creating DELETE request:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending DELETE request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	err = os.Remove(configFile)
	if err != nil {
		log.Fatal("Error deleting config file:", err)
	}

	fmt.Println(string(body))
}
