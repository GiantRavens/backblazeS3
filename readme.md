# Connecting Go Apps to Backblaze B2 with S3

This simple POC shows how you can use the Amazon AWS S3 modules to connect to Backblaze B2 as a starting point for building your own applications.

## Setting Up Your Environment

Make sure you have a running install of Go.

On a Mac, you can use the brew package manager to install Go with:

```shell
  brew install go
```
Move into your $GOPATH directory, clone the repo.

## Configure Your Bucket and Connection Information

1. Create your target bucket in your Backblaze Account Management Page
2. Issue an application key that only has access to that bucket - be sure that 'list all buckets' option is selected.

Make note of the server endpoint associated with the bucket, and enter this information in `.env` to match your bucket, server, and appkey details.

## Run Your App

Issue `go mod init`, then run the app with:

```golang
  go run backblaze_example_app.go
```

The first time you run the app the AWS modules will be downloaded as needed.

A sample file is already included in dir_upload. You'll see the file upload, the bucket contents listed, then the file downloaded again to a separate folder. Finally, the delete function will be exercised and the file will be deleted from the bucket.

From here -  you can explore passing in new filenames, dynamic filenames, and more. 
