package main

import "fmt"

type TemperatureConfig struct {
	value float64
}

func (tc *TemperatureConfig) Get() float64 {
	return tc.value
}

func (tc *TemperatureConfig) Set(value float64) {
	tc.value = value
}

type TemperautreOperations interface {
	Get() float64
}

type TemperatureMonitor struct {
	threshold TemperautreOperations
}

func NewMonitor(threshold TemperautreOperations) *TemperatureMonitor {
	return &TemperatureMonitor{threshold: threshold}
}

func (tm *TemperatureMonitor) checkTemperature(currTemp float64) {
	thresholdTemp := tm.threshold.Get()

	if currTemp > thresholdTemp {
		fmt.Printf("警告：當前溫度 %.2f 超過閾值 %.2f\n", currTemp, thresholdTemp)
	} else {
		fmt.Printf("正常：當前溫度 %.2f 在閾值 %.2f 以下\n", currTemp, thresholdTemp)
	}
}

func main() {
	monitor := NewMonitor(&TemperatureConfig{value: 30.0})

	monitor.checkTemperature(28.5)
	monitor.checkTemperature(33.5)
}
