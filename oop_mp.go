// package main

// import "fmt"


// type PetrolEngine struct {
// 	HorsePower int
// }

// type ElectricEngine struct {
// 	BatteryCapacity int
// }

// func (e * ElectricEngine) Start(){
// 	fmt.Printf("Electric engine started with %d KWH \n",e.BatteryCapacity)
// }


// func (p * PetrolEngine) Start(){
// 	fmt.Printf("Petrol engine started with %d HP\n",p.HorsePower)
// }

// type Engine interface {
// 	Start()
// }

// type Car struct {
// 	Brand string 
// 	Engine Engine
// }


// func  SayHello(e Engine) {
// 	e.Start()
// }
// func main(){
// 	petrolEngine := Car{Brand: "Petrol Car",Engine: &PetrolEngine{HorsePower: 40}}
// 	electricEngine := Car{Brand: "Petrol Car",Engine: &ElectricEngine{BatteryCapacity: 140}}
// 	petrolEngine.Engine.Start()
// 	electricEngine.Engine.Start()
// }

