package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schemaquantitylimits
type Schemaquantitylimits struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// MinFieldNameCharacters - The minimum number of schema field name characters allowed.
	MinFieldNameCharacters *int `json:"minFieldNameCharacters,omitempty"`


	// MaxFieldNameCharacters - The maximum number of schema field name characters allowed.
	MaxFieldNameCharacters *int `json:"maxFieldNameCharacters,omitempty"`


	// MinFieldDescriptionCharacters - The minimum number of schema field description characters allowed.
	MinFieldDescriptionCharacters *int `json:"minFieldDescriptionCharacters,omitempty"`


	// MaxFieldDescriptionCharacters - The maximum number of schema field description characters allowed.
	MaxFieldDescriptionCharacters *int `json:"maxFieldDescriptionCharacters,omitempty"`


	// MinSchemaNameCharacters - The minimum number of schema name characters allowed.
	MinSchemaNameCharacters *int `json:"minSchemaNameCharacters,omitempty"`


	// MaxSchemaNameCharacters - The maximum number of schema name characters allowed.
	MaxSchemaNameCharacters *int `json:"maxSchemaNameCharacters,omitempty"`


	// MinSchemaDescriptionCharacters - The minimum number of schema description characters allowed.
	MinSchemaDescriptionCharacters *int `json:"minSchemaDescriptionCharacters,omitempty"`


	// MaxSchemaDescriptionCharacters - The maximum number of schema description characters allowed.
	MaxSchemaDescriptionCharacters *int `json:"maxSchemaDescriptionCharacters,omitempty"`


	// MaxNumberOfSchemasPerOrg - The maximum number of schema allowed per org.
	MaxNumberOfSchemasPerOrg *int `json:"maxNumberOfSchemasPerOrg,omitempty"`


	// MaxNumberOfFieldsPerSchema - The maximum number of schema fields allowed per schema.
	MaxNumberOfFieldsPerSchema *int `json:"maxNumberOfFieldsPerSchema,omitempty"`


	// MaxNumberOfFieldsPerOrg - The maximum number of schema fields allowed per organization across all of their schemas.
	MaxNumberOfFieldsPerOrg *int `json:"maxNumberOfFieldsPerOrg,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Schemaquantitylimits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schemaquantitylimits

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MinFieldNameCharacters *int `json:"minFieldNameCharacters,omitempty"`
		
		MaxFieldNameCharacters *int `json:"maxFieldNameCharacters,omitempty"`
		
		MinFieldDescriptionCharacters *int `json:"minFieldDescriptionCharacters,omitempty"`
		
		MaxFieldDescriptionCharacters *int `json:"maxFieldDescriptionCharacters,omitempty"`
		
		MinSchemaNameCharacters *int `json:"minSchemaNameCharacters,omitempty"`
		
		MaxSchemaNameCharacters *int `json:"maxSchemaNameCharacters,omitempty"`
		
		MinSchemaDescriptionCharacters *int `json:"minSchemaDescriptionCharacters,omitempty"`
		
		MaxSchemaDescriptionCharacters *int `json:"maxSchemaDescriptionCharacters,omitempty"`
		
		MaxNumberOfSchemasPerOrg *int `json:"maxNumberOfSchemasPerOrg,omitempty"`
		
		MaxNumberOfFieldsPerSchema *int `json:"maxNumberOfFieldsPerSchema,omitempty"`
		
		MaxNumberOfFieldsPerOrg *int `json:"maxNumberOfFieldsPerOrg,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		MinFieldNameCharacters: u.MinFieldNameCharacters,
		
		MaxFieldNameCharacters: u.MaxFieldNameCharacters,
		
		MinFieldDescriptionCharacters: u.MinFieldDescriptionCharacters,
		
		MaxFieldDescriptionCharacters: u.MaxFieldDescriptionCharacters,
		
		MinSchemaNameCharacters: u.MinSchemaNameCharacters,
		
		MaxSchemaNameCharacters: u.MaxSchemaNameCharacters,
		
		MinSchemaDescriptionCharacters: u.MinSchemaDescriptionCharacters,
		
		MaxSchemaDescriptionCharacters: u.MaxSchemaDescriptionCharacters,
		
		MaxNumberOfSchemasPerOrg: u.MaxNumberOfSchemasPerOrg,
		
		MaxNumberOfFieldsPerSchema: u.MaxNumberOfFieldsPerSchema,
		
		MaxNumberOfFieldsPerOrg: u.MaxNumberOfFieldsPerOrg,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Schemaquantitylimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
