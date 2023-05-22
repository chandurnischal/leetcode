package leetcode

// question 2469
func ConvertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, (celsius * 1.8) + 32.0}
}
