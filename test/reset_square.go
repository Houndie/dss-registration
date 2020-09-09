package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/gofrs/uuid"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	apiKey := os.Getenv("SQUARE_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("no square api key found")
	}
	client, err := square.NewClient(apiKey, square.Sandbox, &http.Client{})
	if err != nil {
		return fmt.Errorf("error creating new client: %w", err)
	}

	locations, err := client.ListLocations(context.Background())
	if err != nil {
		return fmt.Errorf("error listing locations: %w", err)
	}

	if len(locations) > 1 {
		return fmt.Errorf("too many locations found: %d", len(locations))
	}

	changes := []*square.InventoryChange{}
	idempotencyKey, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("error generating new idempotency key: %w", err)
	}

	objects := client.ListCatalog(context.Background(), nil)
	for objects.Next() {
		v, ok := objects.Value().CatalogObjectType.(*square.CatalogItem)
		if !ok {
			continue
		}
		switch v.Name {
		case "Mix And Match", "Team Competition", "Solo", "2020 T-Shirt":
			changes = append(changes, &square.InventoryChange{
				Type: &square.InventoryPhysicalCount{
					CatalogObjectID: v.Variations[0].ID,
					Quantity:        "1000",
				},
			})
		}
	}

	if err := objects.Error(); err != nil {
		return fmt.Errorf("error fetching catalog: %w", err)
	}

	_, err = client.BatchChangeInventory(context.Background(), idempotencyKey.String(), changes)
	if err != nil {
		return fmt.Errorf("error changing inventory")
	}
	return nil
}
