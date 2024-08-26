// Code generated by github.com/AlexanderMint/gencodec2. DO NOT EDIT.

package ftypes

import (
	"encoding/json"
)

// MarshalJSON marshals as JSON.
func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		Int int
		Err error
		If  interface{}
	}
	var enc X
	enc.Int = x.Int
	enc.Err = x.Err
	enc.If = x.If
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		Int *int
		Err error
		If  interface{}
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Int != nil {
		x.Int = *dec.Int
	}
	if dec.Err != nil {
		x.Err = dec.Err
	}
	if dec.If != nil {
		x.If = dec.If
	}
	return nil
}
