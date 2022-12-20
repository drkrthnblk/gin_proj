package dbconns

type DbConnType string

const (
	POSTGRESCONN DbConnType = "POSTGRES_CONN"
	DYNAMOCONN   DbConnType = "DYNAMO_CONN"
	REDISCONN    DbConnType = "REDIS_CONN"
	S3CONN       DbConnType = "S3_CONN"
	SQSCONN      DbConnType = "SQS_CONN"
)
