package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopicjourneycustomersession
type Conversationcobrowseeventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
