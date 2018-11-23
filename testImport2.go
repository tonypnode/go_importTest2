package testImport2

import "fmt"

func main(){
	fmt.Println("i am in the main of the imported package")
}

func ImportMe(){
	fmt.Println("I was imported and called yaay for me")
}

func cantImportMe(){
	fmt.Println("na na na na boo boo, can't run me from import")
}

