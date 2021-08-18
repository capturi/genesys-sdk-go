package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata
type Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata struct { 
	// Version
	Version *int `json:"version,omitempty"`


	// ModifiedBy
	ModifiedBy *Wfmbushorttermforecastimportcompletetopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

}

func (u *Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata

	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Version *int `json:"version,omitempty"`
		
		ModifiedBy *Wfmbushorttermforecastimportcompletetopicuserreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		*Alias
	}{ 
		Version: u.Version,
		
		ModifiedBy: u.ModifiedBy,
		
		DateModified: DateModified,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
