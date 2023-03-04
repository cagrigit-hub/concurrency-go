package main

import "testing"

func Test_updateMessage(t *testing.T){
	msg = "Hello, World!"

	wg.Add(2)
	go updateMessage("Goodbye good world!")
	go updateMessage("Goodbye cruel world!")
	wg.Wait()

	if msg != "Goodbye cruel world!" {
		t.Errorf("msg is not 'Goodbye cruel world!': %s", msg)
	}
}