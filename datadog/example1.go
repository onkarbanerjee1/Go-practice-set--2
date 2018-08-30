package main

import (
	"fmt"
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}

	c.Namespace = "PipelineErrors."
	c.Tags = append(c.Tags, "us-east-1a")

	ev := statsd.Event{
		Title:          "Error",
		Text:           "An error occurred now",
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

	if err := c.Event(&ev); err != nil {
		log.Fatal(err)
	}

	fmt.Println("all good")
}
