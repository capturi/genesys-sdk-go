package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Defaultobjective
type Defaultobjective struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// TemplateId - The id of this objective's base template
	TemplateId *string `json:"templateId,omitempty"`


	// Zones - Objective zone specifies min,max points and values for the associated metric
	Zones *[]Objectivezone `json:"zones,omitempty"`


	// Enabled - A flag for whether this objective is enabled for the related metric
	Enabled *bool `json:"enabled,omitempty"`

}

func (u *Defaultobjective) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Defaultobjective

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		TemplateId *string `json:"templateId,omitempty"`
		
		Zones *[]Objectivezone `json:"zones,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		TemplateId: u.TemplateId,
		
		Zones: u.Zones,
		
		Enabled: u.Enabled,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Defaultobjective) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
