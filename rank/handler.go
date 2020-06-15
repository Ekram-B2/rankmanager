package rankmanager

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	l4g "github.com/alecthomas/log4go"

	"github.com/Ekram-B2/suggestionsmanager/config"
)

type responseFormat struct {
	Name string  `json:"name"`
	Rank float64 `json:"rank"`
}

// HandleRequestToDetermineRank is the logic used to return a rank of a real term against a search term
func HandleRequestToDetermineRank(rw http.ResponseWriter, req *http.Request) {
	// 1. Check to see if within the request made, there is a query parameter containing the search term
	searchTerm := req.URL.Query().Get("searchTerm")
	if searchTerm == "" {
		l4g.Error("SYSTEM-ERROR: 'searchTerm' was not found as a query paramter")
		http.Error(rw, "We were unable to find the required search term for this request; please include a 'searchTerm' parameter in your next request.", http.StatusBadRequest)
		return
	}

	// 2. See if the real term is provided
	realTerm := req.URL.Query().Get("realTerm")
	if realTerm == "" {
		l4g.Error("SYSTEM-ERROR: 'realTerm' was not found as a query parameter")
		http.Error(rw, "we were unable to find the required real term parameter for this request; please include a 'realTerm' parameter in your next request.", http.StatusBadRequest)
		return
	}

	// 3. Check if latitudes and longitudes for the realTerm and searchTerm are provided. If these params aren't
	// provided, then this does not mean that there is an error as is it possible to query without this information
	searchTermLat := req.URL.Query().Get("searchTermLat")

	searchTermLng := req.URL.Query().Get("searchTermLng")

	realTermLat := req.URL.Query().Get("realTermLat")

	realTermLng := req.URL.Query().Get("realTermLng")

	// 4. Load configuration
	loadedConfig, err := config.LoadConfiguration(config.GetConfigPath(os.Getenv("CONFIG_OPERATION_TYPE")))
	if err != nil {
		l4g.Error("OPERATION-ERROR: there was an error loading the configuration object: %s", err.Error())
		http.Error(rw, "We were unable to retreive a rank for the real term provided; please try again after waiting some time.", http.StatusInternalServerError)
		return
	}
	// 5. build operation to apply onto score
	getRank := getRank(convertStringToFloat64(searchTermLat),
		convertStringToFloat64(searchTermLng),
		convertStringToFloat64(realTermLat),
		convertStringToFloat64(realTermLng),
		searchTerm,
		realTerm,
		getCharDistanceRanker(loadedConfig.DistanceRankerType),
		getLatLngDistanceRanker(loadedConfig.LatLngDistCalcType),
		getNormalizer(loadedConfig.Normalizer))

	rank := getRank(searchTerm, realTerm, getNormalizer(loadedConfig.Normalizer))

	// 6. Set up reply format to send back to caller
	content := responseFormat{Name: realTerm, Rank: rank}

	// 7. Return response back to caller
	rw.Header().Add("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)

	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(content)
	if err != nil {
		l4g.Error("OPERATION-ERROR: there was an error marshalling to the expected output: %s", err.Error())
		http.Error(rw, "We were unable to retreive a rank for the real term provided; please try again after waiting some time.", http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(b.Bytes())
	if err != nil {
		l4g.Error("OPERATION-ERROR: there was an error encoding content within the response writer: %s", err.Error())
		http.Error(rw, "We were unable to return the rank for this realTerm; please try again later after waiting some time.", http.StatusInternalServerError)
		return
	}
}

func convertStringToFloat64(loc string) float64 {
	value, err := strconv.ParseFloat(loc, 64)

	if err != nil {
		return 0.0
	}
	return float64(value)
}
