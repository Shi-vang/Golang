package main

import "fmt"

type addressInfo struct {
	houseNo string
	street  string
	city    string
	state   string
	pinCode int
}

type person struct {
	name    string
	phoneNo int
	address addressInfo
}

func declareStruct() person {
	dom := person{name: "groot",
		phoneNo: 789561230,
		address: addressInfo{
			houseNo: "C-77",
			street:  "Mall",
			city:    "Junaghar",
			state:   "UP",
			pinCode: 789456,
		},
	}
	fmt.Printf("%+v", dom)
	return dom
}

func (d person) printStruct() {
	d.name = "boot"
	fmt.Printf("%+v", d)
}

func (d *person) colneStruct() {
	d.name = "shoot"
	fmt.Printf("%+v", d)
}
