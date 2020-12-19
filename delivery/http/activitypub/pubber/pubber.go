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

// Now determines server time
func (c *Clock) Now() time.Time {
	return time.Now()
}

// HTTPClient sends http requests.
type HTTPClient struct{}

// Do a request
func (m *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}

// CreatePubber - create activitypub pubber
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

	// pub.HTTPClient
	httpClient := &HTTPClient{}

	userAgent := viper.GetString("pubber.user_agent")
	maxDeliveryDepth := viper.GetInt("pubber.max_delivery_depth")
	maxInboxForwardingDepth := viper.GetInt("pubber.max_inbox_forwarding_depth")

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
