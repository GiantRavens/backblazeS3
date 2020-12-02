# Simple POC for Go with Backblaze

This simple POC shows how you can use the Amazon AWS S3 libs to connect to Backblaze B2 as a starting point for building your own applications.

## Setting Up Your Environment

Make sure you have a running install of Go.

On a Mac, you can use the brew package manager to install both Go and the dep package manager with:

```shell
  brew install go
  brew install dep
```

Next, move the code to your $GOPATH/src directory.

From that directory, resolve dependencies with:

```shell
  dep init
  dep ensure -v
```

## Configure Your Bucket and Connection Information

Create your target bucket in your Backblaze Account Management Page, then issue an application key that only has access to that bucket. Make note of the server the bucket is created on as well.

With that information you can edit the `backblaze_go/config/config.go` to match your settings.

## Choose a File and Directory to Upload/Download

in `backblaze_go/app.go` edit the name of the file you want to upload, then the destination to download it again to prove that you have uploaded, and downloaded a file.

## Run Your App

From your $GOPATH/src/ dir test it out with:

```golang
  go run backblaze_example_app.go
```
