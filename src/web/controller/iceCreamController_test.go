package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/context"
	"github.com/iceCream/src/dataAccess"
	"github.com/iceCream/src/model"
)

func TestInsertIceCream(t *testing.T) {

	var jsonStr = []byte(`{
		"name": "Caramel Chocolate Cheesecake",
		"image_closed": "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/caramel-chocolate-cheesecake-truffles-landing.png",
		"image_open": "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/caramel-chocolate-cheesecake-truffles-landing-open.png",
		"description": "Caramel Cheesecake Ice Cream with Graham Cracker-Covered Cheesecake Truffles & Chocolate Cookie Swirls",
		"story": "In your cheesecake dreams, is it like you\u2019re spooning through a world of caramel cheesecake ice cream swirled with chocolate cookies in a wonderland of truffles filled with cheesecake? Hello? You can wake up now\u2026",
		"sourcing_values": ["Non-GMO", "Cage-Free Eggs", "Fairtrade", "Responsibly Sourced Packaging", "Caring Dairy"],
		"ingredients": ["cream", "skim milk", "water", "liquid sugar (sugar", "water)", "sugar", "corn syrup", "canola oil", "cream cheese (pasteurized milk", "cream", "cheese cultures", "salt", "carob bean gum)", "coconut oil", "egg yolks", "wheat flour", "dried cane syrup", "soybean oil", "graham flour", "eggs", "cocoa (processed with alkali)", "natural flavors", "cocoa", "guar gum", "butteroil", "milk protein concentrate", "corn starch", "salt", "soy lecithin", "tapioca starch", "pectin", "caramelized sugar syrup", "baking soda", "molasses", "honey", "carrageenan", "vanilla extract"],
		"allergy_info": "contains milk, eggs, wheat and soy",
		"dietary_certifications": "Kosher",
		"productId": "2191"
	}`)

	ic := model.IceCream{}
	json.Unmarshal(jsonStr, &ic)

	/*	GetAllIceCreams
		GetIceCreamByProductID
		AddIceCreamObj */

	old := dataAccess.AddIceCreamObj
	dataAccess.AddIceCreamObj = func(ic model.IceCream) error {
		return nil
	}
	defer func() {
		dataAccess.AddIceCreamObj = old
	}()

	old2 := dataAccess.GetIceCreamByProductID
	dataAccess.GetIceCreamByProductID = func(productId string) *model.IceCream {
		return &ic
	}
	defer func() {
		dataAccess.GetIceCreamByProductID = old2
	}()

	old3 := dataAccess.GetAllIceCreams
	dataAccess.GetAllIceCreams = func() *[]model.IceCream {
		ics := []model.IceCream{}
		ics = append(ics, ic)
		return &ics
	}
	defer func() {
		dataAccess.GetAllIceCreams = old3
	}()

	t.Run("Add Ice cream Success", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/ic/insert", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("authorization", "Bearer xyz")
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(InsertIceCream)
		handler.ServeHTTP(w, req)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status code %d", resp.StatusCode)
		}

	})

	t.Run("Get all Ice cream Success", func(t *testing.T) {
		req, err := http.NewRequest("Get", "/ic/getall", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("authorization", "Bearer xyz")
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(GetAllIceCreams)
		handler.ServeHTTP(w, req)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status code %d", resp.StatusCode)
		}

	})

	t.Run("Get Ice cream by productID Success", func(t *testing.T) {
		req, err := http.NewRequest("Get", "/ic/get/2191", nil)
		if err != nil {
			t.Fatal(err)
		}

		vars := map[string]string{
			"productID": "2191",
		}
		context.Set(req, 0, vars)

		req.Header.Set("authorization", "Bearer xyz")
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(GetIceCreamByProductID)
		handler.ServeHTTP(w, req)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status code %d", resp.StatusCode)
		}

	})
}
