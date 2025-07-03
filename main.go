package main

import (
	"amazon-handler/s3handler"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func exemploCreateBucket(client *s3handler.Client) {

	log.Println("Criando bucket...")
	_, err := client.CreateBucket("teste")
	if err != nil {
		log.Printf("Erro ao criar bucket: %s", err.Error())
	} else {
		log.Println("Bucket criado com sucesso.")
	}
}

func exemploDeleteBucket(client *s3handler.Client) {
	log.Println("Deletando bucket...")
	_, err := client.DeleteBucket("teste")
	if err != nil {
		log.Printf("Erro ao deletar bucket: %s", err.Error())
	} else {
		log.Println("Bucket deletado com sucesso.")
	}
}

func exemploDeleteObjects(client *s3handler.Client) {
	log.Println("Deletando objeto...")
	_, err := client.DeleteObjects([]string{"exemplo-teste/exemplo.html"}, "teste")
	if err != nil {
		log.Printf("Erro ao deletar objeto: %s", err.Error())
	} else {
		log.Println("Objeto deletado com sucesso.")
	}
}

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

func exemploUpload(client *s3handler.Client) {
	log.Println("Fazendo upload de objeto...")
	_, err := client.Upload("teste", "exemplo-teste", "/home/angelo/Documentos/Programação/exemplo.html")
	if err != nil {
		log.Printf("Erro ao fazer upload: %s", err.Error())
	} else {
		log.Println("Upload realizado com sucesso.")
	}
}

func main() {
	log.Println("main")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações. (%e)", err)
	}

	client := s3.NewFromConfig(cfg)

	handler := s3handler.NewS3Client(client)

	exemploCreateBucket(handler)
	exemploDeleteBucket(handler)
	exemploDeleteObjects(handler)
	exemploListBuckets(handler)
	exemploListObjects(handler)
	exemploUpload(handler)
}
