package main

import "fmt"

func main() {
	//simple switch
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")

	// case 2:
	// 	fmt.Println("two")

	// case 3:
	// 	fmt.Println("three")

	// }

	// MULTIPLE CONDITION  SWITCH

	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its a weekend")

	// default:
	// 	fmt.Println("its a workday kaam krooooo")
	// }

	//TYPE SWITCH VERY POWERFUL SWITCH

	whoAmI := func(i interface{}) {

		switch t := i.(type) {

		case int:
			fmt.Println("its an integer")

		case string:
			fmt.Println("its a string")

		case bool:
			fmt.Println("its boolean")

		case float32:
			fmt.Println("its a float")

		default:
			fmt.Println("others",t)
		}

	}
	whoAmI(12.5)
}


