package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchateventtopicconversationroutingdata
type Conversationchateventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationchateventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationchateventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationchateventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationchateventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
