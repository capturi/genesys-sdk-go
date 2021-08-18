package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Estimatedwaittimepredictions
type Estimatedwaittimepredictions struct { 
	// Results - Returned upon a successful estimated wait time request.
	Results *[]Predictionresults `json:"results,omitempty"`

}

func (u *Estimatedwaittimepredictions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Estimatedwaittimepredictions

	

	return json.Marshal(&struct { 
		Results *[]Predictionresults `json:"results,omitempty"`
		*Alias
	}{ 
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Estimatedwaittimepredictions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
