package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationaggregatequeryfilter
type Evaluationaggregatequeryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Evaluationaggregatequeryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Evaluationaggregatequerypredicate `json:"predicates,omitempty"`

}

func (u *Evaluationaggregatequeryfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationaggregatequeryfilter

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Evaluationaggregatequeryclause `json:"clauses,omitempty"`
		
		Predicates *[]Evaluationaggregatequerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Clauses: u.Clauses,
		
		Predicates: u.Predicates,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationaggregatequeryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
