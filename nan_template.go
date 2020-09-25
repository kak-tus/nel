package nan

// imports

// nullTemplateType - generated for nullTemplateType
//easyjson:skip
type nullTemplateType struct {
	NullTemplateValue initialTemplateType
	Valid             bool // Valid is true if Value is not NULL
}

func (n nullTemplateType) IsValid() bool {
	return n.Valid
}
