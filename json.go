package main

import (
	// "encoding/json"
	"fmt"

	// "io/ioutil"
	// "log"
	// "os"
)

type Something struct{
	Name string `json:"name"`
	Married bool `json:"marriage"`
	Age float64 `json:"age"`
	Address []Address `json:"address"`
	Marks []int `json:"marks"`
}
type Address struct{
	Street int `json:"street"`
	City string `json:"city"`
}
// //
// func main() {
// 	jsonFile, err := os.Open("something.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer jsonFile.Close()
// 	jsonByteValues, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var something Something
// 	json.Unmarshal(jsonByteValues,&something)
// 	log.Println(something)
// 	fmt.Println(string(jsonByteValues))
// 	newJsonByteValues,err:=json.Marshal(something)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	os.WriteFile("some.json",newJsonByteValues)

	func main(){
		user:=map[int]interface{}{}
		user[1]="name"
		user[2]=true


		users:=map[string]bool{}
		users["golang"]=true
		fmt.Println(users["rahul"])
	}

