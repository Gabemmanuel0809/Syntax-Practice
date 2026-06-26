package main 

import(
	"fmt"
	"reflect"
	"errors"
)

func ageChecker(x int) (error) {

	if(x <= 0) {
		return errors.New("Invalid Age")
	}

	return nil 
}

func main() {
	var age float64 = 7.55

	if(reflect.TypeOf(age).String() == "float64") {
		cnva := int(age)
		fmt.Println(ageChecker(cnva))
	} else {
		fmt.Println(age)
	}
}
