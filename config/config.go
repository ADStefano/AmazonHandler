package config

import (
	"log"
	"os"
	"strconv"
)

type AWSConfig struct {
	Region           string
	Profile          string
	CustomEndpoint   string
	RetryMaxAttempts int
}

// Carrega as configurações da Amazon com base nas variáveis no ambiente
func LoadConfig() (*AWSConfig, error) {
	print("placeholder")

	region := os.Getenv("AWS_REGION")
	if region == "" {
		log.Println("Região não definida, utilizando padrão: sa-east-1")
		region = "sa-east-1"
	}

	profile := os.Getenv("AWS_PROFILE")
	if profile == "" {
		log.Println("Profile não definido, utilizando padrão: default")
		profile = "default"
	}

	retryMaxAttempts := 3
	customEndpoint := os.Getenv("AWS_CUSTOM_ENDPOINT")

	retries := os.Getenv("AWS_RETRY_MAX_ATTEMPTS")
	if retries != "" {
		var err error
		retryMaxAttempts, err = strconv.Atoi(retries)
		if err != nil {
			log.Printf("Erro ao configurar RetryMaxAttempts (%e), utilizando padrâo: 3", err)
			retryMaxAttempts = 3
		}
	}

	return &AWSConfig{
		Region:           region,
		Profile:          profile,
		CustomEndpoint:   customEndpoint,
		RetryMaxAttempts: retryMaxAttempts,
	}, nil
}
