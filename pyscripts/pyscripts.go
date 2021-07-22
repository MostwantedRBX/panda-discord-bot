package pyscripts

import (
	"log"
	"os"
	"os/exec"
)
import _ "embed"

//	I'll be using Python to convert images

//go:embed scripts/ascii-converter.py
var ascii string

func RunScript(script string) {
	if script == "convert" {
		cmd := exec.Command("python", "-c", ascii, "./tacos.png")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Println(cmd.Run())
	}
}
