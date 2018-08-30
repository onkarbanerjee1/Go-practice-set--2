package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, client statsd.Client, snsEvent events.SNSEvent) {
	message := []string{}
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		message = append(message, fmt.Sprintf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message))

	}
	ev := statsd.Event{
		Title:          "Error",
		Text:           strings.Join(message, "\n"),
		Hostname:       "example.com",
		AggregationKey: "Errors",
		Priority:       statsd.Normal,
		SourceTypeName: "Sample",
		AlertType:      statsd.Error,
		Tags:           []string{"PipelineErrors", "Pipeline", "Errors"},
	}

	if err := ev.Check(); err != nil {
		log.Fatal(err)
	}

	if err := client.Event(&ev); err != nil {
		log.Fatal(err)
	}

}

func main() {
	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}

	c.Namespace = "PipelineErrors."
	c.Tags = append(c.Tags, "us-east-1a")

	fmt.Println("all good")

	lambda.Start(handler)
}
