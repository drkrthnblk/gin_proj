package sqsConn

import (
	"sync"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type singleton struct {
	sync.Once
	sqs *sqs.SQS
	sqsSession *session.Session
}

var (
	singletonDB singleton
)

func InitDB() error {
	var initializeError error
	singletonDB.Do(func(){
		sessionSQS, err := session.NewSession(&aws.Config{
			Region: aws.String(viper.GetString("dynamodb.region")),
		})
		initializeError = err
		singletonDB.sqs = sqs.New(sessionSQS)
		singletonDB.sqsSession = sessionSQS
	})
	return initializeError
}

func GetDB() *sqs.SQS {
	return singletonDB.sqs
}

func Gets3Session() *session.Session {
	return singletonDB.sqsSession
}