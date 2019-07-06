package square

type CatalogItemVariation struct {
	ItemId                   string                            `json:"item_id,omitempty"`
	Name                     string                            `json:"name,omitempty"`
	Sku                      string                            `json:"sku,omitempty"`
	Upc                      string                            `json:"string,omitempty"`
	Ordinal                  int                               `json:"ordinal,omitempty"`
	PricingType              string                            `json:"pricing_type,omitempty"`
	PriceMoney               *Money                            `json:"price_money,omitempty"`
	LocationOverrides        []*ItemVariationLocationOverrides `json:"location_overrides,omitempty"`
	TrackInventory           bool                              `json:"track_inventory,omitempty"`
	InventoryAlertType       string                            `json:"inventory_alert_type,omitempty"`
	InventoryAlertThreshold  int                               `json:"inventory_alert_threshold,omitempty"`
	UserData                 string                            `json:"user_data,omitempty"`
	ServiceDuration          int                               `json:"service_duration,omitempty"`
	CatalogMeasurementUnitId string                            `json:"catalog_measurement_unit_id,omitempty"`
}

type ItemVariationLocationOverrides struct {
	LocationId              string `json:"location_id,omitempty"`
	PriceMoney              *Money `json:"price_money,omitempty"`
	PricingType             string `json:"pricing_type,omitempty"`
	TrackInventory          bool   `json:"track_inventory,omitempty"`
	InventoryAlertType      string `json:"inventory_alert_type,omitempty"`
	InventoryAlertThreshold int    `json:"inventory_alert_threshold,omitempty"`
}

func (*CatalogItemVariation) isCatalogObjectType() {}
