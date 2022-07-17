package wtf

import "testing"

func RequireNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("error was not nil as required : %v", err)
	}
}
