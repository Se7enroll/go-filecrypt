package main

import (
	"fmt"
	"os"
)

func main() {
	// Ensure all
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file and decrypt to decrypt a file.")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Simple file encryption.")
	fmt.Println("Usage:")
	fmt.Println("\tGo run . encrypt /path/to/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("\t encrypt \t Encrypts the file using a password.")
	fmt.Println("\t decrypt \t Decrypts the file using a password.")
	fmt.Println("\t help \t Displays the help text..")
}

func encryptHandle() {

}

func decryptHandle() {

}

func getPassword() {

}

func validatePassword() {

}

func validateFile() {

}
