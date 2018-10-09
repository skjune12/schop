package schop

import (
	"reflect"
	"testing"
)

func TestGetFQDN(t *testing.T) {
	want := "cpu.sfc.wide.ad.jp."
	got, err := GetFQDN("203.178.142.142")
	if err != nil {
		t.Fatalf("Failed to test %#v", err)
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestGetAddrs(t *testing.T) {
	want := []string{"203.178.142.142", "2001:200:0:8803:203:178:142:142"}
	got, err := GetAddrs("cpu.sfc.wide.ad.jp.")
	if err != nil {
		t.Fatalf("Failed to test %#v", err)
	}

	if reflect.DeepEqual(want, got) != true {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestIsIPv6(t *testing.T) {
	want := true
	got := IsIPv6("2001:200:0:8803:203:178:142:142")
	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}

func TestIsIPv4(t *testing.T) {
	want := false
	got := IsIPv4("2001:200:0:8803:203:178:142:142")
	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}

func TestSearch(t *testing.T) {
	want := Result{
		Fqdn: "cpu.sfc.wide.ad.jp.",
		IPv4: "203.178.142.142",
		IPv6: "2001:200:0:8803:203:178:142:142",
	}
	got := Search("2001:200:0:8803:203:178:142:142")
	if reflect.DeepEqual(want, got) != true {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}
