package main

import (
	"log"

	minio "github.com/minio/minio-go"
)

func main() {
	// Initialize a new s3Client.
	s3Client, err := minio.New("play.minio.io:9000", "Q3AM3UQ867SPQQA43P2F",
		"zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", true)
	if err != nil {
		log.Fatalln(err)
	}

	// Read file and upload.
	_, err = s3Client.FPutObject("my-bucketname", "my-objectname",
		"~/Videos/my-movie.avi", "application/video")
	if err != nil {
		log.Fatalln(err)
	}
}
