package main
​
import (
	"fmt"
	"log"
​
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)
​
func main() {
​
	queueURL := "http://test-message-queue.default.svc.cluster.local:31000/queue/my_queue"
	region := "us-east-1"
	config := &aws.Config{
		Region:      &region,
		Credentials: credentials.NewStaticCredentials("foo", "bar", ""),
	}
​
	sess, err := session.NewSession(config)
	if err != nil {
		log.Panic("Error while creating session")
	}
	svc := sqs.New(sess)
​
	for i := 100; i < 200; i++ {
		msg := fmt.Sprintf("Hello Msg %v", i+1)
		_, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageBody:  &msg,
			QueueUrl:     &queueURL,
		})
		if err != nil {
			log.Panic("Error while writing message")
		}
	}
}
