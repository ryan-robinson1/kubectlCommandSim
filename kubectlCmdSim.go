/*
	Ryan Robinson, 2021

	kubectlCmdSim simulates the output of kubectl requests locally and can be used for simple testing. The  simulated connector status data
	is stored in a local text file called "connectorData" in this format, where each connector has a corresponding scale:

	connector1:0
	connector2:1
	connector3:0


	To read the status of the connector, call:

	./kubectlCmdSim status connectorName                   ex: ./kubectlCmdSim status connector1

	To change the scale of the connector, call:

	./kubectlCmdSim scale connectorName connectorScale     ex: ./kubectlCmdSim scale connector1 1

	To reset every connector status to zero, call:

	./kubectlCmdSim reset
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var connectorDataFilePath = "connectorData"

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
func printStatus(connector string, filepath string) {
	data := readInData(filepath)
	status := readStatus(connector, data)
	if status != "" {
		fmt.Println(status)
	} else {
		fmt.Println("Error: " + connector + " does not exist")
	}
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
	} else {
		fmt.Println("Error: Unrecognized arguments")
	}
}
func main() {
	args := os.Args[1:]
	handleArgs(args)
}