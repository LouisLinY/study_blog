package main

import(
	"fmt"
)

type Person struct {
	name string
	age int
}

func (p Person)SetValue(name string, age int) {
	p.name = name
	p.age = age
	fmt.Printf("p addr : %p\n", &p)
}

func (p *Person)SetPoint(name string, age int){
	p.name = name
	p.age = age
	fmt.Printf("p addr : %p\n", p)
}

func main(){
	var p1 Person = Person{"yoyo", 5}
	fmt.Printf("p1 addr : %p\n", &p1)
	p1.SetValue("songjiang", 50)
	fmt.Printf("p1 %+v\n", p1)

	var p2 *Person= &Person{"louis", 20}
	fmt.Printf("p2 add : %p \n", p2)
	p2.SetPoint("MAHUATENG", 52)
	fmt.Printf("p2 %+v\n", p2)

    p3 := new(Person{"liudehua", 58})
    fmt.Printf("p3 addr : %p\n", p3)


}



