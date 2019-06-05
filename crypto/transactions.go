package main

type OutPoint struct {
	hash  []byte // hash of the referenced Tx
	index int    // index of the specific output
}

// Transaction Input
type Tx_in struct {
	previous_output OutPoint
	// adresse - ajouter une adresse basique sans signature
}

// Transaction Output
type Tx_out struct {
	Value int
	// ajouter une adresse
}

// Transaction structure
type Tx struct {
	Version      int // Transaction data format version
	Tx_in_count  int // number of input tx
	Inputs       []Tx_in
	Tx_out_count int // number of output tx
	Outputs      []Tx_out
	Locked_tile  int
}
