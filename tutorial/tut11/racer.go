package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration > bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

const tenSecondTimeout = 10 * time.Second

func RacerRefactored(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// struct é usado em vez de bool pq tem o menor footprint
	// channels tem que ser sempre inicializado com make, se usar var ele é inicializado com nil
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

/*

select
-Helps you wait on multiple channels.
-Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.

httptest
-A convenient way of creating test servers so you can have reliable and controllable tests.
-Using the same interfaces as the "real" net/http servers which is consistent and less for you to learn.

*/
