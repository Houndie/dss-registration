package square

type CatalogTax struct {
	Name                   string `json:"name"`
	CalculationPhase       string `json:"calculation_phase"`
	InclusionType          string `json:"inclusion_type"`
	Percentage             string `json:"percentage"`
	AppliesToCustomAmounts bool   `json:"applies_to_custom_amounts"`
	Enabled                bool   `json:"enabled"`
}

func (*CatalogTax) isCatalogObjectType() {}
