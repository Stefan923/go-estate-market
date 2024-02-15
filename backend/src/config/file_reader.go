package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type FileReader[T any] struct {
}

func (reader *FileReader[T]) GetContent(filePath string) *T {
	fileReader, err := reader.LoadFileContent(filePath, "yml")
	if err != nil {
		log.Fatalf("Error while loading file content: %v", err)
	}

	fileContent, err := reader.ParseFileContent(fileReader)
	if err != nil {
		log.Fatalf("Error while parsing file content: %v", err)
	}

	return fileContent
}

func (reader *FileReader[T]) ParseFileContent(fileReader *viper.Viper) (*T, error) {
	var fileContent T
	err := fileReader.Unmarshal(&fileContent)
	if err != nil {
		log.Printf("Unable to parse file: %v", err)
		return nil, err
	}
	return &fileContent, nil
}

func (reader *FileReader[T]) LoadFileContent(fileName string, fileType string) (*viper.Viper, error) {
	fileReader := viper.New()
	fileReader.SetConfigType(fileType)
	fileReader.SetConfigName(fileName)
	fileReader.AddConfigPath(".")
	fileReader.AutomaticEnv()

	err := fileReader.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read file content: %v", err)
		var fileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &fileNotFoundError) {
			return nil, errors.New("could not find file")
		}
		return nil, err
	}
	return fileReader, nil
}
