package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanation
type Wfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

func (o *Wfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanation
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanation) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if startDateString, ok := WfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanationMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := WfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanationMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Status, ok := WfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanationMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedtopicrealtimeadherenceexplanation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
