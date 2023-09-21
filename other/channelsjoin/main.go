package channelsjoin

import "fmt"

func joinChannels(ch1 chan int, ch2 chan int) chan int {
	result := make(chan int)
	isFirstClosed := false
	isSecondClosed := false
	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					fmt.Println("ch1 closed")
					isFirstClosed = true
					ch1 = nil
					if isSecondClosed && isFirstClosed {
						close(result)
						return
					}
					break
				}
				result <- v
			case v, ok := <-ch2:
				if !ok {
					fmt.Println("ch2 closed")
					isSecondClosed = true
					ch2 = nil
					if isSecondClosed && isFirstClosed {
						close(result)
						return
					}
					break
				}
				result <- v
			}
		}
	}()
	return result
}
