package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationgeolocation
type Journeyoutcomeeventsnotificationgeolocation struct { 
	// Country
	Country *string `json:"country,omitempty"`


	// CountryName
	CountryName *string `json:"countryName,omitempty"`


	// Latitude
	Latitude *float32 `json:"latitude,omitempty"`


	// Longitude
	Longitude *float32 `json:"longitude,omitempty"`


	// Locality
	Locality *string `json:"locality,omitempty"`


	// PostalCode
	PostalCode *string `json:"postalCode,omitempty"`


	// Region
	Region *string `json:"region,omitempty"`


	// RegionName
	RegionName *string `json:"regionName,omitempty"`


	// Timezone
	Timezone *string `json:"timezone,omitempty"`


	// Source
	Source *string `json:"source,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationgeolocation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationgeolocation
	
	return json.Marshal(&struct { 
		Country *string `json:"country,omitempty"`
		
		CountryName *string `json:"countryName,omitempty"`
		
		Latitude *float32 `json:"latitude,omitempty"`
		
		Longitude *float32 `json:"longitude,omitempty"`
		
		Locality *string `json:"locality,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		RegionName *string `json:"regionName,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
		
		Source *string `json:"source,omitempty"`
		*Alias
	}{ 
		Country: o.Country,
		
		CountryName: o.CountryName,
		
		Latitude: o.Latitude,
		
		Longitude: o.Longitude,
		
		Locality: o.Locality,
		
		PostalCode: o.PostalCode,
		
		Region: o.Region,
		
		RegionName: o.RegionName,
		
		Timezone: o.Timezone,
		
		Source: o.Source,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationgeolocation) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationgeolocationMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationgeolocationMap)
	if err != nil {
		return err
	}
	
	if Country, ok := JourneyoutcomeeventsnotificationgeolocationMap["country"].(string); ok {
		o.Country = &Country
	}
	
	if CountryName, ok := JourneyoutcomeeventsnotificationgeolocationMap["countryName"].(string); ok {
		o.CountryName = &CountryName
	}
	
	if Latitude, ok := JourneyoutcomeeventsnotificationgeolocationMap["latitude"].(float64); ok {
		LatitudeFloat32 := float32(Latitude)
		o.Latitude = &LatitudeFloat32
	}
	
	if Longitude, ok := JourneyoutcomeeventsnotificationgeolocationMap["longitude"].(float64); ok {
		LongitudeFloat32 := float32(Longitude)
		o.Longitude = &LongitudeFloat32
	}
	
	if Locality, ok := JourneyoutcomeeventsnotificationgeolocationMap["locality"].(string); ok {
		o.Locality = &Locality
	}
	
	if PostalCode, ok := JourneyoutcomeeventsnotificationgeolocationMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
	
	if Region, ok := JourneyoutcomeeventsnotificationgeolocationMap["region"].(string); ok {
		o.Region = &Region
	}
	
	if RegionName, ok := JourneyoutcomeeventsnotificationgeolocationMap["regionName"].(string); ok {
		o.RegionName = &RegionName
	}
	
	if Timezone, ok := JourneyoutcomeeventsnotificationgeolocationMap["timezone"].(string); ok {
		o.Timezone = &Timezone
	}
	
	if Source, ok := JourneyoutcomeeventsnotificationgeolocationMap["source"].(string); ok {
		o.Source = &Source
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationgeolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
