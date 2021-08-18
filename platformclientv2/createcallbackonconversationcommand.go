package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcallbackonconversationcommand
type Createcallbackonconversationcommand struct { 
	// ScriptId - The identifier of the script to be used for the callback
	ScriptId *string `json:"scriptId,omitempty"`


	// QueueId - The identifier of the queue to be used for the callback. Either queueId or routingData is required.
	QueueId *string `json:"queueId,omitempty"`


	// RoutingData - The routing data to be used for the callback. Either queueId or routingData is required.
	RoutingData *Routingdata `json:"routingData,omitempty"`


	// CallbackUserName - The name of the party to be called back.
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// CallbackNumbers - A list of phone numbers for the callback.
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackScheduledTime - The scheduled date-time for the callback as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`


	// CountryCode - The country code to be associated with the callback numbers.
	CountryCode *string `json:"countryCode,omitempty"`


	// ValidateCallbackNumbers - Whether or not to validate the callback numbers for phone number format.
	ValidateCallbackNumbers *bool `json:"validateCallbackNumbers,omitempty"`


	// Data - A map of key-value pairs containing additional data that can be associated to the callback. These values will appear in the attributes property on the conversation participant. Example: { \"notes\": \"ready to close the deal!\", \"customerPreferredName\": \"Doc\" }
	Data *map[string]string `json:"data,omitempty"`


	// CallerId - The phone number displayed to recipients when a phone call is placed as part of the callback. Must conform to the E.164 format. May be overridden by other settings in the system such as external trunk settings. Telco support for \"callerId\" varies.
	CallerId *string `json:"callerId,omitempty"`


	// CallerIdName - The name displayed to recipients when a phone call is placed as part of the callback. May be overridden by other settings in the system such as external trunk settings. Telco support for \"callerIdName\" varies.
	CallerIdName *string `json:"callerIdName,omitempty"`

}

func (u *Createcallbackonconversationcommand) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createcallbackonconversationcommand

	
	CallbackScheduledTime := new(string)
	if u.CallbackScheduledTime != nil {
		
		*CallbackScheduledTime = timeutil.Strftime(u.CallbackScheduledTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallbackScheduledTime = nil
	}
	

	return json.Marshal(&struct { 
		ScriptId *string `json:"scriptId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		RoutingData *Routingdata `json:"routingData,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		ValidateCallbackNumbers *bool `json:"validateCallbackNumbers,omitempty"`
		
		Data *map[string]string `json:"data,omitempty"`
		
		CallerId *string `json:"callerId,omitempty"`
		
		CallerIdName *string `json:"callerIdName,omitempty"`
		*Alias
	}{ 
		ScriptId: u.ScriptId,
		
		QueueId: u.QueueId,
		
		RoutingData: u.RoutingData,
		
		CallbackUserName: u.CallbackUserName,
		
		CallbackNumbers: u.CallbackNumbers,
		
		CallbackScheduledTime: CallbackScheduledTime,
		
		CountryCode: u.CountryCode,
		
		ValidateCallbackNumbers: u.ValidateCallbackNumbers,
		
		Data: u.Data,
		
		CallerId: u.CallerId,
		
		CallerIdName: u.CallerIdName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createcallbackonconversationcommand) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
