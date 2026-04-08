package main

import(
	"os"
	"fmt"
	"reloaded/GO-RELOADED"

)
func main(){
	if len(os.Args) !=3 {
		return
	}
	//take filename as input
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	
	//access file content
	data, err :=os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}
	result:= reloaded.Modify(string(data))
	
	err =os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil{
		fmt.Println("Error writting file", err)
	}

}