package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurements.txt")
	if err != nil{
		panic(err)
	}
	defer measurements.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan(){
		rawData := scanner.Text() 
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := data[location]
		if !ok {
			measurement = Measurement{
				Min: temp,
				Max: temp,
				Sum: temp,
				Count: 1,
			} 
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		data[location] = measurement
	}

	locations := make([]string, 0, len(data))
	for name := range data{
		locations = append(locations, name) 
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for i, name := range locations {
		measurements := data[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f",
			name,
			measurements.Min,
			measurements.Sum/float64(measurements.Count),
			measurements.Max,
		)
		if i < len(locations)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}\n")
}
