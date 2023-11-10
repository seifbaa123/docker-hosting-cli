package subcommands

import (
	"docker-hosting-cli/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Message string `json:"message"`
	Image   struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		HasBuild bool   `json:"hasBuild"`
	} `json:"image"`
}

func Create(projectName string) {
	multipartWriter, requestBody := utils.CreateFormData(projectName)
	req, err := http.NewRequest("POST", apiUrl, &requestBody)
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

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	config := strings.Join([]string{
		fmt.Sprintf("id: %d", response.Image.Id),
		fmt.Sprintf("name: %s", response.Image.Name),
	}, "\n")

	content := []byte(config)
	err = os.WriteFile(configFile, content, 0644)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("Done")
}
