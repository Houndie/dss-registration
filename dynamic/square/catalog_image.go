package square

type CatalogImage struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Caption string `json:"caption"`
}

func (*CatalogImage) isCatalogObjectType() {}
