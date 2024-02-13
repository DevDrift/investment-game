package core

type Event struct {
	Type    string  `json:"type"`
	Percent float64 `json:"percent"`
}

func (e *Event) Calculate(inputProfit float64) (result float64) {
	prc := e.Percent
	result = inputProfit + (inputProfit * prc)
	return
}

/*
e: 50%
1a: p = 1.0 ((p * 0.5) + p)
2a: p = 1.5 ((p * 0.5) + p)
*/
