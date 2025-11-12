package pokeapi

import (
	"fmt"
	"testing"
)

func TestRequests(t *testing.T) {
	err := CommandExplore(nil, nil, nil, "canalave-city-area")
	if err != nil {
		t.Errorf("expected no errors: %v", err)
	}

	err = CommandExplore(nil, nil, nil, "dingdong")
	if err == nil {
		t.Errorf("expected that to not be found bro")
	}
	fmt.Println(err)
}
