package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myJSON struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"datValue"`
	ObjectValue     *myObject `json:"objectValue"`
	NullStringValue *string   `json:"nullStringValue,omitempty"`
	NullIntValue    *int      `json:"nullIntValue"`
	EmptyString     string    `json:"emptyString,omitempty"`
}

type myObject struct {
	ArrayValue []int `json:"arrayValue"`
}

func main() {
	jsonData := `
		{
			"intValue": 1234,
			"boolValue": true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null,
			"extraValue":4321
		}
	`

	var data *myJSON
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json struct: %#v\n", data)
	fmt.Printf("dateValue: %#v\n", data.DateValue)
	fmt.Printf("objectValue: %#v\n", data.ObjectValue)

	//fmt.Printf("json map: %v\n", data)

	// rawDateValue, ok := data["dateValue"]
	// if !ok {
	// 	fmt.Printf("dateValue does not exist\n")
	// 	return
	// }
	// dateValue, ok := rawDateValue.(string)
	// if !ok {
	// 	fmt.Printf("dateValue is note a string\n")
	// 	return
	// }
	// fmt.Printf("date value: %s\n", dateValue)

	// otherInt := 4321

	// data := &myJSON{
	// 	IntValue:    1234,
	// 	BoolValue:   true,
	// 	StringValue: "hello!",
	// 	DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
	// 	ObjectValue: &myObject{
	// 		ArrayValue: []int{1, 2, 3, 4},
	// 	},
	// 	NullStringValue: nil,
	// 	NullIntValue:    &otherInt,
	// }

	// jsondata, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Printf("could not marshal json: %s\n", err)
	// }

	// fmt.Printf("json data: %s\n", jsondata)
}
