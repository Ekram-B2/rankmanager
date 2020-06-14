package rankclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	rankManager "github.com/Ekram-B2/rankmanager/rank"
)

func Test_createURL(t *testing.T) {
	type args struct {
		searchTerm    string
		realTerm      string
		searchTermLat string
		searchTermLng string
		realTermLat   string
		realTermLng   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// 1. Init parts of the application required to run the test (Act)
		{
			name: "NoSpaces",
			args: args{searchTerm: "tor",
				realTerm:      "toronto",
				searchTermLat: "123",
				searchTermLng: "234",
				realTermLat:   "123",
				realTermLng:   "234"},
			want: "http://127.0.0.1:8081/determineRank?searchTerm=tor&realTerm=toronto&searchTermLat=123&searchTermLng=234&realTermLat=123&realTermLng=234",
		},
		{
			name: "HasSpaces",
			args: args{searchTerm: "tor",
				realTerm:      "New York",
				searchTermLat: "123",
				searchTermLng: "234",
				realTermLat:   "123",
				realTermLng:   "234"},
			want: "http://127.0.0.1:8081/determineRank?searchTerm=tor&realTerm=New%20York&searchTermLat=123&searchTermLng=234&realTermLat=123&realTermLng=234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Compute output and check if result matches the expected (Act, Assert)
			if got := createURL(tt.args.searchTerm,
				tt.args.realTerm,
				tt.args.searchTermLat,
				tt.args.searchTermLng,
				tt.args.realTermLat,
				tt.args.realTermLng); got != tt.want {
				t.Errorf("createURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetRank(t *testing.T) {
	// 1. Set up temp to store actual value of the rmclient (Arrange)
	tempClient := rmClient
	// 2. Define the client function to use to use
	// to mock a call to the rank manager service (Arrange)
	inputRank := float32(0.5)
	inputName := "Toronto"
	rmClient = func(string) (*http.Response, error) {
		stream, _ := json.Marshal(rankManager.Rank{Name: inputName, Rank: inputRank})
		r := ioutil.NopCloser(bytes.NewReader([]byte(stream)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	// 2. Call rankClient to draw ouput (Act)
	rank, err := GetRank("", "", "", "", "", "")
	if err != nil {
		t.Fatalf("was not able to get rank")
	}
	// 3. Check to see if rank matches what was expected (Assert)
	if rank.Name != inputName {
		t.Fatalf("actual did not match expected; actual is %v and expected is %v", inputName, rank.Name)
	}

	if inputRank != rank.Rank {
		t.Fatalf("actual did not match expected; actual is %v and expected is %v", inputRank, rank.Rank)
	}
	// 4. Reset the client
	rmClient = tempClient
}
