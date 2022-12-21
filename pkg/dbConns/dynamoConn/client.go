package dynamoConn

import (
	"sync"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type sinleton struct {
	sync.Once
	dynamoDB *dynamodb.DynamoDB
}

var (
	singletonDB singleton
)

func InitDB() error {
	var initializeError error
	singletonDB.Do(func(){
		sessionDynamo, err := session.NewSession(&aws.Config{
			Region: aws.String(viper.GetString("dynamodb.region"))
		})
		initializeError = err
		singletonDB.dynamodb.New(sessionDynamo)
	})
	return err
}

func GetDB() *dynamodb.DynamoDB {
	return singletonDB.dynamoDB
}