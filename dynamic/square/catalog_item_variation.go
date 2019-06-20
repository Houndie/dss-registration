package square

type CatalogItemVariation struct {
	ItemId                   string                            `json:"item_id"`
	Name                     string                            `json:"name"`
	Sku                      string                            `json:"sku"`
	Upc                      string                            `json:"string"`
	Ordinal                  int                               `json:"ordinal"`
	PricingType              string                            `json:"pricing_type"`
	PriceMoney               *Money                            `json:"price_money"`
	LocationOverrides        []*ItemVariationLocationOverrides `json:"location_overrides"`
	TrackInventory           bool                              `json:"track_inventory"`
	InventoryAlertType       string                            `json:"inventory_alert_type"`
	InventoryAlertThreshold  int                               `json:"inventory_alert_threshold"`
	UserData                 string                            `json:"user_data"`
	ServiceDuration          int                               `json:"service_duration"`
	CatalogMeasurementUnitId string                            `json:"catalog_measurement_unit_id"`
}

type ItemVariationLocationOverrides struct {
	LocationId              string `json:"location_id"`
	PriceMoney              *Money `json:"price_money"`
	PricingType             string `json:"pricing_type"`
	TrackInventory          bool   `json:"track_inventory"`
	InventoryAlertType      string `json:"inventory_alert_type"`
	InventoryAlertThreshold int    `json:"inventory_alert_threshold"`
}

func (*CatalogItemVariation) isCatalogObjectType() {}
