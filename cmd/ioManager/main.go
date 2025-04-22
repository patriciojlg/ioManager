package main

import (
	"context"
	"encoding/json"
	"errors"
	"ioManager/controllers"
	"ioManager/models"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, rawEvent json.RawMessage) (models.Response, error) {
	var evt models.Event
	if err := json.Unmarshal(rawEvent, &evt); err != nil {
		return models.Error400Response(err), nil
	}

	switch evt.Command {
	// INPUTS COMMANDS
	case "save-input":
		return controllers.SaveInput(evt.Args), nil
	case "explode-input":
		return controllers.ExplodeInput(evt.Args), nil

	// TASK'S CONFIGURATION COMMANDS
	case "get-task-configuration":
		resp, err := controllers.GetConfigurationTemplate(evt.Args["taskName"].(string))
		if err != nil {
			return models.Error400Response(err), nil
		}
		return resp, err

	//case "explode-input":
	//	return controllers.ExplodeInput(evt.Args)
	default:
		return models.Error400Response(errors.New("command not found")), nil
	}
}

func main() {
	if _, exists := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); exists {
		lambda.Start(handler)
	} else {
		ctx := context.Context(context.Background())
		lambdaResponse, err := handler(ctx, json.RawMessage(`{"command":"save-input","args":{"id":"1234","format":"json"}}`))
		if err != nil {
			log.Fatal(err)
		}
		jsonResponse, err := json.Marshal(lambdaResponse)
		log.Printf(`Response: %v`, string(jsonResponse))
	}
}
