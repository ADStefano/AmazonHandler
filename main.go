package main

import (
	"amazon-handler/s3handler"
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Exemplo de uso da função CreateBucket
func exemploCreateBucket(client *s3handler.Client) {

	log.Println("Criando bucket...")
	_, err := client.CreateBucket("teste")
	if err != nil {
		log.Printf("Erro ao criar bucket: %s", err.Error())
	} else {
		log.Println("Bucket criado com sucesso.")
	}
}

// Exemplo de uso da função DeleteBucket
func exemploDeleteBucket(client *s3handler.Client) {
	log.Println("Deletando bucket...")
	_, err := client.DeleteBucket("teste")
	if err != nil {
		log.Printf("Erro ao deletar bucket: %s", err.Error())
	} else {
		log.Println("Bucket deletado com sucesso.")
	}
}

// Exemplo de uso da função DeleteObjects
func exemploDeleteObjects(client *s3handler.Client) {
	log.Println("Deletando objeto...")
	_, err := client.DeleteObjects([]string{"exemplo-teste/exemplo.html"}, "teste")
	if err != nil {
		log.Printf("Erro ao deletar objeto: %s", err.Error())
	} else {
		log.Println("Objeto deletado com sucesso.")
	}
}

// Exemplo de uso da função ListBuckets
func exemploListBuckets(client *s3handler.Client) {
	log.Println("Listando buckets...")
	buckets, err := client.ListBuckets()
	if err != nil {
		log.Printf("Erro ao listar buckets: %s", err.Error())
		return
	}

	for _, bucket := range buckets {
		log.Printf("Bucket: %s, Criado em: %s", *bucket.Name, bucket.CreationDate.String())
	}
	log.Println("Listagem de buckets concluída com sucesso.")
}

// Exemplo de uso da função ListObjects
func exemploListObjects(client *s3handler.Client) {
	log.Println("Listando objetos...")
	objects, err := client.ListObjects("teste", 5)
	if err != nil {
		log.Printf("Erro ao listar objetos: %s", err.Error())
		return
	}

	for _, item := range objects {
		log.Printf("Key: %s, LastModified: %s, ETag: %s, Size: %d, StorageClass: %v", *item.Key, *item.LastModified, *item.ETag, *item.Size, item.StorageClass)
	}
	log.Println("Listagem de objetos concluída com sucesso.")
}

// Exemplo de uso da função Upload
func exemploUpload(client *s3handler.Client) {
	log.Println("Fazendo upload de objeto...")
	_, err := client.UploadS3("teste", "exemplo-teste", "/home/angelo/Documentos/Programação/exemplo.html")
	if err != nil {
		log.Printf("Erro ao fazer upload: %s", err.Error())
	} else {
		log.Println("Upload realizado com sucesso.")
	}
}

// Exemplo de uso da função Download
func exemploDownload(client *s3handler.Client) {

	log.Println("Baixando objeto...")

	output, err := client.DownloadS3("adstefano", "teste/exemplo.html")

	if err != nil {
		log.Printf("Erro ao baixar objeto: %s", err.Error())
		return
	}

	file, err := os.Create("exemplo.html")
	if err != nil {
		log.Printf("Erro ao criar arquivo: %s", err.Error())
		return
	}

	defer file.Close()

	_, err = io.Copy(file, output.Body)
	if err != nil {
		log.Printf("Erro ao copiar conteúdo: %s", err.Error())
		return
	}

	defer output.Body.Close()

	log.Println("Download concluído com sucesso.")
}

func exemploGetPresignedURL(client *s3handler.Client) {

	log.Println("Gerando URL pré-assinada...")

	url, err := client.GetPreSignedURL("teste", "teste/exemplo.html", 600*time.Second)
	if err != nil {
		log.Printf("Erro ao gerar URL pré-assinada: %s", err.Error())
		return
	}

	log.Printf("URL pré-assinada gerada com sucesso: %s", url.URL)
}

func main() {
	log.Println("main")

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações. (%e)", err)
	}

	client := s3.NewFromConfig(cfg)

	handler := s3handler.NewS3Client(client)

	// exemploCreateBucket(handler)
	// exemploDeleteBucket(handler)
	// exemploDeleteObjects(handler)
	// exemploListBuckets(handler)
	// exemploListObjects(handler)
	// exemploUpload(handler)
	// exemploDownload(handler)
	exemploGetPresignedURL(handler)
	
}
