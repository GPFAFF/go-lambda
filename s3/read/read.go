package read

import (
	file "github.com/GPFAFF/go-lambda/file"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
	"log"
	"os"
)

// TODO fill these in!
const (
	S3_REGION    = "us-east-2"
	S3_LOCALHOST = "http://docker.for.mac.host.internal:4572"
)

func File(s3Bucket string, filename string) {
	// NOTE: you need to store your AWS credentials in ~/.aws/credentials

	// 1) Define your bucket and item names
	bucket := s3Bucket
	item := filename

	downloadedItem, err := os.Create("/tmp/" + item)

	if err != nil {
		fmt.Printf("Error in downloading from file: %v \n", err)
		os.Exit(1)
	}

	fmt.Println(downloadedItem)

	defer downloadedItem.Close()

	// 2) Create an AWS session
	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String(S3_REGION),
		Endpoint: aws.String(S3_LOCALHOST),
	})

	// 3) Create a new AWS S3 downloader
	downloader := s3manager.NewDownloader(sess)

	// 4) Download the item from the bucket. If an error occurs, log it and exit. Otherwise, notify the user that the download succeeded.

	fmt.Println("DDDDD", downloadedItem)

	numBytes, err := downloader.Download(downloadedItem,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
			Range:  aws.String("bytes=0-9"),
		})

	if err != nil {
		log.Fatalf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", downloadedItem.Name(), numBytes, "bytes")

	file.BuildReport(downloadedItem.Name())
}
