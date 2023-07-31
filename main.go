package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func main() {
	data := []byte(`{
		"employees":{
		   "protected":false,
		   "address":{
			  "street":"22 Saint-Lazare",
			  "postalCode":"75003",
			  "city":"Paris",
			  "countryCode":"FRA",
			  "country":"France"
		   },
		   "employee":[
			  {
				 "id":1,
				 "first_name":"Jeanette",
				 "last_name":"Penddreth"
			  },
			  {
				 "id":2,
				 "firstName":"Giavani",
				 "lastName":"Frediani"
			  }
		   ]
		}
	 }`)

	jsonParsed, err := gabs.ParseJSON(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Get value of Protected:\t", jsonParsed.Path("employees.protected").Data())
	fmt.Println("Get value of Country:\t", jsonParsed.Search("employees", "address", "country").Data())
	fmt.Println("ID of first employee:\t", jsonParsed.Path("employees.employee.0.id").String())
	fmt.Println("Check Country Exists:\t", jsonParsed.Exists("employees", "address", "countryCode"))

	addressJson, err := jsonParsed.Search("employees", "address").ChildrenMap()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for key, child := range addressJson {
		fmt.Printf("Key=>%v, Value=>%v\n", key, child.Data().(string))
	}

	employee, err := jsonParsed.Search("employees", "employee").Children()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, child := range employee {
		fmt.Println(child.Data())
	}

	firstEmployee, err := jsonParsed.Search("employees", "employee", "0").Children()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, child := range firstEmployee {
		fmt.Println(child.Data())
	}
}
