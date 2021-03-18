package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Criteria
type Criteria struct { 
	// Key - The criteria key.
	Key *string `json:"key,omitempty"`


	// Values - The criteria values.
	Values *[]string `json:"values,omitempty"`


	// ShouldIgnoreCase - Should criteria be case insensitive.
	ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`

}

// String returns a JSON representation of the model
func (o *Criteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
