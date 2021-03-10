package main

import "fmt"

func main() {

	m := map[string][]string{
		"ball":   {"red", "grey"},
		"animal": {"black"},
	}

	res := append(m["animal"], "brown")
	fmt.Println(res)

	m["food"] = []string{"red", "green"}

	fmt.Println(m["food"])

	for i := range m["food"] {
		fmt.Println(i, m["food"][i])
	}
}
