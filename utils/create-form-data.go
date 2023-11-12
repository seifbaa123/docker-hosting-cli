package utils

import (
	"bytes"
	"docker-hosting-cli/logs"
	"io"
	"mime/multipart"
	"os"
)

func CreateFormData(projectName string) (*multipart.Writer, bytes.Buffer) {
	dir, err := os.Getwd()
	if err != nil {
		logs.Error("could not get the working directory: %s", err.Error())
	}

	zipFilePath := "/tmp/" + GenerateRandomFilename()
	err = ZipDir(dir, zipFilePath)
	if err != nil {
		logs.Error("could not compress the project in zip file: %s", err.Error())
	}

	file, err := os.Open(zipFilePath)
	if err != nil {
		logs.Error("could not compress open the file %s: %s", zipFilePath, err.Error())
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	nameField, err := writer.CreateFormField("name")
	if err != nil {
		logs.Error("could not create form field: %s", err.Error())
	}
	nameField.Write([]byte(projectName))

	fileField, err := writer.CreateFormFile("file", "file.zip")
	if err != nil {
		logs.Error("could not create form file %s:", err.Error())
	}

	_, err = io.Copy(fileField, file)
	if err != nil {
		logs.Error("could not copy file data to form field: %s", err.Error())
	}

	writer.Close()

	return writer, body
}
