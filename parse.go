package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyKey struct {
	streetAddress string
	town          string
	saleDate      string
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Usage: ./Parse <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	rawFile, e := readFileIntoArray(filename)
	if e != nil {
		fmt.Println(e)
		return
	}
	println(len(rawFile))
	//testOne(rawFile)
	//testTwo(rawFile)
	//testThree(rawFile)
	//prettyPrintList(filterByCheapProperties(rawFile))
	//prettyPrintList(filterOutCertainStreets(rawFile))
	prettyPrintList(filterOneInTen(rawFile))
}

/*In the case of duplicates, use the last encountered record.*/
func testOne(rawFile [][]string) {

	m := make(map[MyKey][]string)

	for i := range rawFile {
		key := extractAddressAndDate(rawFile[i])
		m[key] = rawFile[i]
	}

	prettyPrintMap(m)
}

/*Modify the code in case of duplicates to use the first encountered record.*/
func testTwo(rawFile [][]string) {

	m := make(map[MyKey][]string)
	//Test for the existence of a key in the map, and if it doesn't exist, add
	for i := range rawFile {
		key := extractAddressAndDate(rawFile[i])

		_, ok := m[key]
		if !ok {
			m[key] = rawFile[i]
		}
	}

	fmt.Println("Test two")

}

/*Instead of inserting the last record, make sure that no duplicates are entered at all. */
func testThree(rawFile [][]string) {

	toRemove := make(map[MyKey][]string)   //the collection of duplicates to remove
	sliceAsMap := make(map[MyKey][]string) //the collection to return

	for i := range rawFile {
		key := extractAddressAndDate((rawFile[i]))
		//if it already exists, add ot to the toRemove map
		_, ok := sliceAsMap[key]
		if !ok { //if it's not there, add it
			sliceAsMap[key] = rawFile[i]
		} else {
			toRemove[key] = rawFile[i]
			sliceAsMap[key] = rawFile[i]
		}
	}

	for key := range sliceAsMap {
		_, ok := toRemove[key]
		if ok {
			delete(sliceAsMap, key)
		}
	}
	prettyPrintMap(sliceAsMap)
}

func filterByCheapProperties(rawFile [][]string) [][]string {
	threshold := 400000
	valueIndex := 4

	//var lines [][]string
	var propertyList [][]string
	for _, property := range rawFile {
		propertyValue, _ := strconv.Atoi(property[valueIndex])

		if propertyValue > threshold {
			propertyList = append(propertyList, property)
		}
	}

	return propertyList
}
func filterOutCertainStreets(rawFile [][]string) [][]string {
	streetIndex := 1
	var propertyList [][]string
	for _, property := range rawFile {
		if !strings.Contains(property[streetIndex], "AVE") && !strings.Contains(property[streetIndex], "CRES") && !strings.Contains(property[streetIndex], "PL") {
			propertyList = append(propertyList, property)
		}
	}
	return propertyList
}

func filterOneInTen(rawFile [][]string) [][]string {
	var propertyList [][]string
	for index, property := range rawFile {
		if !(index%10 == 0) {
			propertyList = append(propertyList, property)
		} else {
			fmt.Println(index)
		}
	}
	return propertyList
}

/*Helper function to make a map print out nicely*/
func prettyPrintMap(m map[MyKey][]string) {
	for _, value := range m {
		fmt.Println(value)
	}
}

func prettyPrintList(rawFile [][]string) {
	for _, element := range rawFile {
		fmt.Println(element)
	}

}

func readFileIntoArray(filename string) ([][]string, error) {
	file, e := os.Open(filename)

	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	defer file.Close()
	var lines [][]string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if scanner.Text() != "" { //Gotta check for empty lines.
			line := strings.Split(scanner.Text(), "\t")
			lines = append(lines, line)
		}
	}
	return lines, e
}

func extractAddressAndDate(propertySale []string) MyKey {
	return MyKey{streetAddress: propertySale[1], town: propertySale[2], saleDate: propertySale[3]}
}
