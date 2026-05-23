package main 
import "fmt"

type MainGun interface {
    fire() string 
}

type Ship struct {
	Name string 
	Class string
	Loaded bool 
}

func (s *Ship) fire() string {
  if(s.Loaded == true) {
	   s.Loaded = false 
	   return "Main Battery fired a shot"
   } else {
	   return "Main Battery is not loaded yet"
   }
}

func feuer(m MainGun) {
	fmt.Println(m.fire())
}

func main() {
  var s = Ship {Name:"HMS Valiant", Class:"Battleship", Loaded:true}

  feuer(s)
	feuer(s)
}
