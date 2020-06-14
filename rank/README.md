# rankmanager
--
    import "github.com/Ekram-B2/rankmanager/rank"


## Usage

#### func  HandleRequestToDetermineRank

```go
func HandleRequestToDetermineRank(rw http.ResponseWriter, req *http.Request)
```
HandleRequestToDetermineRank is the logic used to return a rank of a real term
against a search term

#### type Rank

```go
type Rank struct {
	Name string  `json:"name"`
	Rank float32 `json:"rank"`
}
```

Rank is the definition for what is retreived from the microservice
