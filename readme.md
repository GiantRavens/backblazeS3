# Connecting Go Apps to Backblaze B2 with S3

This simple POC shows how you can use the Amazon AWS S3 modules to connect to Backblaze B2 as a starting point for building your own applications in Go.

## Setting Up Your Environment

Make sure you have a [running install of Go](https://golang.org/doc/install).

On a Mac, you can use the brew package manager to install Go with:

```shell
  brew install go
```
Move into your $GOPATH directory, clone the repo.

## Configure Your Bucket and Connection Information

If you aren't already signed up for Backblaze B2 [sign up for Backblaze B2 here](https://www.backblaze.com/b2/sign-up.html?referrer=giantravens) - your first 10GB/month are always no charge, after that $5/TB/month.

1. Create your target bucket in your Backblaze Account Management Page. 
2. Issue an application key that only has access to that bucket - be sure that the 'list all buckets' option is selected.

Make note of the server endpoint associated with the bucket, and edit your configuration in `.env` to match your bucket, server, and appkey details.

## Run Your App

Issue `go mod init`, then run the app with:

```golang
  go run backblaze_example_app.go
```

The first time you run the app the AWS modules will be downloaded as needed.

A sample file is already included in dir_upload. When you run this app, you'll see the file upload, the bucket contents listed, then the file downloaded again to a separate folder. Finally, the delete function will be exercised and the file will be deleted from the bucket.

From here -  you can explore passing in different filenames, dynamic filenames, and more. 
