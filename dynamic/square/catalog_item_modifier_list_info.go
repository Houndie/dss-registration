package square

type CatalogItemModifierListInfo struct {
	ModifierListId       string                     `json:"modifier_list_id"`
	ModifierOverrides    []*CatalogModifierOverride `json:"modifier_overrides"`
	MinSelectedModifiers int                        `json:"min_selected_modifiers"`
	MaxSelectedModifiers int                        `json:"max_selected_modifiers"`
	Enabled              bool                       `json:"enabled"`
}

type CatalogModifierOverride struct {
	ModifierId  string `json:modifier_id`
	OnByDefault bool   `json:on_by_default`
}
