package main

import (
	"backblaze_go/backblazeS3"
	"backblaze_go/config"
	"log"
	"fmt"
)

func main() {
	// set logging flags
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// read config
	config.Init()

	// new Backblaze B2 client
	b2Client, err := backblazeS3.NewB2Client(config.GetConfig().B2)
	if err != nil {
		log.Fatalln("[fatal][app] failed to get B2 client", err)
	}

	// list objects in the bucket
	fmt.Println("")
	fmt.Println("First, counting and listing the number of files in your target bucket:")
	b2Client.List()

	// upload object
	fmt.Println("")
	fmt.Println("Now uploading the file you've specified...")
	b2Client.Upload("CS_Nodecraft.pdf", "dir_upload/CS_Nodecraft.pdf")

	fmt.Println("")
	fmt.Println("Listing bucket contents again. Note that if the file already existed, and 'keep all file versions' is set for that bucket - you're uploading and keeping another version and not replacing.")
	b2Client.List()

	// download object
	fmt.Println("")
	fmt.Println("Now downloading that same file to a different folder to confirm that it worked...")
  b2Client.Download("CS_Nodecraft.pdf", "dir_download/CS_Nodecraft.pdf")

	// delete object
	// b2Client.Delete("CS_Nodecraft.pdf")

  fmt.Println("Complete")
	fmt.Println("")
}
