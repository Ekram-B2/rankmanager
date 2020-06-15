package rankclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	l4g "github.com/alecthomas/log4go"

	rankManager "github.com/Ekram-B2/rankmanager/rank"
	"github.com/Ekram-B2/suggestionsmanager/config"
)

type retreiveRank func(string) (*http.Response, error)

var (
	rmClient retreiveRank
)

func init() {
	switch os.Getenv("RANK_MANAGER_CLIENT") {
	case "default":
		rmClient = http.Get
	default:
		rmClient = http.Get
	}
}

// GetRank is the algorithm used to retreive rank for the real term retreived from persistance against the research term
func GetRank(searchTerm, realTerm, realTermLat, searchTermLat, realTermLng, searchTermLng string) (rankManager.Rank, error) {
	// 1. Commit GET request to retreive the rank for this datapoint from the remote server
	resp, err := rmClient(createURL(searchTerm, realTerm, realTermLat, searchTermLat, realTermLng, searchTermLng))
	if err != nil {
		l4g.Error("SYSTEM-ERROR: unable to process request to retreive the rank: %s", err.Error())
		return rankManager.Rank{}, err
	}
	defer resp.Body.Close()
	// 2. Read the contents of the return out to a stram
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to read byte stream from response body: %s", err.Error())
		return rankManager.Rank{}, err
	}
	// 3. Unmarshall the contents out to a rank object
	var returnedRank rankManager.Rank

	err = json.Unmarshal(contents, &returnedRank)
	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to unmarshall byte stream into rank structure: %s", err.Error())
		return rankManager.Rank{}, err
	}

	// 4. Return the rank
	return returnedRank, nil
}

// createURL is what's used to generate the URL to commit a request from the remote server
func createURL(searchTerm, realTerm, realTermLat, searchTermLat, realTermLng, searchTermLng string) string {

	// 1. Return the appropriate URL based on whether the environment is development or production
	if config.IsDevelopmentEnvironment(os.Getenv("DEPLOYMENT_TYPE")) {
		return os.Getenv("DEVELOPMENT_SERVICE_PATH") + createPath(searchTerm, realTerm, realTermLat, searchTermLat, realTermLng, searchTermLng)
	}
	return os.Getenv("PRODUCTION_SERVICE_PATH") + createPath(searchTerm, realTerm, realTermLat, searchTermLat, realTermLng, searchTermLng)
}

func createPath(searchTerm, realTerm, realTermLat, searchTermLat, realTermLng, searchTermLng string) string {
	// 1. Modify terms if there are spaces and convert to lower case
	modifiedSearchTerm := strings.ToLower(strings.ReplaceAll(searchTerm, " ", "%20"))
	modifiedRealTerm := strings.ToLower(strings.ReplaceAll(realTerm, " ", "%20"))
	// 2. Return the appropriate URL based on whether the environment is development or production

	if searchTermLat == "" || searchTermLng == "" {
		return fmt.Sprintf("/rank?searchTerm=%s&realTerm=%s", modifiedSearchTerm, modifiedRealTerm)
	}
	return fmt.Sprintf("/rank?searchTerm=%s&realTerm=%s&searchTermLat=%s&searchTermLng=%s&realTermLat=%s&realTermLng=%s", modifiedSearchTerm,
		modifiedRealTerm,
		searchTermLat,
		searchTermLng,
		realTermLat,
		realTermLng)
}
