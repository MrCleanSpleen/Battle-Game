package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Race struct {
	class   string `json:"class"`
	health  int    `json:"health"`
	agility int    `json:"agility"`
	shield  int    `json:"shield"`
	attacks []Move `json:"attacks"`
}

type Move struct {
	name     string `json:"name"`
	damage   int    `json:"damage"`
	accuracy int    `json:"accuracy"`
	speed    int    `json:"speed"`
}

func main() {
	bytes, err := ioutil.ReadFile("class.json")
	if err != nil {
		fmt.Println("There was an error reading the file")
		fmt.Println("make sure it is in the same directory as main.go, or you have the correct path stated")
		panic(err)
	}
	jsonClass := &Race{}
	err = json.Unmarshal(bytes, jsonClass)
	//decodes it
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("%v: %s <<--ERROR %s\n", e, bytes[:e.Offset], bytes[e.Offset:])
		} else {
			fmt.Println("There was an error decoding the file")
		}
		panic(err)
	}

	fmt.Println(jsonClass.class)
	fmt.Println(jsonClass.attacks[0].name)

}
