package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

var ShadowFolder string = fmt.Sprintf("%v/.stand", os.Getenv("HOME"))

const ShadowAuthFile string = "/auth"
const ShadowGUIFile string = "/gui"
const GUIFetchLink string = "https://github.com/bitspaceorg/STAND-FRONTEND-FOSSHACK"

func GetShadowAuthFilePath() string {
	return fmt.Sprintf("%v%v", ShadowFolder, ShadowAuthFile)
}

func GetShadowGUIFilePath() string {
	return fmt.Sprintf("%v%v", ShadowFolder, ShadowGUIFile)
}

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func AppendToFile(filename, data string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	return err
}
