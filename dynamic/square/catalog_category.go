package square

type CatalogCategory struct {
	Name string `json:"name"`
}

func (*CatalogCategory) isCatalogObjectType() {}
