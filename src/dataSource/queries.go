package dataSource

const (
	//InsertIceCream in DB
	InsertIceCream = "INSERT INTO IceCreams (Name,ImageClosed,ImageOpen,Description,Story, SourcingValues, Ingredients, AllergyInfo, DietaryCertifications, ProductID) VALUES (?,?,?,?,?,?,?,?,?,?)"

	// GetIceCreamByProductID specific
	GetIceCreamByProductID = "Select * from IceCreams where ProductID = ?"

	//GetAllIceCreams in DB
	GetAllIceCreams = "Select * from IceCreams"

	//UpdateIceCreaByProductID = "Select c.Name, c.ManufacturingYear, c.Color, c.CarNumber from Cars c where c.Colour = ?"
)
