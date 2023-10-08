package fecther

import (
	"strings"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	url := "https://www.baidu.com"
	bytes, err := Fetch(url)
	if err != nil {
		t.Error("fetch v% failed", url)
	}
	s := string(bytes)
	if !strings.Contains(s, "baidu") {
		t.Error("fetch v% failed", url)
	}
}
func TestFetchWithRateLimiter(t *testing.T) {
	for i := 0; i < 1000; i++ {
		url := "https://www.baidu.com"
		bytes, err := FetchWithRateLimiter(url, time.Tick(time.Nanosecond))
		if err != nil {
			t.Error("fetch v% failed", url)
		}
		s := string(bytes)
		if !strings.Contains(s, "baidu") {
			t.Error("fetch v% failed", url)
		}
	}
}
