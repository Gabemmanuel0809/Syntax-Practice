package main 
import (
	"fmt"
	"errors"
)

type DepositM interface {
    deposit() string  
}

type DepositErr interface {
	depositCheck() error
}

type User struct {
   Name string 
   Balance float64
}

type Bank struct {
    Name string 
	Users []User
}

func (u *User) deposit() string {
	var money float64 
	fmt.Println("Enter the money to deposit ")
	fmt.Scanf("%f", &money)

    err := u.depositCheck(money)
	if err != nil {
		return err.Error()
	}
	
	return "Deposit Success"
    return "Deposit Success"
}

func (u *User) depositCheck(money float64) error {
	if(money <= 0.0) {
	   return errors.New("Cannot deposit such value")
	}

	u.Balance += money 
	return nil 
}

func DepoMoney(d DepositM) {
	fmt.Println(d.deposit())
}

func checkDepo(m DepositErr) {
   fmt.Println(m.depositCheck())
}


func main() {
     var bank1 = Bank {
         Name: "bfc",
		 Users: []User {
            {Name:"Jane", Balance:0.00},
		 },
	 }

	 DepoMoney(&bank1.Users[0])
}
