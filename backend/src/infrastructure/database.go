package infrastructure

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
	"os"
)

func GetDB() *dynamo.DB {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	const AWS_REGION = "ap-northeast-1"
	DATABASE_ENDPOINT := os.Getenv("DATABASE_ENDPOINT")
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Endpoint:    aws.String(DATABASE_ENDPOINT),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}
	return dynamo.New(session)
}
