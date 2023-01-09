package dbconns

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
	"gorm.io/gorm"
)

type dbClient interface {
	*gorm.DB | *dynamodb.DynamoDB | *s3.S3 | *sqs.SQS
}

type dbSessioon interface {
	*session.Session | *gorm.Session
}

type DbConnSelector interface {
	InitDB() error
	GetClient() (*interface{}, error)
	GetSession() (*interface{}, error)
}
