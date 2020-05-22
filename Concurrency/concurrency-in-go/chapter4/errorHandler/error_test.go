package errorHandler

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandleErr(t *testing.T) {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		response := make(chan *http.Response)
		go func() {
			defer close(response)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case response <- resp:
				}
			}
		}()
		return response
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"www.google.com", "https://baidu.cocm"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("is is ok ? %v", response)
	}
}

// The key thing to note here is
// how we've coupled the potential result with
// the potential error
func TestHandleErr2(t *testing.T) {
	type Result struct {
		Error    error
		Response *http.Response
	}
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})

	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}







