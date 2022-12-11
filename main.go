package main

import (
	"tokobelanja/database"
	_ "tokobelanja/initializer"
)

func init() {
	database.Connect()
}

func main() {

}
