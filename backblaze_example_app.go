package main

import (
	"backblaze_go/backblazeS3"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// set logging flags
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	keyID := os.Getenv("KEY_ID")
	applicationKey := os.Getenv("APPLICATION_KEY")
	bucketName := os.Getenv("BUCKET_NAME")
	endpoint := os.Getenv("ENDPOINT")
	region := os.Getenv("REGION")

	// new Backblaze B2 client
	b2Client, err := backblazeS3.NewB2Client(endpoint, region, keyID, applicationKey, "", bucketName)
	if err != nil {
		log.Fatalln("[fatal][app] failed to get B2 client", err)
	}

	// list objects in the bucket

	listResult, err := b2Client.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range listResult {
		log.Println(result)
	}

	// upload object
	fmt.Println("")
	fmt.Println("Now uploading the file you've specified...")
	if err := b2Client.Upload("CS_Nodecraft.pdf", "dir_upload/CS_Nodecraft.pdf"); err != nil {
		log.Fatal("[err][b2][upload] failed to upload file", err)
	}

	// download object
	if err := b2Client.Download("CS_Nodecraft.pdf", "dir_download/CS_Nodecraft.pdf"); err != nil {
		log.Fatal("[err][b2][download] failed to download file", err)
	}

	// delete object
	if err := b2Client.Delete("CS_Nodecraft.pdf"); err != nil {
		log.Fatal("[err][b2] failed to delete object", err)
	}

}
