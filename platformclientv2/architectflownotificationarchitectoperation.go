package platformclientv2

import (
	"encoding/json"
	"github.com/leekchan/timeutil"
	"reflect"
	"strconv"
	"strings"
)

// Architectflownotificationarchitectoperation - Describes an operation being performed on an Architect object
type Architectflownotificationarchitectoperation struct {
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A unique identifier for this operation, as generated by the initiating Client
	Id *string `json:"id,omitempty"`

	// Complete - Indicates if the operation is complete
	Complete *bool `json:"complete,omitempty"`

	// User
	User *Architectflownotificationuser `json:"user,omitempty"`

	// Client
	Client *Architectflownotificationclient `json:"Client,omitempty"`

	// ActionName - The action being performed
	ActionName *string `json:"actionName,omitempty"`

	// ActionStatus - The action status
	ActionStatus *string `json:"actionStatus,omitempty"`

	// ErrorMessage - The error message, if the action failed
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// ErrorCode - The error code, if the action failed
	ErrorCode *string `json:"errorCode,omitempty"`

	// ErrorMessageParams
	ErrorMessageParams *Architectflownotificationerrormessageparams `json:"errorMessageParams,omitempty"`

	// ErrorDetails - The error details, if the action failed
	ErrorDetails *[]Architectflownotificationerrordetail `json:"errorDetails,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Architectflownotificationarchitectoperation) SetField(field string, fieldValue interface{}) {
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

func (o Architectflownotificationarchitectoperation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{}
		localDateTimeFields := []string{}
		dateFields := []string{}

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil() {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
	_ = timeutil.Timedelta{}
	type Alias Architectflownotificationarchitectoperation

	return json.Marshal(&struct {
		Id *string `json:"id,omitempty"`

		Complete *bool `json:"complete,omitempty"`

		User *Architectflownotificationuser `json:"user,omitempty"`

		Client *Architectflownotificationclient `json:"Client,omitempty"`

		ActionName *string `json:"actionName,omitempty"`

		ActionStatus *string `json:"actionStatus,omitempty"`

		ErrorMessage *string `json:"errorMessage,omitempty"`

		ErrorCode *string `json:"errorCode,omitempty"`

		ErrorMessageParams *Architectflownotificationerrormessageparams `json:"errorMessageParams,omitempty"`

		ErrorDetails *[]Architectflownotificationerrordetail `json:"errorDetails,omitempty"`
		Alias
	}{
		Id: o.Id,

		Complete: o.Complete,

		User: o.User,

		Client: o.Client,

		ActionName: o.ActionName,

		ActionStatus: o.ActionStatus,

		ErrorMessage: o.ErrorMessage,

		ErrorCode: o.ErrorCode,

		ErrorMessageParams: o.ErrorMessageParams,

		ErrorDetails: o.ErrorDetails,
		Alias:        (Alias)(o),
	})
}

func (o *Architectflownotificationarchitectoperation) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationarchitectoperationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationarchitectoperationMap)
	if err != nil {
		return err
	}

	if Id, ok := ArchitectflownotificationarchitectoperationMap["id"].(string); ok {
		o.Id = &Id
	}

	if Complete, ok := ArchitectflownotificationarchitectoperationMap["complete"].(bool); ok {
		o.Complete = &Complete
	}

	if User, ok := ArchitectflownotificationarchitectoperationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}

	if Client, ok := ArchitectflownotificationarchitectoperationMap["Client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}

	if ActionName, ok := ArchitectflownotificationarchitectoperationMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}

	if ActionStatus, ok := ArchitectflownotificationarchitectoperationMap["actionStatus"].(string); ok {
		o.ActionStatus = &ActionStatus
	}

	if ErrorMessage, ok := ArchitectflownotificationarchitectoperationMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}

	if ErrorCode, ok := ArchitectflownotificationarchitectoperationMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}

	if ErrorMessageParams, ok := ArchitectflownotificationarchitectoperationMap["errorMessageParams"].(map[string]interface{}); ok {
		ErrorMessageParamsString, _ := json.Marshal(ErrorMessageParams)
		json.Unmarshal(ErrorMessageParamsString, &o.ErrorMessageParams)
	}

	if ErrorDetails, ok := ArchitectflownotificationarchitectoperationMap["errorDetails"].([]interface{}); ok {
		ErrorDetailsString, _ := json.Marshal(ErrorDetails)
		json.Unmarshal(ErrorDetailsString, &o.ErrorDetails)
	}

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationarchitectoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
