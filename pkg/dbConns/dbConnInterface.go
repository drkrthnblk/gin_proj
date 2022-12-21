package dbConns

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws/session"
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
	GetClient[C dbClient]() C
	GetSession[S dbSession]() S
}
