package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicscreenshare
type Queueconversationvideoeventtopicscreenshare struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// Self - Address and name data for a call endpoint.
	Self *Queueconversationvideoeventtopicaddress `json:"self,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// Context - The room id context (xmpp jid) for the conference session.
	Context *string `json:"context,omitempty"`


	// Sharing - Indicates whether this participant is sharing their screen to the session.
	Sharing *bool `json:"sharing,omitempty"`


	// Provider - The source provider of the screen share.
	Provider *string `json:"provider,omitempty"`


	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// PeerCount - The number of peer participants from the perspective of the participant in the conference.
	PeerCount *interface{} `json:"peerCount,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork - A communication's after-call work data.
	AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

}

func (o *Queueconversationvideoeventtopicscreenshare) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicscreenshare
	
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
		
		Self *Queueconversationvideoeventtopicaddress `json:"self,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Context *string `json:"context,omitempty"`
		
		Sharing *bool `json:"sharing,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		PeerCount *interface{} `json:"peerCount,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		Self: o.Self,
		
		Id: o.Id,
		
		Context: o.Context,
		
		Sharing: o.Sharing,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		PeerCount: o.PeerCount,
		
		DisconnectType: o.DisconnectType,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicscreenshare) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicscreenshareMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicscreenshareMap)
	if err != nil {
		return err
	}
	
	if State, ok := QueueconversationvideoeventtopicscreenshareMap["state"].(string); ok {
		o.State = &State
	}
    
	if Self, ok := QueueconversationvideoeventtopicscreenshareMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if Id, ok := QueueconversationvideoeventtopicscreenshareMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Context, ok := QueueconversationvideoeventtopicscreenshareMap["context"].(string); ok {
		o.Context = &Context
	}
    
	if Sharing, ok := QueueconversationvideoeventtopicscreenshareMap["sharing"].(bool); ok {
		o.Sharing = &Sharing
	}
    
	if Provider, ok := QueueconversationvideoeventtopicscreenshareMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ScriptId, ok := QueueconversationvideoeventtopicscreenshareMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if PeerId, ok := QueueconversationvideoeventtopicscreenshareMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if PeerCount, ok := QueueconversationvideoeventtopicscreenshareMap["peerCount"].(map[string]interface{}); ok {
		PeerCountString, _ := json.Marshal(PeerCount)
		json.Unmarshal(PeerCountString, &o.PeerCount)
	}
	
	if DisconnectType, ok := QueueconversationvideoeventtopicscreenshareMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if connectedTimeString, ok := QueueconversationvideoeventtopicscreenshareMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationvideoeventtopicscreenshareMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Wrapup, ok := QueueconversationvideoeventtopicscreenshareMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationvideoeventtopicscreenshareMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationvideoeventtopicscreenshareMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicscreenshare) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
