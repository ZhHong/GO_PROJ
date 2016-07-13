package main

import "utils"
import "fmt"

func main(){
	utils.LogInfo(fmt.Sprintf("-----------------------begin %d",1))
	utils.LogError("----------------------begin")
	utils.LogWarning("=====================end")
	utils.LogInfo("=======================end ")
}