package s3handler

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// DeleteObjects apaga um ou mais objetos do bucket S3 // TODO renomear para delete objects
func (client *Client) DeleteObjects(objKey []string, bucketName string) (bool, error) {
	log.Printf("Deletando objeto(s) %s do bucket %s \n", objKey, bucketName)

	var objectIds []types.ObjectIdentifier
	for _, key := range objKey {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}

	params := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{Objects: objectIds},
	}

	output, err := client.s3Client.DeleteObjects(context.TODO(), params)

	if err != nil {
		var noBucket *types.NoSuchBucket
		if errors.As(err, &noBucket) {
			log.Printf("Bucket %s não encontrado", bucketName)
			return false, noBucket
		} else {
			log.Printf("Erro ao deletar objeto(s) do bucket %s: %s \n", bucketName, err.Error())
			return false, err
		}
	}

	if len(output.Errors) > 0 {
		log.Printf("Erro ao deletar objeto(s) do bucket: %s", bucketName)
		for _, outErr := range output.Errors {
			log.Printf("%s: %s\n", *outErr.Key, *outErr.Message)
			err = errors.New(*output.Errors[0].Message)
			return false, err
		}
	}

	for _, delObj := range output.Deleted {

		input := &s3.HeadObjectInput{Bucket: aws.String(bucketName), Key: delObj.Key}

		err = client.objNotExistWaiter().Wait(context.TODO(), input, time.Minute)
		if err != nil {
			log.Printf("Erro ao aguardar o objeto ser deletado: %s", *delObj.Key)
			return false, ErrWaiterTimeout
		} else {
			log.Printf("Objeto %s deletado com sucesso do bucket %s", *delObj.Key, bucketName)
		}
	}

	return true, nil
}

// EmptyBucket esvazia um bucket do S3
func (client *Client) EmptyBucket(bucketName string) (bool, error) {
	log.Printf("Esvaziando bucket %s", bucketName)

	objectsList, err := client.ListObjects(bucketName, 1000)
	if err != nil {
		log.Printf("Erro ao buscar os objetos do bucket: %s", bucketName)
		return false, err
	}

	var deleteList []string

	for _, item := range objectsList {
		deleteList = append(deleteList, *item.Key)
	}

	_, err = client.DeleteObjects(deleteList, bucketName)
	if err != nil {
		log.Printf("Erro ao esvaziar bucket: %s, erro ao deletar objetos: %s", bucketName, deleteList)
		return false, err
	}

	log.Printf("Objetos deletados do bucket: %s com sucesso", bucketName)

	return true, nil
}
