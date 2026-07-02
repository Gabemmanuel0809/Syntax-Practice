package main 

import(
  "fmt"
  "reflect"
)

func DataTypeInspector(x any, ch chan string) {
	 // wow theres no other better variable name eh
	 var judgement string 

   // yeah Switch case is much better here
   // and seriously channels are not so very needed here
   if(reflect.TypeOf(x).String() == "int") {
		  judgement = "int"
		  ch <- judgement
	 } else if(reflect.TypeOf(x).String() == "string") {
	 	 judgement = "string"
		 ch <- judgement
	 } else if(reflect.TypeOf(x).String() == "float64") {
		 judgement = "float64"
		 ch <- judgement
	 } else if(reflect.TypeOf(x).String() == "bool") {
		 judgement = "boolean"
		 ch <- judgement
	 } else if(reflect.TypeOf(x).String() == "rune") {
		 judgement = "rune"
		 ch <- judgement
	 } else {
		 judgement = "unknown"
		 ch <- judgement
	 }
}

func main() {
	ch := make(chan string)
  // yeah, even this goorutine is not so needed
  go DataTypeInspector(rune(64), ch)
  result := <- ch 

	fmt.Println("Inspection Result: ", result)
}
