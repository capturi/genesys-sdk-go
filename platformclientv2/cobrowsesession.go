package platformclientv2

import (
	"encoding/json"
	"github.com/leekchan/timeutil"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Cobrowsesession
type Cobrowsesession struct {
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`

	// InitialState - The initial connection state of this communication.
	InitialState *string `json:"initialState,omitempty"`

	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// Self - Address and name data for a call endpoint.
	Self *Address `json:"self,omitempty"`

	// CobrowseSessionId - The co-browse session ID.
	CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`

	// CobrowseRole - This value identifies the role of the co-browse Client within the co-browse session (a Client is a sharer or a viewer).
	CobrowseRole *string `json:"cobrowseRole,omitempty"`

	// Controlling - ID of co-browse participants for which this Client has been granted control (list is empty if this Client cannot control any shared pages).
	Controlling *[]string `json:"controlling,omitempty"`

	// ViewerUrl - The URL that can be used to open co-browse session in web browser.
	ViewerUrl *string `json:"viewerUrl,omitempty"`

	// ProviderEventTime - The time when the provider event which triggered this conversation update happened in the corrected provider clock (milliseconds since 1970-01-01 00:00:00 UTC). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ProviderEventTime *time.Time `json:"providerEventTime,omitempty"`

	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// Provider - The source provider for the co-browse session.
	Provider *string `json:"provider,omitempty"`

	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`

	// Segments - The time line of the participant's call, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`

	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`

	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`

	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Cobrowsesession) SetField(field string, fieldValue interface{}) {
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

func (o Cobrowsesession) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{"ProviderEventTime", "StartAlertingTime", "ConnectedTime", "DisconnectedTime"}
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
	type Alias Cobrowsesession

	ProviderEventTime := new(string)
	if o.ProviderEventTime != nil {

		*ProviderEventTime = timeutil.Strftime(o.ProviderEventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ProviderEventTime = nil
	}

	StartAlertingTime := new(string)
	if o.StartAlertingTime != nil {

		*StartAlertingTime = timeutil.Strftime(o.StartAlertingTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAlertingTime = nil
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
		State *string `json:"state,omitempty"`

		InitialState *string `json:"initialState,omitempty"`

		Id *string `json:"id,omitempty"`

		DisconnectType *string `json:"disconnectType,omitempty"`

		Self *Address `json:"self,omitempty"`

		CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`

		CobrowseRole *string `json:"cobrowseRole,omitempty"`

		Controlling *[]string `json:"controlling,omitempty"`

		ViewerUrl *string `json:"viewerUrl,omitempty"`

		ProviderEventTime *string `json:"providerEventTime,omitempty"`

		StartAlertingTime *string `json:"startAlertingTime,omitempty"`

		ConnectedTime *string `json:"connectedTime,omitempty"`

		DisconnectedTime *string `json:"disconnectedTime,omitempty"`

		Provider *string `json:"provider,omitempty"`

		PeerId *string `json:"peerId,omitempty"`

		Segments *[]Segment `json:"segments,omitempty"`

		Wrapup *Wrapup `json:"wrapup,omitempty"`

		AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`

		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		Alias
	}{
		State: o.State,

		InitialState: o.InitialState,

		Id: o.Id,

		DisconnectType: o.DisconnectType,

		Self: o.Self,

		CobrowseSessionId: o.CobrowseSessionId,

		CobrowseRole: o.CobrowseRole,

		Controlling: o.Controlling,

		ViewerUrl: o.ViewerUrl,

		ProviderEventTime: ProviderEventTime,

		StartAlertingTime: StartAlertingTime,

		ConnectedTime: ConnectedTime,

		DisconnectedTime: DisconnectedTime,

		Provider: o.Provider,

		PeerId: o.PeerId,

		Segments: o.Segments,

		Wrapup: o.Wrapup,

		AfterCallWork: o.AfterCallWork,

		AfterCallWorkRequired: o.AfterCallWorkRequired,
		Alias:                 (Alias)(o),
	})
}

func (o *Cobrowsesession) UnmarshalJSON(b []byte) error {
	var CobrowsesessionMap map[string]interface{}
	err := json.Unmarshal(b, &CobrowsesessionMap)
	if err != nil {
		return err
	}

	if State, ok := CobrowsesessionMap["state"].(string); ok {
		o.State = &State
	}

	if InitialState, ok := CobrowsesessionMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}

	if Id, ok := CobrowsesessionMap["id"].(string); ok {
		o.Id = &Id
	}

	if DisconnectType, ok := CobrowsesessionMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}

	if Self, ok := CobrowsesessionMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}

	if CobrowseSessionId, ok := CobrowsesessionMap["cobrowseSessionId"].(string); ok {
		o.CobrowseSessionId = &CobrowseSessionId
	}

	if CobrowseRole, ok := CobrowsesessionMap["cobrowseRole"].(string); ok {
		o.CobrowseRole = &CobrowseRole
	}

	if Controlling, ok := CobrowsesessionMap["controlling"].([]interface{}); ok {
		ControllingString, _ := json.Marshal(Controlling)
		json.Unmarshal(ControllingString, &o.Controlling)
	}

	if ViewerUrl, ok := CobrowsesessionMap["viewerUrl"].(string); ok {
		o.ViewerUrl = &ViewerUrl
	}

	if providerEventTimeString, ok := CobrowsesessionMap["providerEventTime"].(string); ok {
		ProviderEventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", providerEventTimeString)
		o.ProviderEventTime = &ProviderEventTime
	}

	if startAlertingTimeString, ok := CobrowsesessionMap["startAlertingTime"].(string); ok {
		StartAlertingTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAlertingTimeString)
		o.StartAlertingTime = &StartAlertingTime
	}

	if connectedTimeString, ok := CobrowsesessionMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}

	if disconnectedTimeString, ok := CobrowsesessionMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}

	if Provider, ok := CobrowsesessionMap["provider"].(string); ok {
		o.Provider = &Provider
	}

	if PeerId, ok := CobrowsesessionMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}

	if Segments, ok := CobrowsesessionMap["segments"].([]interface{}); ok {
		SegmentsString, _ := json.Marshal(Segments)
		json.Unmarshal(SegmentsString, &o.Segments)
	}

	if Wrapup, ok := CobrowsesessionMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}

	if AfterCallWork, ok := CobrowsesessionMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}

	if AfterCallWorkRequired, ok := CobrowsesessionMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}

	return nil
}

// String returns a JSON representation of the model
func (o *Cobrowsesession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
