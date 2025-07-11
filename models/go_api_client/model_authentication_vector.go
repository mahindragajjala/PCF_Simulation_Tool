/*
NudmUEAU

UDM UE Authentication Service. � 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// AuthenticationVector - struct for AuthenticationVector
type AuthenticationVector struct {
	Av5GHeAka *Av5GHeAka
	AvEapAkaPrime *AvEapAkaPrime
}

// Av5GHeAkaAsAuthenticationVector is a convenience function that returns Av5GHeAka wrapped in AuthenticationVector
func Av5GHeAkaAsAuthenticationVector(v *Av5GHeAka) AuthenticationVector {
	return AuthenticationVector{
		Av5GHeAka: v,
	}
}

// AvEapAkaPrimeAsAuthenticationVector is a convenience function that returns AvEapAkaPrime wrapped in AuthenticationVector
func AvEapAkaPrimeAsAuthenticationVector(v *AvEapAkaPrime) AuthenticationVector {
	return AuthenticationVector{
		AvEapAkaPrime: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AuthenticationVector) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Av5GHeAka
	err = newStrictDecoder(data).Decode(&dst.Av5GHeAka)
	if err == nil {
		jsonAv5GHeAka, _ := json.Marshal(dst.Av5GHeAka)
		if string(jsonAv5GHeAka) == "{}" { // empty struct
			dst.Av5GHeAka = nil
		} else {
			if err = validator.Validate(dst.Av5GHeAka); err != nil {
				dst.Av5GHeAka = nil
			} else {
				match++
			}
		}
	} else {
		dst.Av5GHeAka = nil
	}

	// try to unmarshal data into AvEapAkaPrime
	err = newStrictDecoder(data).Decode(&dst.AvEapAkaPrime)
	if err == nil {
		jsonAvEapAkaPrime, _ := json.Marshal(dst.AvEapAkaPrime)
		if string(jsonAvEapAkaPrime) == "{}" { // empty struct
			dst.AvEapAkaPrime = nil
		} else {
			if err = validator.Validate(dst.AvEapAkaPrime); err != nil {
				dst.AvEapAkaPrime = nil
			} else {
				match++
			}
		}
	} else {
		dst.AvEapAkaPrime = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Av5GHeAka = nil
		dst.AvEapAkaPrime = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AuthenticationVector)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AuthenticationVector)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AuthenticationVector) MarshalJSON() ([]byte, error) {
	if src.Av5GHeAka != nil {
		return json.Marshal(&src.Av5GHeAka)
	}

	if src.AvEapAkaPrime != nil {
		return json.Marshal(&src.AvEapAkaPrime)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AuthenticationVector) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Av5GHeAka != nil {
		return obj.Av5GHeAka
	}

	if obj.AvEapAkaPrime != nil {
		return obj.AvEapAkaPrime
	}

	// all schemas are nil
	return nil
}

type NullableAuthenticationVector struct {
	value *AuthenticationVector
	isSet bool
}

func (v NullableAuthenticationVector) Get() *AuthenticationVector {
	return v.value
}

func (v *NullableAuthenticationVector) Set(val *AuthenticationVector) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationVector) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationVector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationVector(val *AuthenticationVector) *NullableAuthenticationVector {
	return &NullableAuthenticationVector{value: val, isSet: true}
}

func (v NullableAuthenticationVector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationVector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


