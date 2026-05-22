package main 
import "fmt"

type Product struct {
   Name string 
   Quantity int 
   Price float64
}

func addProd(store []Product) []Product {
  var name string 
	fmt.Println("Enter the new products name: ")
	fmt.Scanf("%s", &name)

	var qty int 
	fmt.Println("Enter the number of quantity: ")
	fmt.Scanf("%d", &qty)

	var price float64 
	fmt.Println("Enter the price of the product: ")
	fmt.Scanf("%f", &price)

	var newProd = Product {
		 Name: name,
		 Quantity: qty,
		 Price: price,
	 }

	store = append(store, newProd)
	fmt.Println("Product succesfully added")
	return store
}

func updateProd(store []Product) {
	var prodName string 
	fmt.Println("Enter the product name you wanted to update: ")
	fmt.Scanf("%s", &prodName)

	for i := range store {
		if(store[i].Name == prodName) {
			var newName string 
			fmt.Println("Enter the new name: ")
            fmt.Scanf("%s", &newName)

			store[i].Name = newName
			return
		 }
	 }
 }

func main() {
   var store = []Product {
      {Name: "Pencil", Quantity:30, Price: 1.50},
	 }

	 addProd(store)  
	 updateProd(store)

	 fmt.Println(store)
}
