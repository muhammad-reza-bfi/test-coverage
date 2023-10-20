package generator

func GenerateOddOrEven(start int64) func() int64 {
	ch := make(chan int64)
	go func(c chan int64) {
		for {
			c <- start
			start += 2
		}
	}(ch)

	return func() int64 {
		return <-ch
	}
}
