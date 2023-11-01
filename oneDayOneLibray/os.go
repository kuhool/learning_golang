package oneDayOneLibray

import (
	"fmt"
	"os"
	"path/filepath"
)

func TestOs() {
	workPath, err := os.Getwd()
	fmt.Println(workPath, err)
}

func TestOs2() {

	AppPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(os.Args[0], filepath.Dir(os.Args[0]), AppPath, err)
}

func TestOs1() {
	homeDir := os.Getenv("HOME")
	fmt.Println(homeDir)
}
