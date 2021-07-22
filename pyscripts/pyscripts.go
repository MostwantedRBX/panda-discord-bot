package pyscripts

import (
	"log"
	"os"
	"os/exec"
)
import _ "embed"

//	I'll be using Python to convert images

//go:embed scripts/hello.py
var hellopy string

func RunScript() {
	cmd := exec.Command("python", "-c", hellopy)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}
