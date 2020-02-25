// main.go
package main

import (
	"context"
	"fmt"

	"github.com/GPFAFF/go-lambda/s3/read"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		read.File(s3.Bucket.Name, s3.Object.Key)
	}
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	fmt.Println("Lambda listening")
	lambda.Start(handler)
}
