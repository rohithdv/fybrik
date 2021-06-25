/*
 * Some Swagger
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AllowableActionDataset the model 'AllowableActionDataset'
type AllowableActionDataset string

// List of allowable_action_dataset
const (
	DENY_ACCESS AllowableActionDataset = "DenyAccess"
	DENY_WRITING AllowableActionDataset = "DenyWriting"
	ALLOW_ACCESS AllowableActionDataset = "AllowAccess"
)

var allowedAllowableActionDatasetEnumValues = []AllowableActionDataset{
	"DenyAccess",
	"DenyWriting",
	"AllowAccess",
}

func (v *AllowableActionDataset) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AllowableActionDataset(value)
	for _, existing := range allowedAllowableActionDatasetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AllowableActionDataset", value)
}

// NewAllowableActionDatasetFromValue returns a pointer to a valid AllowableActionDataset
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllowableActionDatasetFromValue(v string) (*AllowableActionDataset, error) {
	ev := AllowableActionDataset(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllowableActionDataset: valid values are %v", v, allowedAllowableActionDatasetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllowableActionDataset) IsValid() bool {
	for _, existing := range allowedAllowableActionDatasetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to allowable_action_dataset value
func (v AllowableActionDataset) Ptr() *AllowableActionDataset {
	return &v
}

type NullableAllowableActionDataset struct {
	value *AllowableActionDataset
	isSet bool
}

func (v NullableAllowableActionDataset) Get() *AllowableActionDataset {
	return v.value
}

func (v *NullableAllowableActionDataset) Set(val *AllowableActionDataset) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowableActionDataset) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowableActionDataset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowableActionDataset(val *AllowableActionDataset) *NullableAllowableActionDataset {
	return &NullableAllowableActionDataset{value: val, isSet: true}
}

func (v NullableAllowableActionDataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowableActionDataset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

