package send

import (
	"fmt"

	"github.com/GPFAFF/go-lambda/record"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Message sends a simple message to the queue
func Message(vehicle record.VehicleData) {

	// struct gets passed in.
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	qURL := "http://localhost:4576/queue/development-queue"

	service := sqs.New(session, &aws.Config{Endpoint: aws.String("http://localhost:4576")})

	result, err := service.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"VIN": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(vehicle.VIN),
			},
			"OrigDealerID": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(vehicle.OrigDealerID),
			},
			"ProgramCode": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(vehicle.ProgramCode),
			},
			"Date": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(vehicle.Date),
			},
			"Status": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(vehicle.Status),
			},
		},
		MessageBody: aws.String("Program Car Utilization"),
		QueueUrl:    &qURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result)
}
