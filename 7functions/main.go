package main
import ("fmt")
func saycountry(s string ) {
   fmt.Printf("%s, is a country\n",s)
}
func saycity(s string){
	fmt.Printf("%s, is a city\n",s)
}
func maxoftwo(a, b int )int {
	if a>b {
		return a
	}else{
		return b
	}
}

func main(){
	fmt.Println("funtions")
	saycountry("Ethiopia")
	saycity("Addis Ababa")
	var result = maxoftwo(3,5);
	fmt.Println("max is ",result)
}