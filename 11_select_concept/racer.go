package selectconcept

import (
	"fmt"
	"net/http"
	"time"
)

const (
	tenSecondTime = time.Second * 10
)

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTime)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Function timed out")
	}
}
