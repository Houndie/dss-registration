package square

type OrderServiceChargeCalculationPhase string

const (
	OrderServiceChargeCalculationPhaseSubtotalPhase OrderServiceChargeCalculationPhase = "SUBTOTAL_PHASE"
	OrderServiceChargeCalculationPhaseTotalPhase    OrderServiceChargeCalculationPhase = "TOTAL_PHASE"
)

type OrderServiceCharge struct {
	Uid              string                             `json:"uid,omitempty"`
	Name             string                             `json:"name,omitempty"`
	CatalogObjectId  string                             `json:"catalog_object_id,omitempty"`
	Percentage       string                             `json:"percentage,omitempty"`
	AmountMoney      *Money                             `json:"amount_money,omitempty"`
	AppliedMoney     *Money                             `json:"applied_money,omitempty"`
	TotalMoney       *Money                             `json:"total_money,omitempty"`
	TotalTaxMoney    *Money                             `json:"total_tax_money,omitempty"`
	CalculationPhase OrderServiceChargeCalculationPhase `json:"calculation_phase,omitempty"`
	Taxable          bool                               `json:"taxable,omitempty"`
	Taxes            []*OrderLineItemTax                `json:"taxes,omitempty"`
}
