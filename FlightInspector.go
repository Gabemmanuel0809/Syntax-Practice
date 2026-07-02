package main 

import(
  "fmt"
  "time"
)

type Flight struct {
   Aircraft string 
   FlightNo int 
   PilotsReady bool 
   EquipmentsReady bool 
   DamageChecked bool 
}

func InspectFlight(f Flight, ch chan string) {
   var InspectionResult string

   if f.PilotsReady == false {
	   InspectionResult = "Pilots not ready for Flight"
	   ch <- InspectionResult
   } else if f.EquipmentsReady == false {
	  InspectionResult = "Flight Equipments not ready"
	  ch <- InspectionResult
   } else if f.DamageChecked == false {
	  InspectionResult = "Damage not inspected yet"
	  ch <- InspectionResult
   } else {
      InspectionResult = "Flight is ready"
	  ch <- InspectionResult
   }

   time.Sleep(2 * time.Second)
}

func Takeoff(f Flight, ch chan string, ch1 chan string) {
    status := <- ch 
	var finalDecision string 

	if status != "Flight is ready" {
		finalDecision = "Flight is not permitted to takeoff yet"
		ch1 <- finalDecision
	} else {
		finalDecision = "Flight is now permitted to take off"
		ch1 <- finalDecision
	}

	time.Sleep(4 * time.Second)
}

func main() {
    ch := make(chan string)
	  ch1 := make(chan string)

    var f1 = Flight{
		   Aircraft:"Boeng 737",
		   FlightNo:5318,
		   PilotsReady: true,
		   EquipmentsReady: true,
		   DamageChecked: true,
	  }

	  go InspectFlight(f1, ch)
	  go Takeoff(f1, ch, ch1)

	  result := <- ch1 
	  fmt.Println(result)

}
