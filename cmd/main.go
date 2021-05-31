package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// TODO web page to register
// TODO connect to db
// TODO define sql table
// TODO post request to insert data

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
