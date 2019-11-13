// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package coral

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

func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral(in *jlexer.Lexer, out *UserUsernameStatusHistory) {
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
		case "id":
			out.ID = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "createdBy":
			out.CreatedBy = string(in.String())
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral(out *jwriter.Writer, in UserUsernameStatusHistory) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"createdBy\":"
		out.RawString(prefix)
		out.String(string(in.CreatedBy))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserUsernameStatusHistory) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserUsernameStatusHistory) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserUsernameStatusHistory) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserUsernameStatusHistory) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral1(in *jlexer.Lexer, out *UserUsernameStatus) {
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
		case "history":
			if in.IsNull() {
				in.Skip()
				out.History = nil
			} else {
				in.Delim('[')
				if out.History == nil {
					if !in.IsDelim(']') {
						out.History = make([]UserUsernameStatusHistory, 0, 1)
					} else {
						out.History = []UserUsernameStatusHistory{}
					}
				} else {
					out.History = (out.History)[:0]
				}
				for !in.IsDelim(']') {
					var v1 UserUsernameStatusHistory
					(v1).UnmarshalEasyJSON(in)
					out.History = append(out.History, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral1(out *jwriter.Writer, in UserUsernameStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"history\":"
		out.RawString(prefix[1:])
		if in.History == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.History {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserUsernameStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserUsernameStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserUsernameStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserUsernameStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral1(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral2(in *jlexer.Lexer, out *UserToken) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral2(out *jwriter.Writer, in UserToken) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserToken) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserToken) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserToken) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserToken) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral2(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral3(in *jlexer.Lexer, out *UserSuspensionStatusHistory) {
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
		case "id":
			out.ID = string(in.String())
		case "from":
			(out.From).UnmarshalEasyJSON(in)
		case "createdBy":
			out.CreatedBy = string(in.String())
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "modifiedBy":
			if in.IsNull() {
				in.Skip()
				out.ModifiedBy = nil
			} else {
				if out.ModifiedBy == nil {
					out.ModifiedBy = new(string)
				}
				*out.ModifiedBy = string(in.String())
			}
		case "modifiedAt":
			if in.IsNull() {
				in.Skip()
				out.ModifiedAt = nil
			} else {
				if out.ModifiedAt == nil {
					out.ModifiedAt = new(Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ModifiedAt).UnmarshalJSON(data))
				}
			}
		case "message":
			out.Message = string(in.String())
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral3(out *jwriter.Writer, in UserSuspensionStatusHistory) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix)
		(in.From).MarshalEasyJSON(out)
	}
	if in.CreatedBy != "" {
		const prefix string = ",\"createdBy\":"
		out.RawString(prefix)
		out.String(string(in.CreatedBy))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	if in.ModifiedBy != nil {
		const prefix string = ",\"modifiedBy\":"
		out.RawString(prefix)
		out.String(string(*in.ModifiedBy))
	}
	if in.ModifiedAt != nil {
		const prefix string = ",\"modifiedAt\":"
		out.RawString(prefix)
		out.Raw((*in.ModifiedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserSuspensionStatusHistory) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserSuspensionStatusHistory) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserSuspensionStatusHistory) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserSuspensionStatusHistory) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral3(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral4(in *jlexer.Lexer, out *UserSuspensionStatus) {
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
		case "history":
			if in.IsNull() {
				in.Skip()
				out.History = nil
			} else {
				in.Delim('[')
				if out.History == nil {
					if !in.IsDelim(']') {
						out.History = make([]UserSuspensionStatusHistory, 0, 1)
					} else {
						out.History = []UserSuspensionStatusHistory{}
					}
				} else {
					out.History = (out.History)[:0]
				}
				for !in.IsDelim(']') {
					var v4 UserSuspensionStatusHistory
					(v4).UnmarshalEasyJSON(in)
					out.History = append(out.History, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral4(out *jwriter.Writer, in UserSuspensionStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"history\":"
		out.RawString(prefix[1:])
		if in.History == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.History {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserSuspensionStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserSuspensionStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserSuspensionStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserSuspensionStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral4(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral5(in *jlexer.Lexer, out *UserStatus) {
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
		case "suspension":
			(out.SuspensionStatus).UnmarshalEasyJSON(in)
		case "ban":
			(out.BanStatus).UnmarshalEasyJSON(in)
		case "username":
			(out.UsernameStatus).UnmarshalEasyJSON(in)
		case "premod":
			(out.PremodStatus).UnmarshalEasyJSON(in)
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral5(out *jwriter.Writer, in UserStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"suspension\":"
		out.RawString(prefix[1:])
		(in.SuspensionStatus).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"ban\":"
		out.RawString(prefix)
		(in.BanStatus).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		(in.UsernameStatus).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"premod\":"
		out.RawString(prefix)
		(in.PremodStatus).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral5(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral6(in *jlexer.Lexer, out *UserProfile) {
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
		case "id":
			out.ID = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "passwordID":
			out.PasswordID = string(in.String())
		case "lastIssuedAt":
			if in.IsNull() {
				in.Skip()
				out.LastIssuedAt = nil
			} else {
				if out.LastIssuedAt == nil {
					out.LastIssuedAt = new(Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastIssuedAt).UnmarshalJSON(data))
				}
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral6(out *jwriter.Writer, in UserProfile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	if in.PasswordID != "" {
		const prefix string = ",\"passwordID\":"
		out.RawString(prefix)
		out.String(string(in.PasswordID))
	}
	if in.LastIssuedAt != nil {
		const prefix string = ",\"lastIssuedAt\":"
		out.RawString(prefix)
		out.Raw((*in.LastIssuedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserProfile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserProfile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserProfile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserProfile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral6(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral7(in *jlexer.Lexer, out *UserPremodStatus) {
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
		case "active":
			out.Active = bool(in.Bool())
		case "history":
			if in.IsNull() {
				in.Skip()
				out.History = nil
			} else {
				in.Delim('[')
				if out.History == nil {
					if !in.IsDelim(']') {
						out.History = make([]string, 0, 4)
					} else {
						out.History = []string{}
					}
				} else {
					out.History = (out.History)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.History = append(out.History, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral7(out *jwriter.Writer, in UserPremodStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"active\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Active))
	}
	{
		const prefix string = ",\"history\":"
		out.RawString(prefix)
		if in.History == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.History {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserPremodStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserPremodStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserPremodStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserPremodStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral7(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral8(in *jlexer.Lexer, out *UserNotifications) {
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
		case "onReply":
			out.OnReply = bool(in.Bool())
		case "onFeatured":
			out.OnFeatured = bool(in.Bool())
		case "onStaffReplies":
			out.OnStaffReplies = bool(in.Bool())
		case "onModeration":
			out.OnModeration = bool(in.Bool())
		case "digestFrequency":
			out.DigestFrequency = string(in.String())
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral8(out *jwriter.Writer, in UserNotifications) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"onReply\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.OnReply))
	}
	{
		const prefix string = ",\"onFeatured\":"
		out.RawString(prefix)
		out.Bool(bool(in.OnFeatured))
	}
	{
		const prefix string = ",\"onStaffReplies\":"
		out.RawString(prefix)
		out.Bool(bool(in.OnStaffReplies))
	}
	{
		const prefix string = ",\"onModeration\":"
		out.RawString(prefix)
		out.Bool(bool(in.OnModeration))
	}
	{
		const prefix string = ",\"digestFrequency\":"
		out.RawString(prefix)
		out.String(string(in.DigestFrequency))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserNotifications) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserNotifications) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserNotifications) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserNotifications) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral8(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral9(in *jlexer.Lexer, out *UserBanStatusHistory) {
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
		case "id":
			out.ID = string(in.String())
		case "active":
			out.Active = bool(in.Bool())
		case "createdBy":
			out.CreatedBy = string(in.String())
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "message":
			out.Message = string(in.String())
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral9(out *jwriter.Writer, in UserBanStatusHistory) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"active\":"
		out.RawString(prefix)
		out.Bool(bool(in.Active))
	}
	if in.CreatedBy != "" {
		const prefix string = ",\"createdBy\":"
		out.RawString(prefix)
		out.String(string(in.CreatedBy))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserBanStatusHistory) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserBanStatusHistory) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserBanStatusHistory) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserBanStatusHistory) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral9(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral10(in *jlexer.Lexer, out *UserBanStatus) {
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
		case "active":
			out.Active = bool(in.Bool())
		case "history":
			if in.IsNull() {
				in.Skip()
				out.History = nil
			} else {
				in.Delim('[')
				if out.History == nil {
					if !in.IsDelim(']') {
						out.History = make([]UserBanStatusHistory, 0, 1)
					} else {
						out.History = []UserBanStatusHistory{}
					}
				} else {
					out.History = (out.History)[:0]
				}
				for !in.IsDelim(']') {
					var v10 UserBanStatusHistory
					(v10).UnmarshalEasyJSON(in)
					out.History = append(out.History, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral10(out *jwriter.Writer, in UserBanStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"active\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Active))
	}
	{
		const prefix string = ",\"history\":"
		out.RawString(prefix)
		if in.History == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.History {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserBanStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserBanStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserBanStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserBanStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral10(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral11(in *jlexer.Lexer, out *User) {
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
		case "tenantID":
			out.TenantID = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "profiles":
			if in.IsNull() {
				in.Skip()
				out.Profiles = nil
			} else {
				in.Delim('[')
				if out.Profiles == nil {
					if !in.IsDelim(']') {
						out.Profiles = make([]UserProfile, 0, 1)
					} else {
						out.Profiles = []UserProfile{}
					}
				} else {
					out.Profiles = (out.Profiles)[:0]
				}
				for !in.IsDelim(']') {
					var v13 UserProfile
					(v13).UnmarshalEasyJSON(in)
					out.Profiles = append(out.Profiles, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "role":
			out.Role = string(in.String())
		case "notifications":
			(out.Notifications).UnmarshalEasyJSON(in)
		case "status":
			(out.Status).UnmarshalEasyJSON(in)
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "ignoredUsers":
			if in.IsNull() {
				in.Skip()
				out.IgnoredUsers = nil
			} else {
				in.Delim('[')
				if out.IgnoredUsers == nil {
					if !in.IsDelim(']') {
						out.IgnoredUsers = make([]IgnoredUser, 0, 1)
					} else {
						out.IgnoredUsers = []IgnoredUser{}
					}
				} else {
					out.IgnoredUsers = (out.IgnoredUsers)[:0]
				}
				for !in.IsDelim(']') {
					var v14 IgnoredUser
					(v14).UnmarshalEasyJSON(in)
					out.IgnoredUsers = append(out.IgnoredUsers, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tokens":
			if in.IsNull() {
				in.Skip()
				out.Tokens = nil
			} else {
				in.Delim('[')
				if out.Tokens == nil {
					if !in.IsDelim(']') {
						out.Tokens = make([]UserToken, 0, 1)
					} else {
						out.Tokens = []UserToken{}
					}
				} else {
					out.Tokens = (out.Tokens)[:0]
				}
				for !in.IsDelim(']') {
					var v15 UserToken
					(v15).UnmarshalEasyJSON(in)
					out.Tokens = append(out.Tokens, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "lastDownloadedAt":
			if in.IsNull() {
				in.Skip()
				out.LastDownloadedAt = nil
			} else {
				if out.LastDownloadedAt == nil {
					out.LastDownloadedAt = new(Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastDownloadedAt).UnmarshalJSON(data))
				}
			}
		case "imported":
			out.Imported = bool(in.Bool())
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral11(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tenantID\":"
		out.RawString(prefix[1:])
		out.String(string(in.TenantID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	if len(in.Profiles) != 0 {
		const prefix string = ",\"profiles\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v16, v17 := range in.Profiles {
				if v16 > 0 {
					out.RawByte(',')
				}
				(v17).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		out.String(string(in.Role))
	}
	{
		const prefix string = ",\"notifications\":"
		out.RawString(prefix)
		(in.Notifications).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		(in.Status).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"ignoredUsers\":"
		out.RawString(prefix)
		if in.IgnoredUsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.IgnoredUsers {
				if v18 > 0 {
					out.RawByte(',')
				}
				(v19).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tokens\":"
		out.RawString(prefix)
		if in.Tokens == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Tokens {
				if v20 > 0 {
					out.RawByte(',')
				}
				(v21).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"lastDownloadedAt\":"
		out.RawString(prefix)
		if in.LastDownloadedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.LastDownloadedAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"imported\":"
		out.RawString(prefix)
		out.Bool(bool(in.Imported))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral11(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral12(in *jlexer.Lexer, out *TimeRange) {
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
		case "from":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.From).UnmarshalJSON(data))
			}
		case "to":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.To).UnmarshalJSON(data))
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral12(out *jwriter.Writer, in TimeRange) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix[1:])
		out.Raw((in.From).MarshalJSON())
	}
	{
		const prefix string = ",\"to\":"
		out.RawString(prefix)
		out.Raw((in.To).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TimeRange) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TimeRange) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TimeRange) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TimeRange) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral12(l, v)
}
func easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral13(in *jlexer.Lexer, out *IgnoredUser) {
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
		case "id":
			out.ID = string(in.String())
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
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
func easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral13(out *jwriter.Writer, in IgnoredUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IgnoredUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IgnoredUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84c0690eEncodeGitlabComCoralprojectCoralImporterCommonCoral13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IgnoredUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IgnoredUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84c0690eDecodeGitlabComCoralprojectCoralImporterCommonCoral13(l, v)
}
