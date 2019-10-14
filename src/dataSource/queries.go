package dataSource

const (
	//InsertIceCream in DB
	InsertIceCream = "INSERT INTO IceCreams (Name,ImageClosed,ImageOpen,Description,Story, SourcingValues, Ingredients, AllergyInfo, DietaryCertifications, ProductID) VALUES (?,?,?,?,?,?,?,?,?,?)"

	// GetIceCreamByProductID specific
	GetIceCreamByProductID = "Select * from IceCreams where ProductID = ?"

	//GetAllIceCreams in DB
	GetAllIceCreams = "Select * from IceCreams"

	CreateTable = `CREATE TABLE IceCreams
	( Name varchar(50) NOT NULL,
	  ImageClosed Text ,
	  ImageOpen Text ,
	  Description Text ,
	  Story Text ,
	  SourcingValues Text ,
	  Ingredients Text ,
	  AllergyInfo Text ,
	  DietaryCertifications Text ,
	  ProductID varchar(50) NOT NULL
	);`

	//UpdateIceCreaByProductID = "Select c.Name, c.ManufacturingYear, c.Color, c.CarNumber from Cars c where c.Colour = ?"
)
