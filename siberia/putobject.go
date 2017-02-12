package main

import (
	"log"

	minio "github.com/minio/minio-go"
)

func main() {
	// Initialize a new s3Client.
	s3Client, err := minio.New("play.minio.io:9000", "YOUR-ACCESSKEYID",
		"YOUR-SECRETACCESSKEY", true)
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
