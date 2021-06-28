package main
​
import (
	"fmt"
	"log"
	"net/http"
​
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)
​
func Handler(w http.ResponseWriter, r *http.Request) {
	queueURL := "http://localstack.default.svc.cluster.local:31000/queue/my_queue"
	region := "us-east-2"
	config := &aws.Config{
		Region:      &region,
		Credentials: credentials.NewStaticCredentials("foo", "bar", ""),
	}
	sess, err := session.NewSession(config)
	if err != nil {
		log.Panic("Error while creating session")
	}
	svc := sqs.New(sess)
	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("message count %v", i+1)
		_, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageBody:  &msg,
			QueueUrl:     &queueURL,
		})
		if err != nil {
			w.Write([]byte(fmt.Sprintf("failed to send message to input queue: %v", err)))
			return
		}
	}
	w.Write([]byte("successfully sent message to input queue"))
}