package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Qmauditqueryrequest
type Qmauditqueryrequest struct { 
	// Interval - Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Filters - List of filters for the query.
	Filters *[]Qualityauditqueryfilter `json:"filters,omitempty"`


	// Sort - Sort parameter for the query.
	Sort *[]Auditquerysort `json:"sort,omitempty"`

}

func (u *Qmauditqueryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Qmauditqueryrequest

	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Filters *[]Qualityauditqueryfilter `json:"filters,omitempty"`
		
		Sort *[]Auditquerysort `json:"sort,omitempty"`
		*Alias
	}{ 
		Interval: u.Interval,
		
		Filters: u.Filters,
		
		Sort: u.Sort,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Qmauditqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
