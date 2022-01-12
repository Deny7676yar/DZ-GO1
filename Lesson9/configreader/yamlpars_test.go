package configreader

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParsConfigYaml(t *testing.T) {
	filename := "../config.yaml"
	got := ParsConfigYaml(filename)
	want := "John"

	if got != want{
		t.Fatalf("exep: %v, got: %v", want, got)
	}
}

func TestWrongFile(t *testing.T) {
	filename := "..config.yaml"
	assert.Panics(t, func(){
		ParsConfigYaml(filename)
	})
}
