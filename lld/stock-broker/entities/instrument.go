package entities

type Instrument struct {
	Symbol string
	Active bool
}

type Stock struct {
	Instrument
}

type MutualFund struct {
	Instrument
}
