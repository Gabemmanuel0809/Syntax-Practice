package main 
import (
  "fmt"
  "unicode"
  "errors"
)

// Wel this is a bit overengineered for a practice

type NameError interface {
   InvName() error 
}

type AgeError interface {
   InvAge() error
}

type User struct {
    Name string 
    Age int 
}

func (u *User) InvName() error {
    if(u.Name == "") {
           return errors.New("Name cannot be empty")
    }

    for _, ch := range u.Name {
           if(!unicode.IsLetter(ch)) {
                  return fmt.Errorf("Invalid characters found")
           }
     }
    
     return nil
} 

func (u *User) InvAge() error {
        if (u.Age <= 0) {
             return errors.New("Invalid age")
        } 

        if(u.Age > 130) {
              return errors.New("Age too high")
        }

        if(u.Age > 10 && u.Age <= 100) {
             return nil 
         }

         return nil 
}

func checkName(n NameError) {
  fmt.Println(n.InvName())
}

func checkAge(a AgeError) {
  fmt.Println(a.InvAge())
}

func main() {
   var tempName string 
   fmt.Println("Input your name: ")
   fmt.Scan(&tempName)

   var tempAge int 
   fmt.Println("Input your age: ")
   fmt.Scan(&tempAge)

   fmt.Println(tempAge)

   var tempUser = User {
          Name: tempName,
          Age: tempAge,
    }

    checkName(&tempUser)
    checkAge(&tempUser)
}
