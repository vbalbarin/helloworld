package main

import (
	"context"
	helloworld "hello-world"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vbalbarin/helloworld"
)

var (
	sess    *session.Session
	scrtmgr *secretsmanager.SecretsManager
)

func init() {
	sess = session.Must(session.NewSession())
	scrtmgr = secretsmanager.New(sess)
}

func handler(req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()
	return helloworld.NewHandler().Handle(ctx, req)
}

func main() {
	lambda.Start(handler)
}
