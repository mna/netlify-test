package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if r.RequestContext.ResourcePath == "/fail" {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Fail!",
		}, nil
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%v", r.QueryStringParameters) + "\n" + r.Resource + "\n" + r.RequestContext.ResourcePath + "\n" + r.Path + "\n" + os.Getenv("NETLIFY_TEST_GREET"),
	}, nil
}

func main() {
	lambda.Start(handler)
}
