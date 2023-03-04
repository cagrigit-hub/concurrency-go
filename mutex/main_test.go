package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T){
	stdOut := os.Stdout
	r,w,_ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	out,_ := io.ReadAll(r)
	os.Stdout = stdOut
	output := string(out)

	if (!strings.Contains(output, "85800.00")){
		t.Errorf("Expected 85800.00 but got %s", output)
	}
	
}