package square

type CatalogMeasurementUnit struct {
	MeasurementUnit *MeasurementUnit `json:"measurement_unit"`
	Precision       int              `json:"precision"`
}

type MeasurementUnit struct {
	CustomUnit  *MeasurementUnitCustom `json:"custom_unit"`
	AreaUnit    string                 `json:"area_unit"`
	LengthUnit  string                 `json:"length_unit"`
	VolumeUnit  string                 `json:"volume_unit"`
	WeightUnit  string                 `json:"weight_unit"`
	GenericUnit string                 `json:"generic_unit"`
}

type MeasurementUnitCustom struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

func (*CatalogMeasurementUnit) isCatalogObjectType() {}
