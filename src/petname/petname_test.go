package petname

import (
	"testing"
)

// Make sure the generated names exist
func TestPetName(t *testing.T) {
	for i:=0; i<10; i++ {
		name := PetName(i, "-")
		if name == nil {
			t.Fatalf("Did not generate a %d-word name, '%s'", i, name)
		}
	}
}
