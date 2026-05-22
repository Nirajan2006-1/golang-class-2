package main
import "fmt"
func c2f(degree float64) float64 {
return (degree*9/5)+32
}

func suhag(){
fmt.Println("Hello class!\n")
}

func nirajan(name string){
fmt.Println("Hello" + name)
}

func temp(degree float64){
fmt.Println("The temp is" , degree)
}

func age (year int){
fmt.Println("The age is" , year)
}

func option(choice bool){
fmt.Println("This is" , choice)
}

func main(){
fmt.Println (c2f(30))
nirajan(" Ayush")
temp(26.16)
age(20)
}
