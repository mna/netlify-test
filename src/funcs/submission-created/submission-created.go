package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kr/pretty"
)

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(r.Body), &m); err != nil {
		fmt.Println(err)
	} else {
		pretty.Printf("%# v\n", m["payload"])
	}
	for k, v := range r.Headers {
		if strings.HasPrefix(k, "X-") {
			continue
		}
		fmt.Println(k, v)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 204,
	}, nil
}

func main() {
	lambda.Start(handler)
}
