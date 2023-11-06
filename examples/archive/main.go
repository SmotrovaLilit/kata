package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"os"
	"strings"
)

func main() {
	files := []ArchiveFile{
		{"file1.txt", "Content of file 1.txt"},
		{"file2.txt", "Content of file 2.txt"},
	}
	err := createZip("archive.zip", files)
	if err != nil {
		panic(err)
	}
	fmt.Println("Zip-archive created.")
	err = createGzip("archive.gz", "Content of  archive.gz")
	if err != nil {
		panic(err)
	}
	fmt.Println("GZ-archive created.")
	err = createTarGz("archive.tar.gz", files)
	if err != nil {
		panic(err)
	}
	fmt.Println("TarGz-archive created.")
}

type ArchiveFile struct {
	Name    string
	Content string
}

const (
	x int = 0
)

func createZip(name string, files []ArchiveFile) error {
	zipFile, err := os.Create(name)
	if err != nil {
		return err
	}
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	for i := 0; i < len(files); i++ {
		w, err := zipWriter.Create(files[i].Name)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(files[i].Content))
		if err != nil {
			return err
		}
	}

	return nil
}

func createGzip(name string, fileContent string) error {
	outputFile, err := os.Create(name)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	_, err = strings.NewReader(fileContent).WriteTo(gzipWriter)
	if err != nil {
		return err
	}
	return nil
}

func createTarGz(name string, files []ArchiveFile) error {
	outputFile, err := os.Create(name)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()
	for _, file := range files {
		header := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Content)),
		}
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		_, err := tarWriter.Write([]byte(file.Content))
		if err != nil {
			return err
		}
	}
	return nil
}
