package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	/*
		if url == "waat://furhurterwe.geds" {
			return false
		}
		return true
	*/
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}

}

/*
-goroutines, the basic unit of concurrency in Go, which let us check more than one website at the same time.
-anonymous functions, which we used to start each of the concurrent processes that check websites.
-channels, to help organize and control the communication between the different processes, allowing us to avoid a race condition bug.
-the race detector which helped us debug problems with concurrent code
*/
