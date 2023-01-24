package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmtimeoffbalancejobtopictimeoffbalance
type Wfmtimeoffbalancejobtopictimeoffbalance struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// HrisTimeOffTypeId
	HrisTimeOffTypeId *string `json:"hrisTimeOffTypeId,omitempty"`

	// HrisTimeOffTypeSecondaryId
	HrisTimeOffTypeSecondaryId *string `json:"hrisTimeOffTypeSecondaryId,omitempty"`

	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`

	// BalanceMinutesPerDay
	BalanceMinutesPerDay *[]int `json:"balanceMinutesPerDay,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmtimeoffbalancejobtopictimeoffbalance) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Wfmtimeoffbalancejobtopictimeoffbalance) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmtimeoffbalancejobtopictimeoffbalance
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		HrisTimeOffTypeId *string `json:"hrisTimeOffTypeId,omitempty"`
		
		HrisTimeOffTypeSecondaryId *string `json:"hrisTimeOffTypeSecondaryId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		BalanceMinutesPerDay *[]int `json:"balanceMinutesPerDay,omitempty"`
		Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		HrisTimeOffTypeId: o.HrisTimeOffTypeId,
		
		HrisTimeOffTypeSecondaryId: o.HrisTimeOffTypeSecondaryId,
		
		StartDate: StartDate,
		
		BalanceMinutesPerDay: o.BalanceMinutesPerDay,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmtimeoffbalancejobtopictimeoffbalance) UnmarshalJSON(b []byte) error {
	var WfmtimeoffbalancejobtopictimeoffbalanceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmtimeoffbalancejobtopictimeoffbalanceMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if HrisTimeOffTypeId, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["hrisTimeOffTypeId"].(string); ok {
		o.HrisTimeOffTypeId = &HrisTimeOffTypeId
	}
    
	if HrisTimeOffTypeSecondaryId, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["hrisTimeOffTypeSecondaryId"].(string); ok {
		o.HrisTimeOffTypeSecondaryId = &HrisTimeOffTypeSecondaryId
	}
    
	if startDateString, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if BalanceMinutesPerDay, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["balanceMinutesPerDay"].([]interface{}); ok {
		BalanceMinutesPerDayString, _ := json.Marshal(BalanceMinutesPerDay)
		json.Unmarshal(BalanceMinutesPerDayString, &o.BalanceMinutesPerDay)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmtimeoffbalancejobtopictimeoffbalance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}