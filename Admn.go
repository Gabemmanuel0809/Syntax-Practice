package main 
import "fmt"

type User struct {
   Name string 
   Role string 
}

func (u User) RunAdminCapabilities() {
	if(u.Role == "Administrator") {
		fmt.Println("Running Administrator's System")
	} else {
		fmt.Println("Action Denied, Current Role: ", u.Role)
	}
}

func (u User) FireThem(usr *User) {
	if(u.Role == "Administrator") {
    if(usr != nil) {
			usr.Role = "Visitor"
		} else {
			 fmt.Println("User not found")
		}
	} else {
		fmt.Println("Action Denied, Current Role: ", u.Role)
	}
}

func main() {
  var admin = User {Name:"Bernie", Role:"Administrator"}
	var regUsr = User {Name:"Bob", Role:"Employee"}
    
	admin.RunAdminCapabilities()
	regUsr.RunAdminCapabilities()

	// Sorry Bob, You tried tampering with the Administrator tools
	// Admin: Congratulations bob, you where now promoted to Visitor
	admin.FireThem(&regUsr)

	fmt.Println(regUsr)
}
