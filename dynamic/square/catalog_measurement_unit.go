package square

type CatalogMeasurementUnit struct {
	MeasurementUnit *MeasurementUnit `json:"measurement_unit,omitempty"`
	Precision       int              `json:"precision,omitempty"`
}

type MeasurementUnit struct {
	CustomUnit  *MeasurementUnitCustom `json:"custom_unit,omitempty"`
	AreaUnit    string                 `json:"area_unit,omitempty"`
	LengthUnit  string                 `json:"length_unit,omitempty"`
	VolumeUnit  string                 `json:"volume_unit,omitempty"`
	WeightUnit  string                 `json:"weight_unit,omitempty"`
	GenericUnit string                 `json:"generic_unit,omitempty"`
}

type MeasurementUnitCustom struct {
	Name         string `json:"name,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
}

func (*CatalogMeasurementUnit) isCatalogObjectType() {}
