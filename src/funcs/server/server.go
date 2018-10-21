package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kr/pretty"
)

const funcPrefix = "/.netlify/functions/server"

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	path := strings.TrimPrefix(r.Path, funcPrefix)
	if path == "/fail" {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Fail!",
		}, nil
	}
	var body bytes.Buffer
	fmt.Fprintln(&body, "APIGatewayProxyRequest:")
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

	fmt.Fprintln(&body, "\nRequestContext:")
	fmt.Fprintf(&body, "APIID = %s\n", r.RequestContext.APIID)
	fmt.Fprintf(&body, "AccountID = %s\n", r.RequestContext.AccountID)
	fmt.Fprintf(&body, "Authorizer = %v\n", r.RequestContext.Authorizer)
	fmt.Fprintf(&body, "HTTPMethod = %s\n", r.RequestContext.HTTPMethod)
	fmt.Fprintf(&body, "RequestID = %s\n", r.RequestContext.RequestID)
	fmt.Fprintf(&body, "ResourceID = %s\n", r.RequestContext.ResourceID)
	fmt.Fprintf(&body, "ResourcePath = %s\n", r.RequestContext.ResourcePath)
	fmt.Fprintf(&body, "Stage = %s\n", r.RequestContext.Stage)

	fmt.Fprintln(&body, "\nIdentity:")
	pretty.Fprintf(&body, "%#v\n", r.RequestContext.Identity)

	fmt.Fprintln(&body, "\nEnv Vars:")
	envs := os.Environ()
	for _, v := range envs {
		fmt.Fprintf(&body, "%v\n", v)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body.String(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
