package square

type OrderReturnTax struct {
	Uid             string                `json:"uid,omitempty"`
	SourceTaxUid    string                `json:"source_tax_uid,omitempty"`
	CatalogObjectId string                `json:"catalog_object_id,omitempty"`
	Name            string                `json:"name,omitempty"`
	Type            OrderLineItemTaxType  `json:"type,omitempty"`
	Percentage      string                `json:"percentage,omitempty"`
	AppliedMoney    *Money                `json:"applied_money,omitempty"`
	Scope           OrderLineItemTaxScope `json:"scope,omitempty"`
}
