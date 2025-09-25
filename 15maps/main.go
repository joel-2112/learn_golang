package main
import ("fmt")

func main()  {
	fmt.Println("maps in go")
	//initialize the map
	countryCapital := make(map[string]string)
	countryCapital["Amhara"] = "Bahir Dar"
	countryCapital["Tigray"] = "Mekele"
	countryCapital["Afar"] = "Semera"
	for country := range countryCapital{
		fmt.Printf("The capital city of %s is %s\n", country, countryCapital[country])
	}
	delete(countryCapital, "Tigray")
	fmt.Println("the country Tigray is deleted")
	fmt.Println("the updated map is")
	for country := range countryCapital{
		fmt.Printf("the capital city of %s is %s\n", country, countryCapital[country])
	}
}