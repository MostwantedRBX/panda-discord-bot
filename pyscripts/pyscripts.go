package pyscripts

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "embed"
)

//	I'll be using Python to convert images

//go:embed scripts/ascii-converter.py
var ascii string

func RunScript(script string) {
	fmt.Println("Running a Python script")
	if script == "convert" {
		cmd := exec.Command("python", "-c", ascii, "./tacos.png") //	TODO: don't hardcode the filename.
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//	run and print the results and output of the python script
		log.Println(cmd.Run())
	}
}
