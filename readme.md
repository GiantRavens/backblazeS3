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
2. Issue an application key that only has access to that bucket. 

Make note of the server the bucket is created on as well, and enter that information in `config/config.go` to match your information.

## Choose a File and Directory to Upload/Download

in `backblaze_example_app.go` edit the name of the file you want to upload, then the destination to download it again to prove that you have uploaded, and downloaded a file.

## Run Your App

Run the app with:

```golang
  go run backblaze_example_app.go
```

The first time you run the app the AWS modules will be downloaded as needed.

## TODO - Areas to Explore Further
1. return actual errors in functions instead of bool
2. store config info key/secrets management, with viper, or directly in a struct
3. full golang linting i.e. proper Go comments and layout
