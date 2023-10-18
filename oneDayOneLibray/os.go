package oneDayOneLibray

import (
	"fmt"
	"os"
)

func TestOs() {
	homeDir := os.Getenv("HOME")
	fmt.Println(homeDir)
}
