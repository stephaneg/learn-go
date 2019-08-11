package main

import "testing"

func TestGetSquareId(t *testing.T) {

	var id int8
	id = getSquareId(0, 0)
	if id != 0 {
		t.Errorf("square id for 0,0 is incorrect, got %d, want %d", id, 0)
	}

	id = getSquareId(0, 1)
	if id != 0 {
		t.Errorf("square id for 0,1 is incorrect, got %d, want %d", id, 0)
	}

	id = getSquareId(0, 3)
	if id != 1 {
		t.Errorf("square id for 0,3 is incorrect, got %d, want %d", id, 1)
	}

	id = getSquareId(7, 7)
	if id != 8 {
		t.Errorf("square id for 7,7 is incorrect, got %d, want %d", id, 8)
	}

}
