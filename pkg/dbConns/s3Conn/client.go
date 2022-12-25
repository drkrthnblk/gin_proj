package s3Conn

import (
	"sync"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type singleton struct {
	sync.Once
	s3 *s3.S3
	s3Session *session.Session
}

var (
	singletonDB singleton
)

func InitDB() error {
	var initializeError error
	singletonDB.Do(func(){
		sessionS3, err := session.NewSession(&aws.Config{
			Region: aws.String(viper.GetString("dynamodb.region")),
		})
		initializeError = err
		singletonDB.s3 = s3.New(sessionS3)
		singletonDB.s3Session = sessionS3
	})
	return initializeError
}

func GetDB() *s3.S3 {
	return singletonDB.s3
}

func Gets3Session() *session.Session {
	return singletonDB.s3Session
}