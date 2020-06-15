package helloworld

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

// the contains handler logic

// Handler should contain fields referencing aws services
// Here we are using *http.Client as a placeholder
type Handler struct {
	client  *http.Client
	baseURL string
}

// NewHandler returns a new Handler
func NewHandler(c *http.Client, bu string) *Handler {
	return &Handler{client: c, baseURL: bu}
}

//Handle returns an APIGatewayResponse
func (h *Handler) Handle(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := h.client.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	defer resp.Body.Close()

	if ip, err := ioutil.ReadAll(resp.Body); err != nil {
		return events.APIGatewayProxyResponse{}, err
	} else if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	} else {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Hello, %v", string(ip)),
			StatusCode: 200,
		}, nil
	}

}
