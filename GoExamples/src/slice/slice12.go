package main

import "fmt"

func getCountries() []string {
	countries := []string{"USA", "FRA", "CN", "JP", "KOR"}
	neededCountries := countries[:len(countries)-2] // [USA FRA CN]
	//fmt.Println("neededCountries is", neededCountries)
	countriesCopy := make([]string, len(neededCountries)) // create an empty copy
	// fmt.Println("countriesCopy is", countriesCopy)
	copy(countriesCopy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCopy
}

func main() {
	countriesNeeded := getCountries()
	fmt.Println("countriesNeeded is", countriesNeeded)
}
