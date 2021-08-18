package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Call
type Call struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// Direction - The direction of the call
	Direction *string `json:"direction,omitempty"`


	// Recording - True if this call is being recorded.
	Recording *bool `json:"recording,omitempty"`


	// RecordingState - State of recording on this call.
	RecordingState *string `json:"recordingState,omitempty"`


	// Muted - True if this call is muted so that remote participants can't hear any audio from this end.
	Muted *bool `json:"muted,omitempty"`


	// Confined - True if this call is held and the person on this side hears hold music.
	Confined *bool `json:"confined,omitempty"`


	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`


	// RecordingId - A globally unique identifier for the recording associated with this call.
	RecordingId *string `json:"recordingId,omitempty"`


	// Segments - The time line of the participant's call, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`


	// ErrorInfo
	ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime - The timestamp the call was placed on hold in the cloud clock if the call is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// DocumentId - If call is an outbound fax of a document from content management, then this is the id in content management.
	DocumentId *string `json:"documentId,omitempty"`


	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// DisconnectReasons - List of reasons that this call was disconnected. This will be set once the call disconnects.
	DisconnectReasons *[]Disconnectreason `json:"disconnectReasons,omitempty"`


	// FaxStatus - Extra information on fax transmission.
	FaxStatus *Faxstatus `json:"faxStatus,omitempty"`


	// Provider - The source provider for the call.
	Provider *string `json:"provider,omitempty"`


	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// UuiData - User to User Information (UUI) data managed by SIP session application.
	UuiData *string `json:"uuiData,omitempty"`


	// Self - Address and name data for a call endpoint.
	Self *Address `json:"self,omitempty"`


	// Other - Address and name data for a call endpoint.
	Other *Address `json:"other,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AgentAssistantId - UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`

}

func (u *Call) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Call

	
	StartHoldTime := new(string)
	if u.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(u.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	StartAlertingTime := new(string)
	if u.StartAlertingTime != nil {
		
		*StartAlertingTime = timeutil.Strftime(u.StartAlertingTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAlertingTime = nil
	}
	
	ConnectedTime := new(string)
	if u.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(u.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if u.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(u.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	

	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		RecordingId *string `json:"recordingId,omitempty"`
		
		Segments *[]Segment `json:"segments,omitempty"`
		
		ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		StartAlertingTime *string `json:"startAlertingTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		DisconnectReasons *[]Disconnectreason `json:"disconnectReasons,omitempty"`
		
		FaxStatus *Faxstatus `json:"faxStatus,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		UuiData *string `json:"uuiData,omitempty"`
		
		Self *Address `json:"self,omitempty"`
		
		Other *Address `json:"other,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		*Alias
	}{ 
		State: u.State,
		
		Id: u.Id,
		
		Direction: u.Direction,
		
		Recording: u.Recording,
		
		RecordingState: u.RecordingState,
		
		Muted: u.Muted,
		
		Confined: u.Confined,
		
		Held: u.Held,
		
		RecordingId: u.RecordingId,
		
		Segments: u.Segments,
		
		ErrorInfo: u.ErrorInfo,
		
		DisconnectType: u.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		DocumentId: u.DocumentId,
		
		StartAlertingTime: StartAlertingTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		DisconnectReasons: u.DisconnectReasons,
		
		FaxStatus: u.FaxStatus,
		
		Provider: u.Provider,
		
		ScriptId: u.ScriptId,
		
		PeerId: u.PeerId,
		
		UuiData: u.UuiData,
		
		Self: u.Self,
		
		Other: u.Other,
		
		Wrapup: u.Wrapup,
		
		AfterCallWork: u.AfterCallWork,
		
		AfterCallWorkRequired: u.AfterCallWorkRequired,
		
		AgentAssistantId: u.AgentAssistantId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Call) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
