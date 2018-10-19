package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/kr/pretty"
)

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var title string
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		title = lc.ClientContext.Client.AppTitle
	}

	fmt.Printf("%# v\n", pretty.Formatter(r))

	user := r.RequestContext.AccountID
	path := r.Path
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello, %s on path %s of app %s", user, path, title),
	}, nil
}

func main() {
	lambda.Start(handler)
}
