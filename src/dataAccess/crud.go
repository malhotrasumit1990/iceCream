package dataAccess

import (
	"fmt"
	"log"

	"github.com/iceCream/src/dataSource"
	"github.com/iceCream/src/model"
)

//AddIceCreamObj in DB. This method adds one object at a time.
var AddIceCreamObj = func(iceCream model.IceCream) error {

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(dataSource.InsertIceCream)
	if err != nil {
		log.Print(err)
		return err
	}

	svStr := arrayToString(iceCream.SourcingValues)
	ingStr := arrayToString(iceCream.Ingredients)

	_, err = stmt.Exec(iceCream.Name, iceCream.ImageClosed, iceCream.ImageOpen, iceCream.Description, iceCream.Story, svStr, ingStr, iceCream.AllergyInfo, iceCream.DietaryCertifications, iceCream.ProductID)
	if err != nil {
		log.Print(err)
		return err
	}

	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil

}

//GetAllIceCreams gets all rows present in IceCream DB.
var GetAllIceCreams = func() *[]model.IceCream {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in dataAccess.getByName", r)
		}
	}()

	rows, err := db.Query(dataSource.GetAllIceCreams)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	iceCreams := []model.IceCream{}

	for rows.Next() {
		var svStr string
		var ingStr string

		iceCream := model.IceCream{}
		err := rows.Scan(&iceCream.Name, &iceCream.ImageClosed, &iceCream.ImageOpen, &iceCream.Description, &iceCream.Story, &svStr, &ingStr, &iceCream.AllergyInfo, &iceCream.DietaryCertifications, &iceCream.ProductID)
		if err != nil {
			log.Print(err)
			return nil
		}
		iceCream.SourcingValues = stringToArray(svStr)
		iceCream.Ingredients = stringToArray(ingStr)

		iceCreams = append(iceCreams, iceCream)

	}
	return &iceCreams

}

//GetIceCreamByProductID return one IceCream object with specific product Id
var GetIceCreamByProductID = func(productID string) *model.IceCream {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in dataAccess.GetIceCreamByProductID", r)
		}
	}()

	rows, err := db.Query(dataSource.GetIceCreamByProductID, productID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	iceCream := model.IceCream{}
	for rows.Next() {
		var svStr string
		var ingStr string

		err := rows.Scan(&iceCream.Name, &iceCream.ImageClosed, &iceCream.ImageOpen, &iceCream.Description, &iceCream.Story, &svStr, &ingStr, &iceCream.AllergyInfo, &iceCream.DietaryCertifications, &iceCream.ProductID)
		if err != nil {
			log.Print(err)
			return nil
		}
		iceCream.SourcingValues = stringToArray(svStr)
		iceCream.Ingredients = stringToArray(ingStr)
	}

	return &iceCream
}
