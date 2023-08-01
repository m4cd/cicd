package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	want := "abcdefghijklmnop"
	w := make(http.Header)

	w.Set("Authorization", "ApiKey "+want)

	//log.Println(w)
	got, _ := GetAPIKey(w)
	//got := "abcdefghijklmnop"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
