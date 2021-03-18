package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkupdateshifttradestateresponse
type Bulkupdateshifttradestateresponse struct { 
	// Status - The status of the operation
	Status *string `json:"status,omitempty"`


	// OperationId - The ID for the operation
	OperationId *string `json:"operationId,omitempty"`


	// Result - The result of the operation.  Null unless status == Complete
	Result *Bulkupdateshifttradestateresult `json:"result,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
