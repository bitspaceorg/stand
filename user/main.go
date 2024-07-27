package User

import (
	"bufio"
	"fmt"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
	"os"
	"strings"
)

func CreateUser() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	hashedPassword := utils.HashPassword(password)

	userData := fmt.Sprintf("%s:%s\n", username, hashedPassword)

	if _, err := os.Stat(utils.ShadowFolder); os.IsNotExist(err) {
		os.MkdirAll(utils.ShadowFolder, 0700)
	}

	err := utils.AppendToFile(utils.GetShadowAuthFilePath(), userData)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	fmt.Println("User created successfully.")
}

func ValidateUser() bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	hashedPassword := utils.HashPassword(password)

	file, err := os.Open(utils.GetShadowAuthFilePath())
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid shadow file format.")
			return false
		}

		if parts[0] == username && parts[1] == hashedPassword {
			fmt.Println("User authenticated.")
			return true
		}
	}

	fmt.Println("Invalid username or password.")
	return false
}
