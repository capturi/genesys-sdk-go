package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentcardaction - CardAction Object
type Conversationcontentcardaction struct { 
	// VarType - Describes the type of action.
	VarType *string `json:"type,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - Text to be returned as the payload from a ButtonResponse when a button is clicked.
	Payload *string `json:"payload,omitempty"`


	// Url - A URL of a web page to direct the user to.
	Url *string `json:"url,omitempty"`

}

func (o *Conversationcontentcardaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentcardaction
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Payload: o.Payload,
		
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcontentcardaction) UnmarshalJSON(b []byte) error {
	var ConversationcontentcardactionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentcardactionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationcontentcardactionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := ConversationcontentcardactionMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := ConversationcontentcardactionMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Url, ok := ConversationcontentcardactionMap["url"].(string); ok {
		o.Url = &Url
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentcardaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
