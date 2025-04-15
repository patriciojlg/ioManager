package main

import (
	"context"
	"encoding/json"
	"errors"
	"ioManager/controllers"
	"ioManager/models"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, rawEvent json.RawMessage) models.Response {
	var evt models.Event
	if err := json.Unmarshal(rawEvent, &evt); err != nil {
		return models.Error400Response(err)
	}

	switch evt.Command {
	case "save-input":
		return controllers.SaveInput(evt.Args)
	//case "explode-input":
	//	return controllers.ExplodeInput(evt.Args)
	default:
		return models.Error400Response(errors.New("command not found"))
	}
}

func main() {
	lambda.Start(handler)
}
