// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package easyJson

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC80ae7adDecodeAppEasyJson(in *jlexer.Lexer, out *ResponsePattern) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "location":
			out.Location = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeAppEasyJson(out *jwriter.Writer, in ResponsePattern) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"location\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Location))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponsePattern) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeAppEasyJson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponsePattern) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeAppEasyJson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponsePattern) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeAppEasyJson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponsePattern) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeAppEasyJson(l, v)
}
func easyjsonC80ae7adDecodeAppEasyJson1(in *jlexer.Lexer, out *RequestPattern) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "domain":
			out.Domain = string(in.String())
		case "path":
			out.Path = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeAppEasyJson1(out *jwriter.Writer, in RequestPattern) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"domain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Domain))
	}
	{
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Path))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestPattern) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeAppEasyJson1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestPattern) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeAppEasyJson1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestPattern) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeAppEasyJson1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestPattern) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeAppEasyJson1(l, v)
}