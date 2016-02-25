package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type class struct {
	class     string `json:"class"`
	agility   int    `json:"agility"`
	sheild    int    `json:"shield"`
	health    int    `json:"health"`
	name1     string `json:"name1"`
	damage1   int    `json:"damage1"`
	accuracy1 int    `json:"accuracy1"`
	speed1    int    `json:"speed1"`
	name2     string `json:"name2"`
	damage2   int    `json:"damage2"`
	accuracy2 int    `json:"accuracy2"`
	speed2    int    `json:speed2"`
}

func main() {
	// read the file, using the ioutil pacakge  https://golang.org/pkg/io/ioutil/#ReadFile
	bytes, err := ioutil.ReadFile("file.json")
	if err != nil {
		fmt.Println("There was an error reading the file, make sure it is in the same directory as main.go, or you have the correct path stated")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(json.class)
	fmt.Println(json.agility)
	fmt.Println(json.shield)
	fmt.Println(json.health)
	fmt.Println("Attack 1:")
	fmt.Println(json.name1)
	fmt.Println(json.damage1)
	fmt.Println(json.accuracy1)
	fmt.Println(json.speed1)
	fmt.Println("Attack 2:")
	fmt.Println(json.name2)
	fmt.Println(json.damage2)
	fmt.Println(json.accuracy2)
	fmt.Println(json.speed2)
}
