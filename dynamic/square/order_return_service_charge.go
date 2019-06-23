package square

type OrderReturnServiceCharge struct {
	Uid                    string                             `json:"uid"`
	SourceServiceChargeUid string                             `json:"source_service_charge_uid"`
	Name                   string                             `json:"name"`
	CatalogObjectId        string                             `json:"catalog_object_id"`
	Percentage             string                             `json:"percentage"`
	AmountMoney            *Money                             `json:"amount_money"`
	AppliedMoney           *Money                             `json:"applied_money"`
	TotalMoney             *Money                             `json:"total_money"`
	TotalTaxMoney          *Money                             `json:"total_tax_money"`
	CalculationPhase       OrderServiceChargeCalculationPhase `json:"calculation_phase"`
	Taxable                bool                               `json:"taxable"`
	ReturnTaxes            []*OrderReturnTax                  `json:"return_taxes"`
}
