package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingstatusdetailqueryclause
type Routingstatusdetailqueryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Routingstatusdetailquerypredicate `json:"predicates,omitempty"`

}

func (u *Routingstatusdetailqueryclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingstatusdetailqueryclause

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Predicates *[]Routingstatusdetailquerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Predicates: u.Predicates,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Routingstatusdetailqueryclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
