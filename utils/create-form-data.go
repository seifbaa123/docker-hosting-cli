package utils

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func CreateFormData(projectName string) (*multipart.Writer, bytes.Buffer) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error:", err)
	}
	zipFilePath := "/tmp/" + GenerateRandomFilename()

	err = ZipDir(dir, zipFilePath)
	if err != nil {
		log.Fatal("Could not compress the project in zip file:", err)
	}

	file, err := os.Open(zipFilePath)
	if err != nil {
		log.Fatal("Error opening the file:", err)
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	nameField, err := writer.CreateFormField("name")
	if err != nil {
		log.Fatal("Error creating form field:", err)
	}
	nameField.Write([]byte(projectName))

	fileField, err := writer.CreateFormFile("file", "file.zip")
	if err != nil {
		log.Fatal("Error creating form file:", err)
	}

	_, err = io.Copy(fileField, file)
	if err != nil {
		log.Fatal("Error copying file data to form field:", err)
	}

	writer.Close()

	return writer, body
}
