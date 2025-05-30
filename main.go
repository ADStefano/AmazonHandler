package main

import (
	"amazon-handler/s3"
	"log"
)

func main() {
	log.Println("main")
	teste := s3.NewS3Client()
	// teste.CreateBucket("test")
	// teste.DeleteBucket("test-meu-segundo-bucket123")
	// teste.DeleteObjects([]string{"exemplo-teste.html"}, "test")
	// objects, errors := teste.ListObjects("test", 5)

	// if errors != nil{
	// 	print(errors)
	// }

	// for _, item := range objects{
	// 	log.Printf("Key: %s, LastModified: %s, ETag: %s, Size: %d, StorageClass: %v", *item.Key, *item.LastModified, *item.ETag, *item.Size, item.StorageClass)
	// }
	teste.Upload("test", "teste", "/home/angelo/Documentos/Programação/exemplo.html")

}
