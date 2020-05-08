//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

//		"James", "Bond", "Shaken, not stirred"

//		"Miss", "Moneypenny", "Helloooooo, James."

//Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	jm := [][]string{jb, mp}
	fmt.Println(jm)

	for i, v := range jm {
		fmt.Println("record", i)
		for j, val := range v {
			fmt.Println("\tindex value", j, "\tRecord:", val)
		}
	}
}
