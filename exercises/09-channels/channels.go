package main

// TODO:
// Send numbers 1–10 into a channel
// Square them in a worker goroutine
// Collect results

func square(input, output chan int) {
	for num := range input {
		output <- num * num
	}
	close(output)
}

func main() {
	input := make(chan int)
	output := make(chan int)

	go square(input, output)

	go func() {
		for result := range output {
			println(result)
		}
	}()

	for i := 1; i < 10; i++ {
		input <- i
	}

	close(input)
}
