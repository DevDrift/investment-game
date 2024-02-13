package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvent_Calculate(t *testing.T) {
	inputProfit := 100.0
	event := Event{
		Percent: 1.0, // 100 %
	}
	calculateProfit := event.Calculate(inputProfit)
	assert.Equal(t, calculateProfit, 200.0)
	event.Percent = -1.0 // - 100%
	calculateProfit = event.Calculate(inputProfit)
	assert.Equal(t, calculateProfit, 0.0)
	event.Percent = -1.1 // - 110%
	calculateProfit = event.Calculate(inputProfit)
	assert.Equal(t, calculateProfit, -10.000000000000014)
	event.Percent = -2.0 // - 200%
	calculateProfit = event.Calculate(inputProfit)
	assert.Equal(t, calculateProfit, -100.0)
}
