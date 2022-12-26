package main

import (
	"encoding/json"
	"io/ioutil"
)

type Recipie struct {
	Thinly_Sliced_Peeled_Apples string
	sugar                       string
	flour                       string
	cinamon                     string
	nutmeg                      string
	lemon_juice                 string
}

func main() {
	jsonAppleRecipie := `{
	"Thinly Sliced Peeled Apples" : "6 Cups",
	"sugar" : "3/4 cup",
	"flour" : "2 tablespoons",
	"cinamon" : "1/4 teaspoon",
	"nutmeg" : "1/8 tablesppon",
	"lemon juice" : "1 tablespoon"}`

	newMap := make(map[string]string)
	err2 := json.Unmarshal([]byte(jsonAppleRecipie), &newMap)
	if err2 != nil {
		panic(err2)
	}

	file, _ := json.MarshalIndent(newMap, "", "")

	_ = ioutil.WriteFile("recipie.json", file, 0644)

}
