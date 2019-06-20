package square

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const (
	CatalogObjectItemType            = "ITEM"
	CatalogObjectItemVariationType   = "ITEM_VARIATION"
	CatalogObjectModifierType        = "MODIFIER"
	CatalogObjectModifierListType    = "MODIFIER_LIST"
	CatalogObjectCategoryType        = "CATEGORY"
	CatalogObjectDiscountType        = "DISCOUNT"
	CatalogObjectTaxType             = "TAX"
	CatalogObjectImageType           = "IMAGE"
	CatalogObjectMeasurementUnitType = "MEASUREMENT_UNIT"
)

type catalogObject struct {
	Type                  string                  `json:"type"`
	Id                    string                  `json:"id"`
	UpdatedAt             string                  `json:"updated_at"`
	Version               int                     `json:"version"`
	IsDeleted             bool                    `json:"is_deleted"`
	CatalogV1Ids          []*CatalogV1Id          `json:"catalog_v1_ids"`
	PresentAtAllLocations bool                    `json:"present_at_all_locations"`
	PresentAtLocationIds  []string                `json:"present_at_location_ids"`
	AbsentAtLocationIds   []string                `json:"absent_at_location_ids"`
	ImageId               string                  `json:"image_id"`
	ItemData              *CatalogItem            `json:"item_data"`
	CategoryData          *CatalogCategory        `json:"category_data"`
	ItemVariationData     *CatalogItemVariation   `json:"item_variation_data"`
	TaxData               *CatalogTax             `json:"tax_data"`
	DiscountData          *CatalogDiscount        `json:"discount_data"`
	ModifierListData      *CatalogModifierList    `json:"modifier_list_data"`
	ModifierData          *CatalogModifier        `json:"modifier_data"`
	ImageData             *CatalogImage           `json:"image_data"`
	MeasurementUnitData   *CatalogMeasurementUnit `json:"catalog_measurement_unit"`
}

type CatalogObjectType interface {
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
	CatalogObjectType     CatalogObjectType
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
		cJson.Type = CatalogObjectItemType
	case *CatalogCategory:
		cJson.CategoryData = t
		cJson.Type = CatalogObjectCategoryType
	case *CatalogItemVariation:
		cJson.ItemVariationData = t
		cJson.Type = CatalogObjectItemVariationType
	case *CatalogTax:
		cJson.TaxData = t
		cJson.Type = CatalogObjectTaxType
	case *CatalogDiscount:
		cJson.DiscountData = t
		cJson.Type = CatalogObjectDiscountType
	case *CatalogModifierList:
		cJson.ModifierListData = t
		cJson.Type = CatalogObjectModifierListType
	case *CatalogModifier:
		cJson.ModifierData = t
		cJson.Type = CatalogObjectModifierType
	case *CatalogImage:
		cJson.ImageData = t
		cJson.Type = CatalogObjectImageType
	case *CatalogMeasurementUnit:
		cJson.MeasurementUnitData = t
		cJson.Type = CatalogObjectMeasurementUnitType
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
	case CatalogObjectItemType:
		c.CatalogObjectType = cJson.ItemData
	case CatalogObjectCategoryType:
		c.CatalogObjectType = cJson.CategoryData
	case CatalogObjectItemVariationType:
		c.CatalogObjectType = cJson.ItemVariationData
	case CatalogObjectTaxType:
		c.CatalogObjectType = cJson.TaxData
	case CatalogObjectDiscountType:
		c.CatalogObjectType = cJson.DiscountData
	case CatalogObjectModifierListType:
		c.CatalogObjectType = cJson.ModifierListData
	case CatalogObjectModifierType:
		c.CatalogObjectType = cJson.ModifierData
	case CatalogObjectImageType:
		c.CatalogObjectType = cJson.ImageData
	case CatalogObjectMeasurementUnitType:
		c.CatalogObjectType = cJson.MeasurementUnitData
	default:
		return fmt.Errorf("Found unknown catalog object type %s", cJson.Type)
	}
	return nil
}
