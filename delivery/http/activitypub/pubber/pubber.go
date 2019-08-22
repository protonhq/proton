package pubber

import (
	"net/http"
	"time"

	"github.com/go-fed/activity/deliverer"
	"github.com/go-fed/activity/pub"
	"github.com/spf13/viper"
	"golang.org/x/time/rate"
)

// Clock determines the time.
type Clock struct{}

func (c *Clock) Now() time.Time {
	return time.Now()
}

// HttpClient sends http requests.
type HttpClient struct{}

func (m *HttpClient) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}

func CreatePubber() (pub.Pubber, pub.HandlerFunc) {
	// pub.Clock
	clock := &Clock{}

	// pub.SocialFederateApplication
	app := &Application{}

	// pub.Callbacker
	socialCallbacker := &SocialCallbacker{}
	federatedCallbacker := &FederatedCallbacker{}

	// pub.Deliverer
	deliverer := createDeliverer()

	// pub.HttpClient
	httpClient := &HttpClient{}

	userAgent := viper.GetString("pubber.userAgent")
	maxDeliveryDepth := viper.GetInt("pubber.maxDeliveryDepth")
	maxInboxForwardingDepth := viper.GetInt("pubber.maxInboxForwardingDepth")

	pubber := pub.NewPubber(clock, app, socialCallbacker, federatedCallbacker, deliverer, httpClient, userAgent, maxDeliveryDepth, maxInboxForwardingDepth)
	asHandler := pub.ServeActivityPubObject(app, clock)
	return pubber, asHandler
}

func createDeliverer() pub.Deliverer {
	opts := deliverer.DeliveryOptions{}
	opts.InitialRetryTime = 60 * time.Second
	opts.MaximumRetryTime = time.Hour
	opts.BackoffFactor = 1.5
	opts.MaxRetries = 30
	opts.RateLimit = rate.NewLimiter(0.5, 30)

	dpool := deliverer.NewDelivererPool(opts)

	return dpool
}
