package entities

type Instrument struct {
	ID     int
	Symbol string
	Active bool
}

type Stock struct {
	Instrument
}

type MutualFund struct {
	Instrument
}
