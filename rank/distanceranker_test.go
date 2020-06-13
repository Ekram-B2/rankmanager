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

func testNormalizer(num, searchTermLen, realTermLen int) float32 {
	return float32(num)
}

func Test_rank_normalizer(t *testing.T) {

	tests := []struct {
		name             string
		searchTermLength int
		realTermLength   int
		num              int
		normalizer       normalizer
	}{
		// 1. Set up what is necessary to write test
		{
			name:             "normalizerDefault",
			searchTermLength: 10,
			realTermLength:   12,
			num:              5,
			normalizer:       defaultNormalizer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 2. Compute output and see if the output matches expectations
			actual := tt.normalizer(tt.num,
				tt.searchTermLength,
				tt.realTermLength)

			// 3. Check to see if the output fit within the expected bounds
			if actual < 0 && actual > 1 {
				t.Fatalf("the actual does not fit within the expected bounds")
			}
		})
	}
}

func Test_rank_distanceRanker(t *testing.T) {

	tests := []struct {
		name           string
		searchTerm     string
		realTerm       string
		normalizer     normalizer
		distanceRanker distanceRanker
		want           float32
	}{
		// 1. Set up what is necessary to write test
		{
			name:           "distanceRankerDefault",
			searchTerm:     "tor",
			realTerm:       "toronto",
			normalizer:     testNormalizer,
			distanceRanker: defaultDistanceRanker,
			want:           4.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 2. Compute output and see if the output matches expectations
			actual := tt.distanceRanker(tt.searchTerm,
				tt.realTerm,
				tt.normalizer)

			// 3. Check to see if the output fit within the expected bounds
			if actual != tt.want {
				t.Fatalf("the actual does not match what is expected; expected %v but got %v", tt.want, actual)
			}
		})
	}
}
