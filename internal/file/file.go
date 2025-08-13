package file

import (
	"fmt"
	"log"
	"os"
)

var baseRoute string = "assets/" 

func CreateFile(filename string){

	file, err := os.Create(baseRoute + filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	fmt.Println("File created at", baseRoute + filename)
}

func FileExists(filename string) bool{

	if file, err := os.Stat(baseRoute + filename); err == nil{
		fmt.Println(file.Name())
		return true
	}

	return false
}

func MakeDirectory(foldername string){
	err := os.Mkdir(foldername, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(filename string) ([]byte, error){
	file, err := os.ReadFile(baseRoute + filename)
	if err != nil {
		log.Fatal(err)
		return file, err
	}
	return file, err

}


func WriteFile(filename string, content []byte){
	err := os.WriteFile(baseRoute + filename, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}