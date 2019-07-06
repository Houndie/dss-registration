package square

type CatalogCategory struct {
	Name string `json:"name,omitempty"`
}

func (*CatalogCategory) isCatalogObjectType() {}
