package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicerrorbody
type Queueconversationscreenshareeventtopicerrorbody struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Details
	Details *[]Queueconversationscreenshareeventtopicdetail `json:"details,omitempty"`


	// Errors
	Errors *[]Queueconversationscreenshareeventtopicerrorbody `json:"errors,omitempty"`

}

func (u *Queueconversationscreenshareeventtopicerrorbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationscreenshareeventtopicerrorbody

	

	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Details *[]Queueconversationscreenshareeventtopicdetail `json:"details,omitempty"`
		
		Errors *[]Queueconversationscreenshareeventtopicerrorbody `json:"errors,omitempty"`
		*Alias
	}{ 
		Message: u.Message,
		
		Code: u.Code,
		
		Status: u.Status,
		
		EntityId: u.EntityId,
		
		EntityName: u.EntityName,
		
		MessageWithParams: u.MessageWithParams,
		
		MessageParams: u.MessageParams,
		
		ContextId: u.ContextId,
		
		Details: u.Details,
		
		Errors: u.Errors,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicerrorbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
