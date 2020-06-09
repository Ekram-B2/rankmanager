package rankmanager

import "testing"

func Test_rankmanager_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "AGreaterThanB",
			args: args{a: 6, b: -1},
			want: -1,
		},
		{
			name: "BGreaterThanA",
			args: args{a: -1, b: 6},
			want: -1,
		},
		{
			name: "AEqualToB",
			args: args{a: 6, b: 6},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Nothing to init to set up the test (Arrange)

			// Peform operation and check to see if output matches the expected (Act, Assert)
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rankmanager_normalizer(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Apply operation to get output (Act)
	actualOut := defaultNormalizer(10, 11, 12)

	// 3. Determine if the output fits the required bounds (Assert)
	if actualOut < 0.0 && actualOut > 1.0 {
		t.Fatalf("the output is outside of the bounds of the what is required to be a normalizer")
	}
}

func testNormalizer(num, searchTermLen, realTermLen int) float32 {
	return float32(num)
}
func Test_rankmanager_defaultDistanceCalculator(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Apply operation to get output (Act)
	realTerm := "tor"
	actualOut := defaultDistanceCalculator("", realTerm, testNormalizer)

	// 3. Define expected out
	expectedOut := len(realTerm)

	// 4. See if the actual matches with what was expected (Assert)
	if actualOut != float32(len(realTerm)) {
		t.Fatalf("actual did not match expected; actual was %v and expected was %v", actualOut, expectedOut)
	}
}
