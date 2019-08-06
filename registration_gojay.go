// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "github.com/francoispqt/gojay"

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Registration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyID, v.ID)
	enc.StringKey(keyMethod, v.Method)
	enc.AddInterfaceKeyOmitEmpty(keyRegisterOptions, &v.RegisterOptions)
}

// IsNil returns wether the structure is nil value or not.
func (v *Registration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Registration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyID:
		return dec.String(&v.ID)
	case keyMethod:
		return dec.String(&v.Method)
	case keyRegisterOptions:
		return dec.Interface(&v.RegisterOptions)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Registration) NKeys() int { return 3 }

// compile time check whether the Registration implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*Registration)(nil)
var _ gojay.UnmarshalerJSONObject = (*Registration)(nil)

type registrations []Registration

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *registrations) MarshalJSONArray(enc *gojay.Encoder) {
	vv := *v
	for i := range vv {
		enc.ObjectOmitEmpty(&vv[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *registrations) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *registrations) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := Registration{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// compile time check whether the registrations implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interface.
var _ gojay.MarshalerJSONArray = (*registrations)(nil)
var _ gojay.UnmarshalerJSONArray = (*registrations)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RegistrationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyRegistrations, (*registrations)(&v.Registrations))
}

// IsNil returns wether the structure is nil value or not.
func (v *RegistrationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RegistrationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyRegistrations {
		return dec.Array((*registrations)(&v.Registrations))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RegistrationParams) NKeys() int { return 1 }

// compile time check whether the RegistrationParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*RegistrationParams)(nil)
var _ gojay.UnmarshalerJSONObject = (*RegistrationParams)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDocumentSelector {
		return dec.Array(&v.DocumentSelector)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentRegistrationOptions) NKeys() int { return 1 }

// compile time check whether the TextDocumentRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*TextDocumentRegistrationOptions)(nil)
var _ gojay.UnmarshalerJSONObject = (*TextDocumentRegistrationOptions)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Unregistration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyID, v.ID)
	enc.StringKey(keyMethod, v.Method)
}

// IsNil returns wether the structure is nil value or not.
func (v *Unregistration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Unregistration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyID:
		return dec.String(&v.ID)
	case keyMethod:
		return dec.String(&v.Method)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Unregistration) NKeys() int { return 2 }

// compile time check whether the Unregistration implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*Unregistration)(nil)
var _ gojay.UnmarshalerJSONObject = (*Unregistration)(nil)

type unregisterations []Unregistration

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *unregisterations) MarshalJSONArray(enc *gojay.Encoder) {
	vv := *v
	for i := range vv {
		enc.ObjectOmitEmpty(&vv[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *unregisterations) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *unregisterations) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := Unregistration{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// compile time check whether the unregisterations implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interface.
var _ gojay.MarshalerJSONArray = (*unregisterations)(nil)
var _ gojay.UnmarshalerJSONArray = (*unregisterations)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *UnregistrationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyUnregisterations, (*unregisterations)(&v.Unregisterations))
}

// IsNil returns wether the structure is nil value or not.
func (v *UnregistrationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *UnregistrationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyUnregisterations {
		return dec.Array((*unregisterations)(&v.Unregisterations))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *UnregistrationParams) NKeys() int { return 1 }

// compile time check whether the UnregistrationParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*UnregistrationParams)(nil)
var _ gojay.UnmarshalerJSONObject = (*UnregistrationParams)(nil)
