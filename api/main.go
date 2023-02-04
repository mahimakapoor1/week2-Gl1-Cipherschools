package main

import (
	"log"
	"gin/basic/database"
	"gin/basic/routers"
)
func init(){
	database.Setup()
	database.Example()
}
func main(){
	database.Setup()
	engine:=routers.Router()
	err:=engine.Run("127.0.0.1:8080")
	if err!=nil{
		log.Fatal(err)
	}

}