package main

import "fmt"

func fhallo1(name string) string {
	if name == "" {
		return "hallo"
	} else {
		return "haro " + name
	}

}
func main() {
	h := fhallo1("jon")
	fmt.Println(h)

	fmt.Println(fhallo1(""))
}
