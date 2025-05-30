package modules

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// func Clear(){
// 	//fmt.Println("\033[H\033[2J") //VSCODE
// }

func Clear() {
	var cmd *exec.Cmd
	var err error

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println("[System] Gagal membersihkan layar:", err)
	}
}
