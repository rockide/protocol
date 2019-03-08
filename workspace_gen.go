// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceFolder) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "name":
		return dec.String(&v.Name)
	case "uRI":
		return dec.String(&v.URI)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceFolder) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceFolder) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("name", v.Name)
	enc.StringKey("uRI", v.URI)
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceFolder) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DidChangeWorkspaceFoldersParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "event":
		if v.Event == nil {
			v.Event = &WorkspaceFoldersChangeEvent{}
		}
		return dec.Object(v.Event)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DidChangeWorkspaceFoldersParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DidChangeWorkspaceFoldersParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("event", v.Event)
}

// IsNil returns wether the structure is nil value or not
func (v *DidChangeWorkspaceFoldersParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceFoldersChangeEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceFoldersChangeEvent) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceFoldersChangeEvent) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceFoldersChangeEvent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DidChangeConfigurationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DidChangeConfigurationParams) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DidChangeConfigurationParams) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *DidChangeConfigurationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ConfigurationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ConfigurationParams) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ConfigurationParams) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *ConfigurationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ConfigurationItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "scopeURI":
		return dec.String(&v.ScopeURI)
	case "section":
		return dec.String(&v.Section)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ConfigurationItem) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ConfigurationItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("scopeURI", v.ScopeURI)
	enc.StringKey("section", v.Section)
}

// IsNil returns wether the structure is nil value or not
func (v *ConfigurationItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DidChangeWatchedFilesParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DidChangeWatchedFilesParams) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DidChangeWatchedFilesParams) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *DidChangeWatchedFilesParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *FileEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "type":
		return dec.Float64(&v.Type)
	case "uri":
		return dec.String((*string)(&v.URI))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *FileEvent) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *FileEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key("type", v.Type)
	enc.StringKey("uri", string(v.URI))
}

// IsNil returns wether the structure is nil value or not
func (v *FileEvent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DidChangeWatchedFilesRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DidChangeWatchedFilesRegistrationOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DidChangeWatchedFilesRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *DidChangeWatchedFilesRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *FileSystemWatcher) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "globPattern":
		return dec.String(&v.GlobPattern)
	case "kind":
		return dec.Float64((*float64)(&v.Kind))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *FileSystemWatcher) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *FileSystemWatcher) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("globPattern", v.GlobPattern)
	enc.Float64Key("kind", float64(v.Kind))
}

// IsNil returns wether the structure is nil value or not
func (v *FileSystemWatcher) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceSymbolParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "query":
		return dec.String(&v.Query)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceSymbolParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceSymbolParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("query", v.Query)
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceSymbolParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ExecuteCommandParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "command":
		return dec.String(&v.Command)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ExecuteCommandParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ExecuteCommandParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("command", v.Command)
}

// IsNil returns wether the structure is nil value or not
func (v *ExecuteCommandParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ExecuteCommandRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ExecuteCommandRegistrationOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ExecuteCommandRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *ExecuteCommandRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ApplyWorkspaceEditParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "label":
		return dec.String(&v.Label)
	case "edit":
		return dec.Object(&v.Edit)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ApplyWorkspaceEditParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ApplyWorkspaceEditParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("label", v.Label)
	enc.ObjectKey("edit", &v.Edit)
}

// IsNil returns wether the structure is nil value or not
func (v *ApplyWorkspaceEditParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ApplyWorkspaceEditResponse) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "applied":
		return dec.Bool(&v.Applied)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ApplyWorkspaceEditResponse) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ApplyWorkspaceEditResponse) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("applied", v.Applied)
}

// IsNil returns wether the structure is nil value or not
func (v *ApplyWorkspaceEditResponse) IsNil() bool { return v == nil }
