package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkerrorrelationship
type Bulkerrorrelationship struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// Retryable
	Retryable *bool `json:"retryable,omitempty"`


	// Entity
	Entity *Relationship `json:"entity,omitempty"`


	// Details
	Details *[]Bulkerrordetail `json:"details,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkerrorrelationship) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}