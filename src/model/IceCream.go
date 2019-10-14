package model

//IceCream model
type IceCream struct {
	Name                  string   `json:"name,omitempty"`
	ImageClosed           string   `json:"image_closed,omitempty"`
	ImageOpen             string   `json:"image_open,omitempty"`
	Description           string   `json:"description,omitempty"`
	Story                 string   `json:"story,omitempty"`
	SourcingValues        []string `json:"sourcing_values,omitempty"`
	Ingredients           []string `json:"ingredients,omitempty"`
	AllergyInfo           string   `json:"allergy_info,omitempty"`
	DietaryCertifications string   `json:"dietary_certifications,omitempty"`
	ProductID             string   `json:"productId,omitempty"`
}

//NewIceCream gives a new iceCream instance
func NewIceCream(name string, image_open string, image_closed string, story string, description string, allergy_info string, sourcing_values []string, ingredients []string, dietary_certifications string, productId string) *IceCream {

	iceCream := IceCream{Name: name, ImageClosed: image_closed, ImageOpen: image_open, Description: description, Story: story, SourcingValues: sourcing_values, Ingredients: ingredients, AllergyInfo: allergy_info, DietaryCertifications: dietary_certifications, ProductID: productId}
	return &iceCream
}
