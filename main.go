package main

import (
	"fmt"
	"os"

	filecrypt "github.com/se7enroll/go-filecrypt/lib"
	"golang.org/x/term"
)

func main() {
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
	fmt.Println("\t encrypt \t Encrypts the file (in place) using a password.")
	fmt.Println("\t decrypt \t Decrypts the file (in place) using a password.")
	fmt.Println("\t help \t Displays the help text..")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("Missing file path argument.")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		println("File not found")
	}

	password := getPassword()
	fmt.Println("\nEncrypting fiel....")
	filecrypt.Encrypt(file, password)
	fmt.Println("Successfully encrypte the file.")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		println("Missing file path argument.")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		println("File not found")
	}

	password := getPassword()
	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\nFile successfully decrypted.")
}

func getPassword() []byte {
	fmt.Print("Enter password:")
	password, _ := term.ReadPassword(0)
	return password
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
