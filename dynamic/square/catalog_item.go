package square

type CatalogItem struct {
	Name                    string                         `json:"name"`
	Description             string                         `json:"description"`
	Abbreviation            string                         `json:"abbreviation"`
	LabelColor              string                         `json:"label_color"`
	AvailableOnline         bool                           `json:"available_online"`
	AvailableForPickup      bool                           `json:"available_for_pickup"`
	AvailableElectronically bool                           `json:"available_electronically"`
	CategoryId              string                         `json:"category_id"`
	TaxIds                  []string                       `json:"tax_ids"`
	ModifierListInfo        []*CatalogItemModifierListInfo `json:"modifier_list_info"`
	Variations              []*CatalogObject               `json:"variations"`
	ProductType             string                         `json:"product_type"`
	SkipModifierScreen      bool                           `json:"skip_modifier_screen"`
}

func (*CatalogItem) isCatalogObjectType() {}
