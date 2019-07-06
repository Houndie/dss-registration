package square

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type CatalogObjectType string

const (
	CatalogObjectTypeItem            CatalogObjectType = "ITEM"
	CatalogObjectTypeItemVariation   CatalogObjectType = "ITEM_VARIATION"
	CatalogObjectTypeModifier        CatalogObjectType = "MODIFIER"
	CatalogObjectTypeModifierList    CatalogObjectType = "MODIFIER_LIST"
	CatalogObjectTypeCategory        CatalogObjectType = "CATEGORY"
	CatalogObjectTypeDiscount        CatalogObjectType = "DISCOUNT"
	CatalogObjectTypeTax             CatalogObjectType = "TAX"
	CatalogObjectTypeImage           CatalogObjectType = "IMAGE"
	CatalogObjectTypeMeasurementUnit CatalogObjectType = "MEASUREMENT_UNIT"
)

type catalogObject struct {
	Type                  CatalogObjectType       `json:"type,omitempty"`
	Id                    string                  `json:"id,omitempty"`
	UpdatedAt             string                  `json:"updated_at,omitempty"`
	Version               int                     `json:"version,omitempty"`
	IsDeleted             bool                    `json:"is_deleted,omitempty"`
	CatalogV1Ids          []*CatalogV1Id          `json:"catalog_v1_ids,omitempty"`
	PresentAtAllLocations bool                    `json:"present_at_all_locations,omitempty"`
	PresentAtLocationIds  []string                `json:"present_at_location_ids,omitempty"`
	AbsentAtLocationIds   []string                `json:"absent_at_location_ids,omitempty"`
	ImageId               string                  `json:"image_id,omitempty"`
	ItemData              *CatalogItem            `json:"item_data,omitempty"`
	CategoryData          *CatalogCategory        `json:"category_data,omitempty"`
	ItemVariationData     *CatalogItemVariation   `json:"item_variation_data,omitempty"`
	TaxData               *CatalogTax             `json:"tax_data,omitempty"`
	DiscountData          *CatalogDiscount        `json:"discount_data,omitempty"`
	ModifierListData      *CatalogModifierList    `json:"modifier_list_data,omitempty"`
	ModifierData          *CatalogModifier        `json:"modifier_data,omitempty"`
	ImageData             *CatalogImage           `json:"image_data,omitempty"`
	MeasurementUnitData   *CatalogMeasurementUnit `json:"catalog_measurement_unit,omitempty"`
}

type catalogObjectType interface {
	isCatalogObjectType()
}

type CatalogObject struct {
	Id                    string
	UpdatedAt             string
	Version               int
	IsDeleted             bool
	CatalogV1Ids          []*CatalogV1Id
	PresentAtAllLocations bool
	PresentAtLocationIds  []string
	AbsentAtLocationIds   []string
	ImageId               string
	CatalogObjectType     catalogObjectType
}

func (c *CatalogObject) MarshalJSON() ([]byte, error) {
	cJson := catalogObject{
		Id:                    c.Id,
		UpdatedAt:             c.UpdatedAt,
		Version:               c.Version,
		IsDeleted:             c.IsDeleted,
		CatalogV1Ids:          c.CatalogV1Ids,
		PresentAtAllLocations: c.PresentAtAllLocations,
		PresentAtLocationIds:  c.PresentAtLocationIds,
		AbsentAtLocationIds:   c.AbsentAtLocationIds,
		ImageId:               c.ImageId,
	}
	switch t := c.CatalogObjectType.(type) {
	case *CatalogItem:
		cJson.ItemData = t
		cJson.Type = CatalogObjectTypeItem
	case *CatalogCategory:
		cJson.CategoryData = t
		cJson.Type = CatalogObjectTypeCategory
	case *CatalogItemVariation:
		cJson.ItemVariationData = t
		cJson.Type = CatalogObjectTypeItemVariation
	case *CatalogTax:
		cJson.TaxData = t
		cJson.Type = CatalogObjectTypeTax
	case *CatalogDiscount:
		cJson.DiscountData = t
		cJson.Type = CatalogObjectTypeDiscount
	case *CatalogModifierList:
		cJson.ModifierListData = t
		cJson.Type = CatalogObjectTypeModifierList
	case *CatalogModifier:
		cJson.ModifierData = t
		cJson.Type = CatalogObjectTypeModifier
	case *CatalogImage:
		cJson.ImageData = t
		cJson.Type = CatalogObjectTypeImage
	case *CatalogMeasurementUnit:
		cJson.MeasurementUnitData = t
		cJson.Type = CatalogObjectTypeMeasurementUnit
	default:
		return nil, errors.New("Found unknown catalog object data type")
	}
	json, err := json.Marshal(&cJson)
	return json, errors.Wrap(err, "Error marshaling json catalog object")
}

func (c *CatalogObject) UnmarshalJSON(data []byte) error {
	cJson := &catalogObject{}
	err := json.Unmarshal(data, &cJson)
	if err != nil {
		return errors.Wrap(err, "Error unmarshaling catalog object")
	}
	c.Id = cJson.Id
	c.UpdatedAt = cJson.UpdatedAt
	c.Version = cJson.Version
	c.IsDeleted = cJson.IsDeleted
	c.CatalogV1Ids = cJson.CatalogV1Ids
	c.PresentAtAllLocations = cJson.PresentAtAllLocations
	c.PresentAtLocationIds = cJson.PresentAtLocationIds
	c.AbsentAtLocationIds = cJson.AbsentAtLocationIds
	c.ImageId = cJson.ImageId

	switch cJson.Type {
	case CatalogObjectTypeItem:
		c.CatalogObjectType = cJson.ItemData
	case CatalogObjectTypeCategory:
		c.CatalogObjectType = cJson.CategoryData
	case CatalogObjectTypeItemVariation:
		c.CatalogObjectType = cJson.ItemVariationData
	case CatalogObjectTypeTax:
		c.CatalogObjectType = cJson.TaxData
	case CatalogObjectTypeDiscount:
		c.CatalogObjectType = cJson.DiscountData
	case CatalogObjectTypeModifierList:
		c.CatalogObjectType = cJson.ModifierListData
	case CatalogObjectTypeModifier:
		c.CatalogObjectType = cJson.ModifierData
	case CatalogObjectTypeImage:
		c.CatalogObjectType = cJson.ImageData
	case CatalogObjectTypeMeasurementUnit:
		c.CatalogObjectType = cJson.MeasurementUnitData
	default:
		return fmt.Errorf("Found unknown catalog object type %s", cJson.Type)
	}
	return nil
}
