package send

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Message sends a simple message to the queue
func Message(message string) {

	fmt.Println(message)
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	qURL := "http://localhost:4576/queue/development-queue"

	service := sqs.New(session, &aws.Config{Endpoint: aws.String("http://localhost:4576")})

	result, err := service.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"UtilizationReport": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(message),
			},
		},
		MessageBody: aws.String("Program Car Utilization"),
		QueueUrl:    &qURL,
	})

	if err != nil {
		fmt.Println("Errorzzz", err)
		return
	}

	fmt.Println("Success", *result)
}
