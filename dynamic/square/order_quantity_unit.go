package square

type OrderQuantityUnit struct {
	MeasurementUnit *MeasurementUnit `json:"measurement_unit"`
	Precision       int              `json:"precision"`
}
