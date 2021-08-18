package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coversheet
type Coversheet struct { 
	// Notes - Text to be added to the coversheet
	Notes *string `json:"notes,omitempty"`


	// Locale - Locale, e.g. = en-US
	Locale *string `json:"locale,omitempty"`

}

func (u *Coversheet) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coversheet

	

	return json.Marshal(&struct { 
		Notes *string `json:"notes,omitempty"`
		
		Locale *string `json:"locale,omitempty"`
		*Alias
	}{ 
		Notes: u.Notes,
		
		Locale: u.Locale,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Coversheet) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
