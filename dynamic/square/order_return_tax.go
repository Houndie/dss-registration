package square

type OrderReturnTax struct {
	Uid             string                `json:"uid"`
	SourceTaxUid    string                `json:"source_tax_uid"`
	CatalogObjectId string                `json:"catalog_object_id"`
	Name            string                `json:"name"`
	Type            OrderLineItemTaxType  `json:"type"`
	Percentage      string                `json:"percentage"`
	AppliedMoney    *Money                `json:"applied_money"`
	Scope           OrderLineItemTaxScope `json:"scope"`
}
