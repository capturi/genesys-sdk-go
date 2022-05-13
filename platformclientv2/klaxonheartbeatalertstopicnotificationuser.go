package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxonheartbeatalertstopicnotificationuser
type Klaxonheartbeatalertstopicnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Klaxonheartbeatalertstopicnotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxonheartbeatalertstopicnotificationuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Klaxonheartbeatalertstopicnotificationuser) UnmarshalJSON(b []byte) error {
	var KlaxonheartbeatalertstopicnotificationuserMap map[string]interface{}
	err := json.Unmarshal(b, &KlaxonheartbeatalertstopicnotificationuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KlaxonheartbeatalertstopicnotificationuserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DisplayName, ok := KlaxonheartbeatalertstopicnotificationuserMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatalertstopicnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
