package platformclientv2

import (
	"encoding/json"
	"github.com/leekchan/timeutil"
	"reflect"
	"strconv"
	"strings"
)

// Scimv2createuser - Defines the creation of a SCIM user.
type Scimv2createuser struct {
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`

	// Active - Indicates whether the user's administrative status is active.
	Active *bool `json:"active,omitempty"`

	// UserName - The user's Genesys Cloud email address. Must be unique.
	UserName *string `json:"userName,omitempty"`

	// DisplayName - The display name of the user.
	DisplayName *string `json:"displayName,omitempty"`

	// Password - The new password for the Genesys Cloud user. Does not return an existing password. When creating a user, if a password is not supplied, then a password will be randomly generated that is 40 characters in length and contains five characters from each of the password policy groups.
	Password *string `json:"password,omitempty"`

	// Title - The user's title.
	Title *string `json:"title,omitempty"`

	// PhoneNumbers - The list of the user's phone numbers.
	PhoneNumbers *[]Scimphonenumber `json:"phoneNumbers,omitempty"`

	// Emails - The list of the user's email addresses.
	Emails *[]Scimemail `json:"emails,omitempty"`

	// ExternalId - The external ID of the user. Set by the provisioning Client. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readWrite\".
	ExternalId *string `json:"externalId,omitempty"`

	// Groups - The list of groups that the user is a member of. This list is immutable per SCIM RFC and may only be updated using the GROUPS resource endpoint.
	Groups *[]Scimv2groupreference `json:"groups,omitempty"`

	// Roles - The list of roles assigned to the user.
	Roles *[]Scimuserrole `json:"roles,omitempty"`

	// UrnIetfParamsScimSchemasExtensionEnterprise20User - The URI of the schema for the enterprise user.
	UrnIetfParamsScimSchemasExtensionEnterprise20User *Scimv2enterpriseuser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`

	// UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User - The URI of the schema for the Genesys Cloud user.
	UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User *Scimuserextensions `json:"urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimv2createuser) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Scimv2createuser) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{}
		localDateTimeFields := []string{}
		dateFields := []string{}

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil() {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_ = timeutil.Timedelta{}
	type Alias Scimv2createuser

	return json.Marshal(&struct {
		Schemas *[]string `json:"schemas,omitempty"`

		Active *bool `json:"active,omitempty"`

		UserName *string `json:"userName,omitempty"`

		DisplayName *string `json:"displayName,omitempty"`

		Password *string `json:"password,omitempty"`

		Title *string `json:"title,omitempty"`

		PhoneNumbers *[]Scimphonenumber `json:"phoneNumbers,omitempty"`

		Emails *[]Scimemail `json:"emails,omitempty"`

		ExternalId *string `json:"externalId,omitempty"`

		Groups *[]Scimv2groupreference `json:"groups,omitempty"`

		Roles *[]Scimuserrole `json:"roles,omitempty"`

		UrnIetfParamsScimSchemasExtensionEnterprise20User *Scimv2enterpriseuser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`

		UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User *Scimuserextensions `json:"urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User,omitempty"`
		Alias
	}{
		Schemas: o.Schemas,

		Active: o.Active,

		UserName: o.UserName,

		DisplayName: o.DisplayName,

		Password: o.Password,

		Title: o.Title,

		PhoneNumbers: o.PhoneNumbers,

		Emails: o.Emails,

		ExternalId: o.ExternalId,

		Groups: o.Groups,

		Roles: o.Roles,

		UrnIetfParamsScimSchemasExtensionEnterprise20User: o.UrnIetfParamsScimSchemasExtensionEnterprise20User,

		UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User: o.UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User,
		Alias: (Alias)(o),
	})
}

func (o *Scimv2createuser) UnmarshalJSON(b []byte) error {
	var Scimv2createuserMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2createuserMap)
	if err != nil {
		return err
	}

	if Schemas, ok := Scimv2createuserMap["schemas"].([]interface{}); ok {
		SchemasString, _ := json.Marshal(Schemas)
		json.Unmarshal(SchemasString, &o.Schemas)
	}

	if Active, ok := Scimv2createuserMap["active"].(bool); ok {
		o.Active = &Active
	}

	if UserName, ok := Scimv2createuserMap["userName"].(string); ok {
		o.UserName = &UserName
	}

	if DisplayName, ok := Scimv2createuserMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}

	if Password, ok := Scimv2createuserMap["password"].(string); ok {
		o.Password = &Password
	}

	if Title, ok := Scimv2createuserMap["title"].(string); ok {
		o.Title = &Title
	}

	if PhoneNumbers, ok := Scimv2createuserMap["phoneNumbers"].([]interface{}); ok {
		PhoneNumbersString, _ := json.Marshal(PhoneNumbers)
		json.Unmarshal(PhoneNumbersString, &o.PhoneNumbers)
	}

	if Emails, ok := Scimv2createuserMap["emails"].([]interface{}); ok {
		EmailsString, _ := json.Marshal(Emails)
		json.Unmarshal(EmailsString, &o.Emails)
	}

	if ExternalId, ok := Scimv2createuserMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}

	if Groups, ok := Scimv2createuserMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}

	if Roles, ok := Scimv2createuserMap["roles"].([]interface{}); ok {
		RolesString, _ := json.Marshal(Roles)
		json.Unmarshal(RolesString, &o.Roles)
	}

	if UrnIetfParamsScimSchemasExtensionEnterprise20User, ok := Scimv2createuserMap["urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"].(map[string]interface{}); ok {
		UrnIetfParamsScimSchemasExtensionEnterprise20UserString, _ := json.Marshal(UrnIetfParamsScimSchemasExtensionEnterprise20User)
		json.Unmarshal(UrnIetfParamsScimSchemasExtensionEnterprise20UserString, &o.UrnIetfParamsScimSchemasExtensionEnterprise20User)
	}

	if UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User, ok := Scimv2createuserMap["urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User"].(map[string]interface{}); ok {
		UrnIetfParamsScimSchemasExtensionGenesysPurecloud20UserString, _ := json.Marshal(UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User)
		json.Unmarshal(UrnIetfParamsScimSchemasExtensionGenesysPurecloud20UserString, &o.UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User)
	}

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2createuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
