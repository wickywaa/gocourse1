package main

import "testing"



gu

func TestDivide(t *testing.T) {

	_, err := divide(100.0, 10.0)
	if err != nil {
		t.Error("got an error weh we should not have")
	}
}

func TestBadDivide(t *testing.T) {

	_, err := divide(100.0, 0)
	if err = nil {
		t.Error(" did not got an error weh we should not have")
	}
}
