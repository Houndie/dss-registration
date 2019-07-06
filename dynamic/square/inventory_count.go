package square

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type inventoryCount struct {
	CatalogObjectId   string `json:"catalog_object_id,omitempty"`
	CatalogObjectType string `json:"catalog_object_type,omitempty"`
	State             string `json:"state,omitempty"`
	LocationId        string `json:"location_id,omitempty"`
	Quantity          string `json:"quantity,omitempty"`
	CalculatedAt      string `json:"calculated_at,omitempty"`
}

type InventoryCount struct {
	CatalogObjectId   string
	CatalogObjectType string
	State             string
	LocationId        string
	Quantity          float64
	CalculatedAt      string
}

func (i *InventoryCount) MarshalJSON() ([]byte, error) {
	iJson := &inventoryCount{
		CatalogObjectId:   i.CatalogObjectId,
		CatalogObjectType: i.CatalogObjectType,
		State:             i.State,
		LocationId:        i.LocationId,
		Quantity:          fmt.Sprintf("%v", i.Quantity),
		CalculatedAt:      i.CalculatedAt,
	}
	bytes, err := json.Marshal(&iJson)
	return bytes, errors.Wrap(err, "Error marshaling InventoryCount")
}

func (i *InventoryCount) UnmarshalJSON(data []byte) error {
	iJson := &inventoryCount{}
	err := json.Unmarshal(data, &iJson)
	if err != nil {
		return errors.Wrap(err, "Error unmarshaling InventoryCount")
	}
	quantity, err := strconv.ParseFloat(iJson.Quantity, 64)
	if err != nil {
		return errors.Wrap(err, "error parsing InventoryCount.Quantity as float")
	}
	i.CatalogObjectId = iJson.CatalogObjectId
	i.CatalogObjectType = iJson.CatalogObjectType
	i.State = iJson.State
	i.LocationId = iJson.LocationId
	i.Quantity = quantity
	i.CalculatedAt = iJson.CalculatedAt
	return nil
}
