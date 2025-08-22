//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b BandwidthUsage) AverageBandwidth() Bytes {
	var total Bytes

	for _, value := range b.amount {
		total += value
	}

	return total / Bytes(len(b.amount))
}

func (c CpuTemp) AverageTemp() Celcius {
	var total Celcius

	for _, value := range c.temp {
		total += value
	}

	return total / Celcius(len(c.temp))
}

func (m MemoryUsage) AverageMemory() Bytes {
	var total Bytes

	for _, value := range m.amount {
		total += value
	}

	return total / Bytes(len(m.amount))
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	dashboard := Dashboard{
		BandwidthUsage: BandwidthUsage{
			[]Bytes{50_000, 100_000, 130_000, 80_000, 90_000},
		},
		CpuTemp: CpuTemp{
			[]Celcius{50, 51, 53, 51, 52},
		},
		MemoryUsage: MemoryUsage{
			[]Bytes{800_000, 800_000, 810_000, 820_000, 800_000},
		},
	}

	fmt.Printf("Average Bandwidth: %d bytes/sec\n",
		dashboard.AverageBandwidth())

	fmt.Printf("Average CPU Temp: %.2f C\n",
		dashboard.AverageTemp())

	fmt.Printf("Average Memory Usage: %d bytes\n",
		dashboard.AverageMemory())
}
