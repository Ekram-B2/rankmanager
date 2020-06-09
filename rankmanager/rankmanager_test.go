package rankmanager

import "testing"

func Test_rankmanager_latlngDistCalculator(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Calculate output for operation (Act)
	actualOut := latlngDistCalculator(12.23, 45.45, 16.2, 50.2)

	// 3. Check to see if the output fit within the expected bounds
	if actualOut < 0 && actualOut > 1 {
		t.Fatalf("the actual does not fit within the expected bounds")
	}
}
