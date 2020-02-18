package receive

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Message() {

	fmt.Println("Blah blah blah")
	// struct gets passed in.
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	service := sqs.New(session, &aws.Config{Endpoint: aws.String("http://localhost:4576")})

	qURL := "http://localhost:4576/queue/development-queue"

	result, err := service.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &qURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(20),
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		fmt.Println("Error", err)
	}

	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
		return
	}

	fmt.Println(result)
}
