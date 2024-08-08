package cars

// CalculateWorkingCarsPerHour calculates how many working cars are produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const (
		groupCost      = 95000
		individualCost = 10000
	)

	groups := carsCount / 10
	individuals := carsCount % 10
	totalCost := (groups * groupCost) + (individuals * individualCost)

	return uint(totalCost)
}
