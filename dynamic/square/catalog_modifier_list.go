package square

type CatalogModifierList struct {
	Name          string           `json:"name"`
	SelectionType string           `json:"selection_type"`
	Modifiers     []*CatalogObject `json:"modifiers"`
}

func (*CatalogModifierList) isCatalogObjectType() {}
