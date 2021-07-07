package websocket

// WebhookEndpoint contains properties about the fake "endpoint" used to
// format the webhook event.
type WebhookEndpoint struct {
	APIVersion *string `json:"api_version"`
}

// WebhookEvent represents incoming webhook event messages sent by Stripe.
type WebhookEvent struct {
	Endpoint              WebhookEndpoint   `json:"endpoint"`
	EventPayload          string            `json:"event_payload"`
	HTTPHeaders           map[string]string `json:"http_headers"`
	Type                  string            `json:"type"`
	WebhookConversationID string            `json:"webhook_conversation_id"`
	WebhookID             string            `json:"webhook_id"`
}

// WebhookResponse represents outgoing webhook response messages sent to
// Stripe.
type WebhookResponse struct {
	ForwardURL            string            `json:"forward_url"`
	Status                int               `json:"status"`
	HTTPHeaders           map[string]string `json:"http_headers"`
	Body                  string            `json:"body"`
	Type                  string            `json:"type"`
	WebhookConversationID string            `json:"webhook_conversation_id"`
	WebhookID             string            `json:"webhook_id"`
}

// NewWebhookResponse returns a new webhookResponse message.
func NewWebhookResponse(webhookID, webhookConversationID, forwardURL string, status int, body string, headers map[string]string) *OutgoingMessage {
	return &OutgoingMessage{
		WebhookResponse: &WebhookResponse{
			WebhookID:             webhookID,
			WebhookConversationID: webhookConversationID,
			ForwardURL:            forwardURL,
			Status:                status,
			Body:                  body,
			HTTPHeaders:           headers,
			Type:                  "webhook_response",
		},
	}
}
go get -u github.com/stripe/stripe-go
package main

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"

  "github.com/stripe/stripe-go"
  "github.com/stripe/stripe-go/paymentintent"
)

func main() {
  // This is your real test secret API key.
  stripe.Key = "sk_test_51J75cDIwrPLAhbrNKRfE4l0ALowryM7JQbQm0maoJqzahbDMjmqzYylXZQXSY0SShVUFgSugn2HFRNf8McG6SIwY00Ope27Tap"

  http.HandleFunc("/webhook", handleWebhook)
  addr := "localhost:4242"
  log.Printf("Listening on %s", addr)
  log.Fatal(http.ListenAndServe(addr, nil))
}

func handleWebhook(w http.ResponseWriter, req *http.Request) {
  const MaxBodyBytes = int64(65536)
  req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
  payload, err := ioutil.ReadAll(req.Body)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
    w.WriteHeader(http.StatusServiceUnavailable)
    return
  }

  event := stripe.Event{}

  if err := json.Unmarshal(payload, &event); err != nil {
    fmt.Fprintf(os.Stderr, "⚠️  Webhook error while parsing basic request. %v\n", err.Error())
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // Unmarshal the event data into an appropriate struct depending on its Type
  switch event.Type {
  case "payment_intent.succeeded":
    var paymentIntent stripe.PaymentIntent
    err := json.Unmarshal(event.Data.Raw, &paymentIntent)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
      w.WriteHeader(http.StatusBadRequest)
      return
    }
    log.Printf("Successful payment for %d.", paymentIntent.Amount)
    // Then define and call a func to handle the successful payment intent.
    // handlePaymentIntentSucceeded(paymentIntent)
  case "payment_method.attached":
    var paymentMethod stripe.PaymentMethod
    err := json.Unmarshal(event.Data.Raw, &paymentMethod)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
      w.WriteHeader(http.StatusBadRequest)
      return
    }
    // Then define and call a func to handle the successful attachment of a PaymentMethod.
    // handlePaymentMethodAttached(paymentMethod)
  default:
    fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
  }

  w.WriteHeader(http.StatusOK)
}
