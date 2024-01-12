package main

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "POST request Successful")
	name := r.FormValue("name")
	price := r.FormValue("price")
	fmt.Fprintf(w, "name = %s \n", name)
	fmt.Fprintf(w, "price = %s \n", price)
}
func filee(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}

}
func main() {
	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	// http.HandlerFunc("/form", filee)///
	// fmt.Printf("staring server at port 8080\n")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	// var gram1, rate1 float64
	const fix = 1000
	// fmt.Print("enter gram you want")
	// fmt.Scanln(&gram1)
	// fmt.Printf("your gram is %f \n", gram1)

	// fmt.Print("enter rate of product in kg")
	// fmt.Scanln(&rate1)
	// fmt.Printf("your rate is %f \n", rate1)

	// t := gram1 * rate1
	// total := t / fix

	// fmt.Printf("%f customer should give", total)

	var gram2, rate2 float64
	// const fix = 1000
	fmt.Print("enter amount you want")
	fmt.Scanln(&gram2)
	fmt.Printf("your gram is %f \n", gram2)

	fmt.Print("enter rate of product in kg")
	fmt.Scanln(&rate2)
	fmt.Printf("your rate is %f \n", rate2)

	t := gram2 * fix
	total := t / rate2

	fmt.Printf("%f this many gram we have to give customer", total)

}
func gram_u_want() float32 {
	var i float32
	fmt.Println("enter gram")
	fmt.Scanln(&i)
	return i
}
