package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Limitsentitylisting
type Limitsentitylisting struct { 
	// Entities
	Entities *[]Limit `json:"entities,omitempty"`

}

func (u *Limitsentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Limitsentitylisting

	

	return json.Marshal(&struct { 
		Entities *[]Limit `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Limitsentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
