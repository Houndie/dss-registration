package populate

import "fmt"

type ErrMissingCatalogItem struct {
	Name string
}

func (e *ErrMissingCatalogItem) Error() string {
	return fmt.Sprintf("expected catalog item with name %s was not found", e.Name)
}

type ErrUnxpectedVariationCount struct {
	Name  string
	Count int
}

func (e *ErrUnxpectedVariationCount) Error() string {
	return fmt.Sprintf("expected catalog item %s to only have one variation, instead found %d", e.Name, e.Count)
}

type ErrMissingVariation struct {
	Name          string
	VariationName string
}

func (e *ErrMissingVariation) Error() string {
	return fmt.Sprintf("expected catalog item %s to contain a variation named %s, but it was not found", e.Name, e.VariationName)
}
