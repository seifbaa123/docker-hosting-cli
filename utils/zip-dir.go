package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipDir(sourceDir string, outputZipFile string) error {
	// Create a new zip file
	zipFile, err := os.Create(outputZipFile)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the source folder and add files to the zip archive
	return filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Open the file for reading
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			// Create a new zip file entry
			zipFileInfo, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			zipFileInfo.Name, _ = filepath.Rel(sourceDir, filePath)

			// Add the file to the zip archive
			fileWriter, err := zipWriter.CreateHeader(zipFileInfo)
			if err != nil {
				return err
			}

			// Copy the file's contents to the zip archive
			_, err = io.Copy(fileWriter, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
