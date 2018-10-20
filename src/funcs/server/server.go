package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	path := r.QueryStringParameters["path"]
	if path == "fail" {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Fail!",
		}, nil
	}
	var body bytes.Buffer
	for k, v := range r.Headers {
		fmt.Fprintf(&body, "%s = %v\n", k, v)
	}
	fmt.Fprintf(&body, "Method = %s\n", r.HTTPMethod)
	fmt.Fprintf(&body, "Path = %s\n", r.Path)
	fmt.Fprintf(&body, "PathParameters = %v\n", r.PathParameters)
	fmt.Fprintf(&body, "QueryStringParameters = %v\n", r.QueryStringParameters)
	fmt.Fprintf(&body, "Resource = %s\n", r.Resource)
	fmt.Fprintf(&body, "StageVariables = %v\n", r.StageVariables)
	fmt.Fprintf(&body, "NETLIFY_TEST_GREET = %s\n", os.Getenv("NETLIFY_TEST_GREET"))

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body.String(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
