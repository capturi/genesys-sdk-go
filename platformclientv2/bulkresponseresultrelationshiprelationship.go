package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultrelationshiprelationship
type Bulkresponseresultrelationshiprelationship struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Relationship `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorrelationship `json:"error,omitempty"`

}

func (u *Bulkresponseresultrelationshiprelationship) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkresponseresultrelationshiprelationship

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Entity *Relationship `json:"entity,omitempty"`
		
		VarError *Bulkerrorrelationship `json:"error,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Success: u.Success,
		
		Entity: u.Entity,
		
		VarError: u.VarError,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bulkresponseresultrelationshiprelationship) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
