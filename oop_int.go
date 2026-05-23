// package main

// import "fmt"

// type NewPerson struct {
// 	Name string
// }

// type Robot struct {
// 	ID string
// }

// func (p * NewPerson) Greet(){
// 	fmt.Printf("Hi, I am %s \n",p.Name)
// }


// func (r * Robot) Greet(){
// 	fmt.Printf("Beep boop. ID: %s \n",r.ID)
// }

// type Greeter interface {
//  Greet()	
// }

// func SayHello(g Greeter){
// 	g.Greet()
// }

// func main(){
// 	me := NewPerson{Name:"Siddhant"}
// 	robot := Robot{ID: "simple id"}
// 	SayHello(&me)
// 	SayHello(&robot)
// }