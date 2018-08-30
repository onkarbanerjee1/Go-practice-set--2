package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/medivo/nucleus/log"
	"github.com/medivo/nucleus/web"
	"github.com/medivo/security_defaults/go/awssess"
)

type snsMsg struct {
	Message string
}

// SNSMsgTypeHdr holds the header for SNS message.
const SNSMsgTypeHdr = "x-amz-sns-message-type"

// Message Types
const (
	SubscriptionMsg = "SubscriptionConfirmation"
	NotificationMsg = "Notification"
)

var (
	snssvc *sns.SNS
	c      *statsd.Client
)

func main() {
	if sess, err := awssess.New(); err != nil {
		log.Error("failed to initialize AWS:", strings.Replace(err.Error(), "\n", " ", -1))

	} else {
		snssvc = sns.New(sess)
	}

	// We need the watcher to be initialized before we can get the handler.
	routes := web.RoutesMap{"/notify": Notify, "/check": Check}
	err := web.Setup(web.Config{
		Address: ":3000",
		Routes:  routes,
		Logfile: filepath.Join("log", "development.log"),
	})
	if err != nil {
		fmt.Println("Error in webs etup", err)
	}

	err = web.Server.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
		fmt.Println("Stopping HTTPS server:", err)
		log.Error("Stopping HTTPS server:", err)
	}

}

func snsConfirmSubscription(body io.ReadCloser) error {
	subscribeOutput := struct{ TopicARN, Token string }{}
	dec := json.NewDecoder(body)
	if err := dec.Decode(&subscribeOutput); err != nil {
		return err
	}

	params := &sns.ConfirmSubscriptionInput{
		Token:                     &subscribeOutput.Token,
		TopicArn:                  &subscribeOutput.TopicARN,
		AuthenticateOnUnsubscribe: aws.String("true"),
	}

	resp, err := snssvc.ConfirmSubscription(params)
	log.Debug("Subscription response:", resp)

	return err
}

// Notify handles the corresponding API call.
func Notify(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	msgType, body := r.Header.Get(SNSMsgTypeHdr), r.Body
	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}

	c.Namespace = "PipelineErrors."
	c.Tags = append(c.Tags, "us-east-1a")

	switch msgType {
	case NotificationMsg:
		errOut := decodeSNS(body)

		ev := statsd.Event{
			Title:          "Pipeline Error",
			Text:           fmt.Sprintf("A Pipeline error %s occurred", errOut),
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

		// b, err := ioutil.ReadAll(body)
		// if err != nil {
		// 	fmt.Println("Got notif error", err)
		// }
		// fmt.Println("Body is", string(b))
	case SubscriptionMsg:
		if err := snsConfirmSubscription(body); err != nil {
			log.Error("Subscription confirmation failed:", err)
		}
	default:
		log.Error("Unknown message request received:", msgType)
	}

	r.Header.Set("X-Extra", msgType)
}

// Check handles Check comment
func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request, it is alive")
}

func decodeSNS(body io.ReadCloser) log.ErrOut {
	msg := snsMsg{}
	dec := json.NewDecoder(body)
	if err := dec.Decode(&msg); err != nil {
		fmt.Println("got err", err)
	}

	errOut := log.ErrOut{}
	err := json.Unmarshal([]byte(msg.Message), &errOut)

	if err != nil {
		fmt.Println("err is", err)
	}
	fmt.Println("errOut is", errOut)

	return errOut

}
