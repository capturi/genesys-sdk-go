package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordinguploadreportrequest
type Recordinguploadreportrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateSince - Report will include uploads since this date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSince *time.Time `json:"dateSince,omitempty"`

	// UploadStatus - Report will include uploads with this status
	UploadStatus *string `json:"uploadStatus,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordinguploadreportrequest) SetField(field string, fieldValue interface{}) {
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

func (o Recordinguploadreportrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateSince", }
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
	type Alias Recordinguploadreportrequest
	
	DateSince := new(string)
	if o.DateSince != nil {
		
		*DateSince = timeutil.Strftime(o.DateSince, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSince = nil
	}
	
	return json.Marshal(&struct { 
		DateSince *string `json:"dateSince,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		Alias
	}{ 
		DateSince: DateSince,
		
		UploadStatus: o.UploadStatus,
		Alias:    (Alias)(o),
	})
}

func (o *Recordinguploadreportrequest) UnmarshalJSON(b []byte) error {
	var RecordinguploadreportrequestMap map[string]interface{}
	err := json.Unmarshal(b, &RecordinguploadreportrequestMap)
	if err != nil {
		return err
	}
	
	if dateSinceString, ok := RecordinguploadreportrequestMap["dateSince"].(string); ok {
		DateSince, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSinceString)
		o.DateSince = &DateSince
	}
	
	if UploadStatus, ok := RecordinguploadreportrequestMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordinguploadreportrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}