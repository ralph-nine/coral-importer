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

func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral(in *jlexer.Lexer, out *StorySettings) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "mode":
			if in.IsNull() {
				in.Skip()
				out.Mode = nil
			} else {
				if out.Mode == nil {
					out.Mode = new(string)
				}
				*out.Mode = string(in.String())
			}
		case "moderation":
			if in.IsNull() {
				in.Skip()
				out.Moderation = nil
			} else {
				if out.Moderation == nil {
					out.Moderation = new(string)
				}
				*out.Moderation = string(in.String())
			}
		case "messageBox":
			if in.IsNull() {
				in.Skip()
				out.MessageBox = nil
			} else {
				if out.MessageBox == nil {
					out.MessageBox = new(MessageBox)
				}
				(*out.MessageBox).UnmarshalEasyJSON(in)
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral(out *jwriter.Writer, in StorySettings) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Mode != nil {
		const prefix string = ",\"mode\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Mode))
	}
	if in.Moderation != nil {
		const prefix string = ",\"moderation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Moderation))
	}
	if in.MessageBox != nil {
		const prefix string = ",\"messageBox\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.MessageBox).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StorySettings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StorySettings) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StorySettings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StorySettings) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral(l, v)
}
func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral1(in *jlexer.Lexer, out *StoryMetadata) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "title":
			out.Title = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "image":
			out.Image = string(in.String())
		case "section":
			out.Section = string(in.String())
		case "publishedAt":
			if in.IsNull() {
				in.Skip()
				out.PublishedAt = nil
			} else {
				if out.PublishedAt == nil {
					out.PublishedAt = new(Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.PublishedAt).UnmarshalJSON(data))
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral1(out *jwriter.Writer, in StoryMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Title != "" {
		const prefix string = ",\"title\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	if in.Author != "" {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Image != "" {
		const prefix string = ",\"image\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Image))
	}
	if in.Section != "" {
		const prefix string = ",\"section\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Section))
	}
	if in.PublishedAt != nil {
		const prefix string = ",\"publishedAt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.PublishedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StoryMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StoryMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StoryMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StoryMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral1(l, v)
}
func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral2(in *jlexer.Lexer, out *StoryCommentCounts) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "action":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Action = make(map[string]int)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 int
					v1 = int(in.Int())
					(out.Action)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "status":
			(out.Status).UnmarshalEasyJSON(in)
		case "moderationQueue":
			(out.ModerationQueue).UnmarshalEasyJSON(in)
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral2(out *jwriter.Writer, in StoryCommentCounts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"action\":"
		out.RawString(prefix[1:])
		if in.Action == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Action {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				out.Int(int(v2Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		(in.Status).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"moderationQueue\":"
		out.RawString(prefix)
		(in.ModerationQueue).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StoryCommentCounts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StoryCommentCounts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StoryCommentCounts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StoryCommentCounts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral2(l, v)
}
func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral3(in *jlexer.Lexer, out *Story) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "_id":
			easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral4(in, &out.MongoID)
		case "tenantID":
			out.TenantID = string(in.StringIntern())
		case "id":
			out.ID = string(in.String())
		case "siteID":
			out.SiteID = string(in.StringIntern())
		case "url":
			out.URL = string(in.String())
		case "commentCounts":
			(out.CommentCounts).UnmarshalEasyJSON(in)
		case "settings":
			(out.Settings).UnmarshalEasyJSON(in)
		case "metadata":
			(out.Metadata).UnmarshalEasyJSON(in)
		case "closedAt":
			if in.IsNull() {
				in.Skip()
				out.ClosedAt = nil
			} else {
				if out.ClosedAt == nil {
					out.ClosedAt = new(Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ClosedAt).UnmarshalJSON(data))
				}
			}
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "importedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ImportedAt).UnmarshalJSON(data))
			}
		case "extra":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Extra = make(map[string]interface{})
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 interface{}
					if m, ok := v3.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v3.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v3 = in.Interface()
					}
					(out.Extra)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral3(out *jwriter.Writer, in Story) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_id\":"
		out.RawString(prefix[1:])
		easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral4(out, in.MongoID)
	}
	{
		const prefix string = ",\"tenantID\":"
		out.RawString(prefix)
		out.String(string(in.TenantID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"siteID\":"
		out.RawString(prefix)
		out.String(string(in.SiteID))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"commentCounts\":"
		out.RawString(prefix)
		(in.CommentCounts).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"settings\":"
		out.RawString(prefix)
		(in.Settings).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		(in.Metadata).MarshalEasyJSON(out)
	}
	if in.ClosedAt != nil {
		const prefix string = ",\"closedAt\":"
		out.RawString(prefix)
		out.Raw((*in.ClosedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"importedAt\":"
		out.RawString(prefix)
		out.Raw((in.ImportedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"extra\":"
		out.RawString(prefix)
		if in.Extra == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.Extra {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				if m, ok := v4Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v4Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v4Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Story) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Story) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Story) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Story) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral3(l, v)
}
func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral4(in *jlexer.Lexer, out *ObjectID) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "$oid":
			out.OID = string(in.String())
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral4(out *jwriter.Writer, in ObjectID) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"$oid\":"
		out.RawString(prefix[1:])
		out.String(string(in.OID))
	}
	out.RawByte('}')
}
func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral5(in *jlexer.Lexer, out *MessageBox) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "enabled":
			out.Enabled = bool(in.Bool())
		case "content":
			out.Content = string(in.String())
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral5(out *jwriter.Writer, in MessageBox) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"enabled\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Enabled))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessageBox) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessageBox) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessageBox) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessageBox) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral5(l, v)
}
func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral6(in *jlexer.Lexer, out *CommentStatusCounts) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "APPROVED":
			out.Approved = int(in.Int())
		case "NONE":
			out.None = int(in.Int())
		case "PREMOD":
			out.Premod = int(in.Int())
		case "REJECTED":
			out.Rejected = int(in.Int())
		case "SYSTEM_WITHHELD":
			out.SystemWithheld = int(in.Int())
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral6(out *jwriter.Writer, in CommentStatusCounts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"APPROVED\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Approved))
	}
	{
		const prefix string = ",\"NONE\":"
		out.RawString(prefix)
		out.Int(int(in.None))
	}
	{
		const prefix string = ",\"PREMOD\":"
		out.RawString(prefix)
		out.Int(int(in.Premod))
	}
	{
		const prefix string = ",\"REJECTED\":"
		out.RawString(prefix)
		out.Int(int(in.Rejected))
	}
	{
		const prefix string = ",\"SYSTEM_WITHHELD\":"
		out.RawString(prefix)
		out.Int(int(in.SystemWithheld))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentStatusCounts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentStatusCounts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentStatusCounts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentStatusCounts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral6(l, v)
}
func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral7(in *jlexer.Lexer, out *CommentModerationQueueCounts) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "total":
			out.Total = int(in.Int())
		case "queues":
			(out.Queues).UnmarshalEasyJSON(in)
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral7(out *jwriter.Writer, in CommentModerationQueueCounts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Total))
	}
	{
		const prefix string = ",\"queues\":"
		out.RawString(prefix)
		(in.Queues).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentModerationQueueCounts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentModerationQueueCounts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentModerationQueueCounts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentModerationQueueCounts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral7(l, v)
}
func easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral8(in *jlexer.Lexer, out *CommentModerationCountsPerQueue) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "unmoderated":
			out.Unmoderated = int(in.Int())
		case "pending":
			out.Pending = int(in.Int())
		case "reported":
			out.Reported = int(in.Int())
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
func easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral8(out *jwriter.Writer, in CommentModerationCountsPerQueue) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"unmoderated\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Unmoderated))
	}
	{
		const prefix string = ",\"pending\":"
		out.RawString(prefix)
		out.Int(int(in.Pending))
	}
	{
		const prefix string = ",\"reported\":"
		out.RawString(prefix)
		out.Int(int(in.Reported))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentModerationCountsPerQueue) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentModerationCountsPerQueue) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE08b5de9EncodeGithubComCoralprojectCoralImporterCommonCoral8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentModerationCountsPerQueue) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentModerationCountsPerQueue) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE08b5de9DecodeGithubComCoralprojectCoralImporterCommonCoral8(l, v)
}
