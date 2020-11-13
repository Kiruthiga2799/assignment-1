package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bolType, _ := json.Marshal(true)
	fmt.Println(string(bolType))
	intType, _ := json.Marshal(17)
	fmt.Println(string(intType))
	fltType, _ := json.Marshal(9.14)
	fmt.Println(string(fltType))
	strType, _ := json.Marshal("Json")
	fmt.Println(string(strType))
	slcA := []string{"air", "water", "land"}
	slcB, _ := json.Marshal(slcA)
	fmt.Println(string(slcB))
	mapA := map[string]int{"air": 1, "land": 2}
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))
}
