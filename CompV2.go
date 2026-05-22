package main 
import "fmt"
 
type Employee struct {
	 Name string
	 Role string
}

type Company struct {
	 Name string 
	 Employees []Employee
}

func (c *Company) Fire() { 
	var inpName string
	fmt.Println("Input the employee name: ") 
	fmt.Scanf("%s", &inpName) 
	for i, value := range c.Employees {
		 if(inpName == value.Name) {
			 c.Employees = append(c.Employees[:i], c.Employees[i+1:]...)
			 fmt.Println("Emloyee Fired")
		 }
	} 
} 

func main() {
	 var cmp = Company {
		  Name: "XnL Corporation",
		  Employees: []Employee {
			  {Name:"bob", Role:"Security Personnel"},
			  {Name:"Linda", Role:"Financial officer"},
		   },
	 }
	  
	cmp.Fire("bob") // Sorry bob, we have to do it for the greater good

	fmt.Println(cmp.Employees)
}
