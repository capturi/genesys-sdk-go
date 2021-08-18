package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fieldconfigs
type Fieldconfigs struct { 
	// Org
	Org *Fieldconfig `json:"org,omitempty"`


	// Person
	Person *Fieldconfig `json:"person,omitempty"`


	// Group
	Group *Fieldconfig `json:"group,omitempty"`


	// ExternalContact
	ExternalContact *Fieldconfig `json:"externalContact,omitempty"`

}

func (u *Fieldconfigs) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fieldconfigs

	

	return json.Marshal(&struct { 
		Org *Fieldconfig `json:"org,omitempty"`
		
		Person *Fieldconfig `json:"person,omitempty"`
		
		Group *Fieldconfig `json:"group,omitempty"`
		
		ExternalContact *Fieldconfig `json:"externalContact,omitempty"`
		*Alias
	}{ 
		Org: u.Org,
		
		Person: u.Person,
		
		Group: u.Group,
		
		ExternalContact: u.ExternalContact,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Fieldconfigs) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
