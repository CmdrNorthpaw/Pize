package main

type DummyDrone struct {
	SDKMode bool
	Airborne bool
	VideoOn bool

	Position Position
	Spped int16
}

type Position struct {
	X float64
	Y float64
	Z float64
	Rotation float64
}