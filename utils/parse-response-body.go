package utils

import (
	"docker-hosting-cli/logs"
	"docker-hosting-cli/types"
	"encoding/json"
	"io"
	"net/http"
)

func ParseResponseBody(req *http.Request) types.Response {
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
	var response types.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		logs.Error("failed on decoding JSON: %s", err.Error())
	}

	return response
}
