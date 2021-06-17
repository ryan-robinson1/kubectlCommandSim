/*
	Ryan Robinson, 2021
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var connectorDataFilePath = "podData"

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func readStatus(connector string, data []string) string {
	for _, str := range data {
		if str[:len(connector)] == connector {
			return str[len(str)-1:]
		}
	}
	return ""
}
func readDeploymentNumber(connectorType string, data []string) int {
	i := 0
	for _, str := range data {
		if str[:len(connectorType)] == connectorType {
			i++
		}
	}
	return i
}
func readInData(filepath string) []string {
	data, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
func removeFile(filepath string) []string {
	data := readInData(filepath)
	err := os.Remove(filepath)
	if err != nil {
		panic(err)
	}
	return data
}
func createNewData(filepath string, connector string, newStatus string, data []string) []string {
	for i, con := range data {
		if con[:len(connector)] == connector {
			data[i] = con[:len(con)-1] + newStatus
			break
		}
	}
	return data
}
func createBlankData(filepath string, data []string) []string {
	for i, con := range data {
		data[i] = con[:len(con)-1] + "0"
	}
	return data
}
func replaceWithNewData(filepath string, data []string) {
	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, v := range data {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func printStatus(connector string, filepath string) {
	data := readInData(filepath)
	status := readStatus(connector, data)
	if status != "" {
		fmt.Println(status)
	} else {
		fmt.Println("Error: " + connector + " does not exist")
	}
}
func writeStatus(filepath string, connector string, newStatus string) {
	data := removeFile(filepath)
	newData := createNewData(filepath, connector, newStatus, data)
	replaceWithNewData(filepath, newData)
}
func resetStatus(filepath string) {
	data := removeFile(filepath)
	newData := createBlankData(filepath, data)
	replaceWithNewData(filepath, newData)
}

func getDeploymentNumber(filepath string, userClass string) int {
	data := readInData(filepath)
	count := readDeploymentNumber(userClass, data)
	return count
}
func handleArgs(args []string) {
	if args[0] == "status" && len(args) == 2 {
		printStatus(args[1], connectorDataFilePath)
	} else if args[0] == "scale" && len(args) == 3 && !isNumeric(args[2]) {
		fmt.Println("Error: " + args[2] + " is not a scaleable number")
	} else if args[0] == "scale" && len(args) == 3 && isNumeric(args[2]) {
		writeStatus(connectorDataFilePath, args[1], args[2])
		fmt.Println(args[2])
	} else if args[0] == "reset" && len(args) == 1 {
		resetStatus(connectorDataFilePath)
	} else if args[0] == "getDeploymentNumber" && len(args) == 2 {
		fmt.Println(getDeploymentNumber(connectorDataFilePath, args[1]))
	} else {
		fmt.Println("Error: Unrecognized arguments")
	}
}
func main() {
	args := os.Args[1:]
	handleArgs(args)
}

