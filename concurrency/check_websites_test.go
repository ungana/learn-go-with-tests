package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "fake://google.com"
}

func slowStubCheckWebsites(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"fake://google.com",
		"https://bing.com",
	}

	want := map[string]bool{
		"https://google.com": true,
		"fake://google.com":  false,
		"https://bing.com":   true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func BenchmarkCheckWebites(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubCheckWebsites, urls)
	}
}
