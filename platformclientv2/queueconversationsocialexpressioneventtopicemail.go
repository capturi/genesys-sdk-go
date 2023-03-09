package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicemail
type Queueconversationsocialexpressioneventtopicemail struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// InitialState
	InitialState *string `json:"initialState,omitempty"`

	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`

	// AutoGenerated - Indicates that the email was auto-generated like an Out of Office reply.
	AutoGenerated *bool `json:"autoGenerated,omitempty"`

	// Subject - The subject for the initial email that started this conversation.
	Subject *string `json:"subject,omitempty"`

	// Provider - The source provider of the email.
	Provider *string `json:"provider,omitempty"`

	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`

	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`

	// MessagesSent - The number of email messages sent by this participant.
	MessagesSent *int `json:"messagesSent,omitempty"`

	// ErrorInfo - Detailed information about an error response.
	ErrorInfo *Queueconversationsocialexpressioneventtopicerrordetails `json:"errorInfo,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// StartHoldTime - The timestamp the email was placed on hold in the cloud clock if the email is currently on hold.
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// MessageId - A globally unique identifier for the stored content of this communication.
	MessageId *string `json:"messageId,omitempty"`

	// Direction - Whether an email is inbound or outbound.
	Direction *string `json:"direction,omitempty"`

	// DraftAttachments - A list of uploaded attachments on the email draft.
	DraftAttachments *[]Queueconversationsocialexpressioneventtopicattachment `json:"draftAttachments,omitempty"`

	// Spam - Indicates if the inbound email was marked as spam.
	Spam *bool `json:"spam,omitempty"`

	// Wrapup - Call wrap up or disposition data.
	Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`

	// AfterCallWork - A communication's after-call work data.
	AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`

	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationsocialexpressioneventtopicemail) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationsocialexpressioneventtopicemail) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartHoldTime","ConnectedTime","DisconnectedTime", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
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
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicemail
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if o.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(o.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		AutoGenerated *bool `json:"autoGenerated,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		MessagesSent *int `json:"messagesSent,omitempty"`
		
		ErrorInfo *Queueconversationsocialexpressioneventtopicerrordetails `json:"errorInfo,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DraftAttachments *[]Queueconversationsocialexpressioneventtopicattachment `json:"draftAttachments,omitempty"`
		
		Spam *bool `json:"spam,omitempty"`
		
		Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		InitialState: o.InitialState,
		
		Held: o.Held,
		
		AutoGenerated: o.AutoGenerated,
		
		Subject: o.Subject,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		MessagesSent: o.MessagesSent,
		
		ErrorInfo: o.ErrorInfo,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		MessageId: o.MessageId,
		
		Direction: o.Direction,
		
		DraftAttachments: o.DraftAttachments,
		
		Spam: o.Spam,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		Alias:    (Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicemail) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicemailMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicemailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicemailMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := QueueconversationsocialexpressioneventtopicemailMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := QueueconversationsocialexpressioneventtopicemailMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Held, ok := QueueconversationsocialexpressioneventtopicemailMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if AutoGenerated, ok := QueueconversationsocialexpressioneventtopicemailMap["autoGenerated"].(bool); ok {
		o.AutoGenerated = &AutoGenerated
	}
    
	if Subject, ok := QueueconversationsocialexpressioneventtopicemailMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if Provider, ok := QueueconversationsocialexpressioneventtopicemailMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ScriptId, ok := QueueconversationsocialexpressioneventtopicemailMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if PeerId, ok := QueueconversationsocialexpressioneventtopicemailMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if MessagesSent, ok := QueueconversationsocialexpressioneventtopicemailMap["messagesSent"].(float64); ok {
		MessagesSentInt := int(MessagesSent)
		o.MessagesSent = &MessagesSentInt
	}
	
	if ErrorInfo, ok := QueueconversationsocialexpressioneventtopicemailMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if DisconnectType, ok := QueueconversationsocialexpressioneventtopicemailMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if startHoldTimeString, ok := QueueconversationsocialexpressioneventtopicemailMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if connectedTimeString, ok := QueueconversationsocialexpressioneventtopicemailMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationsocialexpressioneventtopicemailMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if MessageId, ok := QueueconversationsocialexpressioneventtopicemailMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if Direction, ok := QueueconversationsocialexpressioneventtopicemailMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DraftAttachments, ok := QueueconversationsocialexpressioneventtopicemailMap["draftAttachments"].([]interface{}); ok {
		DraftAttachmentsString, _ := json.Marshal(DraftAttachments)
		json.Unmarshal(DraftAttachmentsString, &o.DraftAttachments)
	}
	
	if Spam, ok := QueueconversationsocialexpressioneventtopicemailMap["spam"].(bool); ok {
		o.Spam = &Spam
	}
    
	if Wrapup, ok := QueueconversationsocialexpressioneventtopicemailMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationsocialexpressioneventtopicemailMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationsocialexpressioneventtopicemailMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
