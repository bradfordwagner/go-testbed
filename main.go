package main

import (
	"context"
	"encoding/json"
	"github.com/bradfordwagner/go-util/log"
	"github.com/hashicorp/vault-client-go"
)

func main() {
	l := log.Log()
	vaultClient, err := vault.New(
		vault.WithEnvironment(),
	)

	healthStatus, err := vaultClient.System.ReadHealthStatus(context.Background())
	if err != nil {
		l.With("error", err).Error("failed to read health status")
		return
	}
	// Convert healthStatus to pretty-printed JSON
	prettyJSON, err := json.MarshalIndent(healthStatus.Data, "", "  ")
	if err != nil {
		l.With("error", err).Error("failed to marshal health status to JSON")
		return
	}

	// Print the status
	l.Infof("Vault health status: \n%s", string(prettyJSON))
}
