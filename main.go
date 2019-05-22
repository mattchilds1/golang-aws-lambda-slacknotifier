package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// CodebuildResponseDetail - Response structure from AWS CodeBuild state change event detail
type CodebuildResponseDetail struct {
	Pipeline    string `json:"pipeline"`
	ExecutionID string `json:"execution-id"`
	Stage       string `json:"stage"`
	State       string `json:"state"`
}

// CodebuildResponse - Response structure from AWS CodeBuild state change event
type CodebuildResponse struct {
	Detail CodebuildResponseDetail `json:"detail"`
}

func handler(ctx context.Context, snsEvent events.SNSEvent) {

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		// each record has: record.EventSource, snsRecord.Timestamp, snsRecord.Message

		codebuildResponse := CodebuildResponse{}
		json.Unmarshal([]byte(snsRecord.Message), &codebuildResponse)

		sendToSlack(codebuildResponse)

	}
}

func buildStatusEmoji(state string) (r string) {
	formattedState := state + " ⚠"
	if state == "SUCCEEDED" {
		formattedState = "SUCCEEDED ✅"
	}
	return formattedState
}

func buildCodeBuildURL(pipelineName string) (r string) {
	return fmt.Sprintf("https://%s.console.aws.amazon.com/codesuite/codepipeline/pipelines/%s/view", os.Getenv("REGION"), pipelineName)
}

func sendToSlack(codebuildResponse CodebuildResponse) {
	attachment1 := slack.Attachment{}
	attachment1.AddAction(slack.Action{Type: "button", Text: "Go to build 🏗️", Url: buildCodeBuildURL(codebuildResponse.Detail.Pipeline), Style: "primary"})
	payload := slack.Payload{
		Text:        fmt.Sprintf("*Pipeline:* %s \n *Stage:* %s \n *State:* %s", codebuildResponse.Detail.Pipeline, codebuildResponse.Detail.Stage, buildStatusEmoji(codebuildResponse.Detail.State)),
		Username:    "Wilbur CodeBuild",
		Channel:     "#aws-notification",
		IconEmoji:   "",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.Send(os.Getenv("WEBHOOK_URL"), "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}

func main() {
	lambda.Start(handler)
}
