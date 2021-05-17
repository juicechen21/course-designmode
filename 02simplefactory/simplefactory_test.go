package simplefactory

import (
	"fmt"
	"testing"
)

func TestHelloAPI_Say(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("xiaoming")
	fmt.Println(s)
}