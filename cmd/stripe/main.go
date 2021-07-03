package main

import "github.com/stripe/stripe-cli/pkg/cmd"

func main() {
	cmd.Execute()
}
stripe webhook_endpoints create \
  --url="https://example.com/my/webhook/endpoint" \
  -d "enabled_events[]"="charge.failed" \
  -d "enabled_events[]"="charge.succeeded"
