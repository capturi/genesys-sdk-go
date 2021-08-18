package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Managementunit - Management Unit object for Workforce Management
type Managementunit struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// BusinessUnit - The business unit to which this management unit belongs
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`


	// StartDayOfWeek - Start day of week for scheduling and forecasting purposes. Moving to Business Unit
	StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`


	// TimeZone - The time zone for the management unit in standard Olson format.  Moving to Business Unit
	TimeZone *string `json:"timeZone,omitempty"`


	// Settings - The configuration settings for this management unit
	Settings *Managementunitsettingsresponse `json:"settings,omitempty"`


	// Metadata - Version info metadata for this management unit. Deprecated, use settings.metadata
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Divisionreference `json:"division,omitempty"`


	// Version - The version of the underlying entity.  Deprecated, use field from settings.metadata instead
	Version *int `json:"version,omitempty"`


	// DateModified - The date and time at which this entity was last modified.  Deprecated, use field from settings.metadata instead. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The user who last modified this entity.  Deprecated, use field from settings.metadata instead
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Managementunit) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Managementunit

	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		Settings *Managementunitsettingsresponse `json:"settings,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		Division *Divisionreference `json:"division,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		BusinessUnit: u.BusinessUnit,
		
		StartDayOfWeek: u.StartDayOfWeek,
		
		TimeZone: u.TimeZone,
		
		Settings: u.Settings,
		
		Metadata: u.Metadata,
		
		Division: u.Division,
		
		Version: u.Version,
		
		DateModified: DateModified,
		
		ModifiedBy: u.ModifiedBy,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Managementunit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
