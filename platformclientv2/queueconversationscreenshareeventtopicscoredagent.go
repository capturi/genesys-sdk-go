package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicscoredagent
type Queueconversationscreenshareeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationscreenshareeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

func (u *Queueconversationscreenshareeventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationscreenshareeventtopicscoredagent

	

	return json.Marshal(&struct { 
		Agent *Queueconversationscreenshareeventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: u.Agent,
		
		Score: u.Score,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
