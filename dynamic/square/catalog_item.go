package square

type CatalogItem struct {
	Name                    string                         `json:"name,omitempty"`
	Description             string                         `json:"description,omitempty"`
	Abbreviation            string                         `json:"abbreviation,omitempty"`
	LabelColor              string                         `json:"label_color,omitempty"`
	AvailableOnline         bool                           `json:"available_online,omitempty"`
	AvailableForPickup      bool                           `json:"available_for_pickup,omitempty"`
	AvailableElectronically bool                           `json:"available_electronically,omitempty"`
	CategoryId              string                         `json:"category_id,omitempty"`
	TaxIds                  []string                       `json:"tax_ids,omitempty"`
	ModifierListInfo        []*CatalogItemModifierListInfo `json:"modifier_list_info,omitempty"`
	Variations              []*CatalogObject               `json:"variations,omitempty"`
	ProductType             string                         `json:"product_type,omitempty"`
	SkipModifierScreen      bool                           `json:"skip_modifier_screen,omitempty"`
}

func (*CatalogItem) isCatalogObjectType() {}
