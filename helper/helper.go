package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	ps := exec.Command("powershell", "cls", "clear", "/c")
	ps.Stdout = os.Stdout
	ps.Run()
}

func Backhandling() {
	fmt.Println("Tekan 'Enter' untuk kembali ke menu utama")
	fmt.Scanln()
	ClearScreen()
}
