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

func GenerateOddOrEvenRaw(start int64) <-chan int64 {
	ch := make(chan int64)
	go func(c chan int64) {
		for {
			c <- start
			start += 2
		}
	}(ch)

	return ch
}

func FanIn(channels ...<-chan int64) <-chan int64 {
	out := make(chan int64)
	for _, channel := range channels {
		go func(out chan int64, channel <-chan int64) {
			for {
				out <- <-channel
			}
		}(out, channel)
	}

	return out
}
