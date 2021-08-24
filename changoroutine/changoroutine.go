package changoroutine

type adder interface {
	AddCounter(int64)
}

func TryChanGoroutine(intChan chan<- int, add adder) {
	// Add the logic to complete the tests in here

	add.AddCounter(64)

	// go add.AddCounter(64)
	// intChan <- 10
	// allSum := 0
	// for i := 0; i <= 10; i++ {
	// 	allSum += i
	// 	add.AddCounter(64)
	// }
	// intChan <- allSum
}
