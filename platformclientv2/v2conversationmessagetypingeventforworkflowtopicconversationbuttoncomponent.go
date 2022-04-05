package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponent
type V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponent struct { 
	// Title
	Title *string `json:"title,omitempty"`


	// Actions
	Actions *V2conversationmessagetypingeventforworkflowtopicconversationcontentactions `json:"actions,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponent
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Actions *V2conversationmessagetypingeventforworkflowtopicconversationcontentactions `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponent) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponentMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponentMap)
	if err != nil {
		return err
	}
	
	if Title, ok := V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponentMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Actions, ok := V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponentMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationbuttoncomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}