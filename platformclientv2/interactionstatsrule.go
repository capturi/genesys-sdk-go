package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Interactionstatsrule
type Interactionstatsrule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Name of the rule
	Name *string `json:"name,omitempty"`


	// Dimension - The dimension of concern.
	Dimension *string `json:"dimension,omitempty"`


	// DimensionValue - The value of the dimension.
	DimensionValue *string `json:"dimensionValue,omitempty"`


	// Metric - The metric to be assessed.
	Metric *string `json:"metric,omitempty"`


	// MediaType - The media type.
	MediaType *string `json:"mediaType,omitempty"`


	// NumericRange - The comparison descriptor used against the metric's value.
	NumericRange *string `json:"numericRange,omitempty"`


	// Statistic - The statistic of concern for the metric.
	Statistic *string `json:"statistic,omitempty"`


	// Value - The threshold value.
	Value *float64 `json:"value,omitempty"`


	// Enabled - Indicates if the rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`


	// InAlarm - Indicates if the rule is in alarm state.
	InAlarm *bool `json:"inAlarm,omitempty"`


	// NotificationUsers - The ids of users who will be notified of alarm state change.
	NotificationUsers *[]User `json:"notificationUsers,omitempty"`


	// AlertTypes - A collection of notification methods.
	AlertTypes *[]string `json:"alertTypes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Interactionstatsrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Interactionstatsrule
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		DimensionValue *string `json:"dimensionValue,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		NumericRange *string `json:"numericRange,omitempty"`
		
		Statistic *string `json:"statistic,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		InAlarm *bool `json:"inAlarm,omitempty"`
		
		NotificationUsers *[]User `json:"notificationUsers,omitempty"`
		
		AlertTypes *[]string `json:"alertTypes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Dimension: o.Dimension,
		
		DimensionValue: o.DimensionValue,
		
		Metric: o.Metric,
		
		MediaType: o.MediaType,
		
		NumericRange: o.NumericRange,
		
		Statistic: o.Statistic,
		
		Value: o.Value,
		
		Enabled: o.Enabled,
		
		InAlarm: o.InAlarm,
		
		NotificationUsers: o.NotificationUsers,
		
		AlertTypes: o.AlertTypes,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Interactionstatsrule) UnmarshalJSON(b []byte) error {
	var InteractionstatsruleMap map[string]interface{}
	err := json.Unmarshal(b, &InteractionstatsruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := InteractionstatsruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := InteractionstatsruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Dimension, ok := InteractionstatsruleMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if DimensionValue, ok := InteractionstatsruleMap["dimensionValue"].(string); ok {
		o.DimensionValue = &DimensionValue
	}
    
	if Metric, ok := InteractionstatsruleMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if MediaType, ok := InteractionstatsruleMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if NumericRange, ok := InteractionstatsruleMap["numericRange"].(string); ok {
		o.NumericRange = &NumericRange
	}
    
	if Statistic, ok := InteractionstatsruleMap["statistic"].(string); ok {
		o.Statistic = &Statistic
	}
    
	if Value, ok := InteractionstatsruleMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Enabled, ok := InteractionstatsruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if InAlarm, ok := InteractionstatsruleMap["inAlarm"].(bool); ok {
		o.InAlarm = &InAlarm
	}
    
	if NotificationUsers, ok := InteractionstatsruleMap["notificationUsers"].([]interface{}); ok {
		NotificationUsersString, _ := json.Marshal(NotificationUsers)
		json.Unmarshal(NotificationUsersString, &o.NotificationUsers)
	}
	
	if AlertTypes, ok := InteractionstatsruleMap["alertTypes"].([]interface{}); ok {
		AlertTypesString, _ := json.Marshal(AlertTypes)
		json.Unmarshal(AlertTypesString, &o.AlertTypes)
	}
	
	if SelfUri, ok := InteractionstatsruleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Interactionstatsrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
