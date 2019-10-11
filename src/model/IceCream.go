package model

//IceCream model
type IceCream struct {
	Name                  string   `json:"name"`
	ImageClosed           string   `json:"image_closed"`
	ImageOpen             string   `json:"image_open"`
	Description           string   `json:"description"`
	Story                 string   `json:"story"`
	SourcingValues        []string `json:"sourcing_values"`
	Ingredients           []string `json:"ingredients"`
	AllergyInfo           string   `json:"allergy_info"`
	DietaryCertifications string   `json:"dietary_certifications"`
	ProductID             string   `json:"productId"`
}

//NewIceCream gives a new iceCream instance
func NewIceCream(name string, image_open string, image_closed string, story string, description string, allergy_info string, sourcing_values []string, ingredients []string, dietary_certifications string, productId string) *IceCream {

	iceCream := IceCream{Name: name, ImageClosed: image_closed, ImageOpen: image_open, Description: description, Story: story, SourcingValues: sourcing_values, Ingredients: ingredients, AllergyInfo: allergy_info, DietaryCertifications: dietary_certifications, ProductID: productId}
	return &iceCream
}
