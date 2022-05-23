
package test

import (
	"log"
	"fmt"
	"os"
)

func hello() {
	/*
	file, err := os.Open("/home/martin2/")

	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}

	defer file.Close()

	fileList,_ := file.Readdir(0)
	
	fmt.Printf("\nName\t\tSize\tIsDirectory  Last Modification\n")
	
	for _, files := range fileList {		
		fmt.Printf("\n%-15s %-7v %-12v %v", files.Name(), files.Size(), files.IsDir(), files.ModTime())
	}
	*/

	fmt.Println("!... Hello World ...!")
}
