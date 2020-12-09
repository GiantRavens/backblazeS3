# Simple POC for Go with Backblaze

This simple POC shows how you can use the Amazon AWS S3 libs to connect to Backblaze B2 as a starting point for building your own applications.

## Setting Up Your Environment

Make sure you have a running install of Go.

On a Mac, you can use the brew package manager to install Go with:

```shell
  brew install go
```
Move into your $GOPATH directory, clone the repo.

In that directory issue:

`go mod init`

## Configure Your Bucket and Connection Information

1. Create your target bucket in your Backblaze Account Management Page
2. Issue an application key that only has access to that bucket - be sure that 'list all buckets' option is selected.

Make note of the server the bucket is created, and enter that information in `.env` to match your bucket, server, and appkey.

## Run Your App

Run the app with:

```golang
  go run backblaze_example_app.go
```

The first time you run the app the AWS modules will be downloaded as needed.

From here -  you can explore passing in filenames, changing the upload and download directories, and more.
