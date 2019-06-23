package square

type OrderServiceChargeCalculationPhase string

const (
	OrderServiceChargeCalculationPhaseSubtotalPhase OrderServiceChargeCalculationPhase = "SUBTOTAL_PHASE"
	OrderServiceChargeCalculationPhaseTotalPhase    OrderServiceChargeCalculationPhase = "TOTAL_PHASE"
)

type OrderServiceCharge struct {
	Uid              string                             `json:"uid"`
	Name             string                             `json:"name"`
	CatalogObjectId  string                             `json:"catalog_object_id"`
	Percentage       string                             `json:"percentage"`
	AmountMoney      *Money                             `json:"amount_money"`
	AppliedMoney     *Money                             `json:"applied_money"`
	TotalMoney       *Money                             `json:"total_money"`
	TotalTaxMoney    *Money                             `json:"total_tax_money"`
	CalculationPhase OrderServiceChargeCalculationPhase `json:"calculation_phase"`
	Taxable          bool                               `json:"taxable"`
	Taxes            []*OrderLineItemTax                `json:"taxes"`
}
