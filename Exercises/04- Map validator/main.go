package main

import "fmt"

func main() {

	dictionary := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}
	var promptValue string
	var output []string
	var presence bool

	for {
		fmt.Printf("Ingresa un valor: ")
		fmt.Scanf("%s\n", &promptValue)
		_, presence = dictionary[promptValue]
		if promptValue == "0" {
			break
		} else if presence == true {
			output = append(output, dictionary[promptValue])
		} else {
			output = append(output, "No encontrado")
		}
	}
	fmt.Println(output)

}
