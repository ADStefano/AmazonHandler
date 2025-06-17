package main

import (
	"amazon-handler/s3"
	"log"
)

func exemploCreateBucket(client *s3.Client) {

	log.Println("Criando bucket...")
	_, err := client.CreateBucket("meu-segundo-bucket123")
	if err != nil {
		log.Printf("Erro ao criar bucket: %s", err.Error())
	} else {
		log.Println("Bucket criado com sucesso.")
	}
}

func exemploDeleteBucket(client *s3.Client) {
	log.Println("Deletando bucket...")
	_, err := client.DeleteBucket("meu-segundo-bucket123")
	if err != nil {
		log.Printf("Erro ao deletar bucket: %s", err.Error())
	} else {
		log.Println("Bucket deletado com sucesso.")
	}
}

func exemploDeleteObjects(client *s3.Client) {
	log.Println("Deletando objeto...")
	_, err := client.DeleteObjects([]string{"exemplo-teste.html"}, "test")
	if err != nil {
		log.Printf("Erro ao deletar objeto: %s", err.Error())
	} else {
		log.Println("Objeto deletado com sucesso.")
	}
}

func exemploListBuckets(client *s3.Client) {
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

func exemploListObjects(client *s3.Client) {
	log.Println("Listando objetos...")
	objects, err := client.ListObjects("test", 5)
	if err != nil {
		log.Printf("Erro ao listar objetos: %s", err.Error())
		return
	}

	for _, item := range objects {
		log.Printf("Key: %s, LastModified: %s, ETag: %s, Size: %d, StorageClass: %v", *item.Key, *item.LastModified, *item.ETag, *item.Size, item.StorageClass)
	}
	log.Println("Listagem de objetos concluída com sucesso.")
}

func exemploUpload(client *s3.Client) {
	log.Println("Fazendo upload de objeto...")
	_, err := client.Upload("test", "exemplo-teste", "/home/angelo/Documentos/Programação/exemplo.html")
	if err != nil {
		log.Printf("Erro ao fazer upload: %s", err.Error())
	} else {
		log.Println("Upload realizado com sucesso.")
	}
}

func main() {
	log.Println("main")
	client := s3.NewS3Client()

	exemploCreateBucket(client)
	exemploDeleteBucket(client)
	exemploDeleteObjects(client)
	exemploListBuckets(client)
	exemploListObjects(client)
	exemploUpload(client)
}
