package main

func main() {
	//var card string = "Ace of spade"
	// splic1 := newDeck()
	// splic2 := newDeck()
	// hand1, remainhand1 := deal(splic1, 3)
	// splic1.print()
	// fmt.Println(hand1, remainhand1)
	// splic1.saveToFile("splice1 file")
	// fmt.Println("  ")
	// hand2, remainhand2 := deal(splic2, 7)
	// fmt.Println(hand2, remainhand2)
	// splic2.print()
	// splic2.saveToFile("splice2 file")

	// splic3 := newDeckFromFile("splice1 file")
	// splic3.print()
	// fmt.Println(" ")
	// splic3.shuffle()
	// splic3.print()

	//declareStruct().printStruct()
	// dom := person{name: "groot",
	// 	phoneNo: 789561230,
	// 	address: addressInfo{
	// 		houseNo: "C-77",
	// 		street:  "Mall",
	// 		city:    "Junaghar",
	// 		state:   "UP",
	// 		pinCode: 789456,
	// 	},
	// }

	//dom.colneStruct()
	//fmt.Printf("%+v", dom)
	//createMap()

	//var colorM map[string]string
	// colorM := map[string]string{
	// 	"white": "nh",
	// }
	// fmt.Println(colorM)

	// colorN := make(map[int]string)
	// colorN[2] = "jdfv"
	// //fmt.Println(colorN)
	// printMap(colorN)

	// en := english{}
	// sp := spanish{}

	// printGreeting(en)
	// printGreeting(sp)

	linksArr := []string{
		"http://google.com",
		"http://amzon.com",
		"http://facebook.com",
		"http://ttt.com",
		"http://golang.org",
	}

	for _, link := range linksArr {
		go checkLinks(link)
	}
}
