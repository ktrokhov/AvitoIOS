package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Struct struct {
	Params []Params `json:"params"`
}
type Values struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Params []Params `json:"params,omitempty"`
}
type Params struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Value  string   `json:"value"`
	Values []Values `json:"values,omitempty"`
}

type Draft struct {
	Values []struct {
		ID    int         `json:"id"`
		Value interface{} `json:"value"`
	} `json:"values"`
}

func OpenJson1Struct() *Draft {
	jsonFile, err := os.Open("Draft_value.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	val := &Draft{}
	json.Unmarshal(byteValue, val)
	return val
}


func OpenJson2Struct() *Struct {
	jsonFile, err := os.Open("Structure.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	val := &Struct{}
	json.Unmarshal(byteValue, val)

	return val
}

func newFunc(dstID int, srcVal interface{}, params []Params) bool {
	for i, record := range params {
		if record.ID == dstID {
			if len(record.Values) == 0 {
				params[i].Value = srcVal.(string)
			}
			for _, value := range record.Values {
				if value.ID == int(srcVal.(float64)) {
					params[i].Value = value.Title
					break
				}
			}
			return true
		}
		for j := range params[i].Values {
			if newFunc(dstID, srcVal, params[i].Values[j].Params) {

			}
		}
	}
	return false
}

func main() {
	draft := OpenJson1Struct()
	struc := OpenJson2Struct()
	for _, v := range draft.Values {
		fmt.Printf("%v\n", newFunc(v.ID, v.Value, struc.Params))
	}
	b, _ := json.MarshalIndent(struc, "", "\t")
	ioutil.WriteFile("Structure_with_values.json", b, 0644)
}
