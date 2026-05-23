// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	age int
// }

// func (p *Person) Greet(){
// 	fmt.Printf("Hello, my name is %s and I am %d years old.\n",p.Name,p.age);
// }

// func (p *Person) setAge(age int){
// 	p.age = age;
// }

// func (p *Person) getAge() int {
// 	return p.age;
// }

// func main(){
// 	me := Person{Name:"Siddhant",age:12}
// 	me.setAge(13);
// 	fmt.Printf("Hello, my name is %s and I am %d years old.\n",me.Name,me.getAge());
// }