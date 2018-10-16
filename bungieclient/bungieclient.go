package bungieclient

import (
	"net/http"

	"github.com/dghubble/sling"
)

const baseURL = "https://bungie.net/Platform/"

type BungieClient struct{
	sling *sling.Sling
}

func NewBungieClient(client *http.Client) *BungieClient {
	return &BungieClient{
		sling: sling.New().Client(client).Base(baseURL).Set("X-API-Key", "lolol"),
	}
}