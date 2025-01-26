package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://search.brave.com",
		"http://proton.me",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://search.brave.com": true,
		"http://proton.me":        true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:\t%v\ngot: \t%v", want, got)
	}
}
