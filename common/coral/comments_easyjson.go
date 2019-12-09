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

func easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral(in *jlexer.Lexer, out *RevisionPerspective) {
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
		case "score":
			out.Score = float64(in.Float64())
		case "model":
			out.Model = string(in.String())
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
func easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral(out *jwriter.Writer, in RevisionPerspective) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"score\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Score))
	}
	{
		const prefix string = ",\"model\":"
		out.RawString(prefix)
		out.String(string(in.Model))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RevisionPerspective) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RevisionPerspective) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RevisionPerspective) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RevisionPerspective) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral(l, v)
}
func easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral1(in *jlexer.Lexer, out *RevisionMetadata) {
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
		case "akismet":
			if in.IsNull() {
				in.Skip()
				out.Akismet = nil
			} else {
				if out.Akismet == nil {
					out.Akismet = new(bool)
				}
				*out.Akismet = bool(in.Bool())
			}
		case "perspective":
			if in.IsNull() {
				in.Skip()
				out.Perspective = nil
			} else {
				if out.Perspective == nil {
					out.Perspective = new(RevisionPerspective)
				}
				(*out.Perspective).UnmarshalEasyJSON(in)
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
func easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral1(out *jwriter.Writer, in RevisionMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Akismet != nil {
		const prefix string = ",\"akismet\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(*in.Akismet))
	}
	if in.Perspective != nil {
		const prefix string = ",\"perspective\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Perspective).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RevisionMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RevisionMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RevisionMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RevisionMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral1(l, v)
}
func easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral2(in *jlexer.Lexer, out *Revision) {
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
		case "body":
			out.Body = HTML(in.String())
		case "actionCounts":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.ActionCounts = make(map[string]int)
				} else {
					out.ActionCounts = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 int
					v1 = int(in.Int())
					(out.ActionCounts)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "metadata":
			(out.Metadata).UnmarshalEasyJSON(in)
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
func easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral2(out *jwriter.Writer, in Revision) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.Raw((in.Body).MarshalJSON())
	}
	{
		const prefix string = ",\"actionCounts\":"
		out.RawString(prefix)
		if in.ActionCounts == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.ActionCounts {
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
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		(in.Metadata).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Revision) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Revision) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Revision) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Revision) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral2(l, v)
}
func easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral3(in *jlexer.Lexer, out *CommentTag) {
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
		case "type":
			out.Type = string(in.String())
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
func easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral3(out *jwriter.Writer, in CommentTag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentTag) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentTag) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentTag) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentTag) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral3(l, v)
}
func easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral4(in *jlexer.Lexer, out *Comment) {
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
		case "ancestorIDs":
			if in.IsNull() {
				in.Skip()
				out.AncestorIDs = nil
			} else {
				in.Delim('[')
				if out.AncestorIDs == nil {
					if !in.IsDelim(']') {
						out.AncestorIDs = make([]string, 0, 4)
					} else {
						out.AncestorIDs = []string{}
					}
				} else {
					out.AncestorIDs = (out.AncestorIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.AncestorIDs = append(out.AncestorIDs, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "parentID":
			out.ParentID = string(in.String())
		case "parentRevisionID":
			out.ParentRevisionID = string(in.String())
		case "authorID":
			out.AuthorID = string(in.String())
		case "storyID":
			out.StoryID = string(in.String())
		case "revisions":
			if in.IsNull() {
				in.Skip()
				out.Revisions = nil
			} else {
				in.Delim('[')
				if out.Revisions == nil {
					if !in.IsDelim(']') {
						out.Revisions = make([]Revision, 0, 1)
					} else {
						out.Revisions = []Revision{}
					}
				} else {
					out.Revisions = (out.Revisions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Revision
					(v4).UnmarshalEasyJSON(in)
					out.Revisions = append(out.Revisions, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "status":
			out.Status = string(in.String())
		case "actionCounts":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.ActionCounts = make(map[string]int)
				} else {
					out.ActionCounts = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 int
					v5 = int(in.Int())
					(out.ActionCounts)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
			}
		case "childIDs":
			if in.IsNull() {
				in.Skip()
				out.ChildIDs = nil
			} else {
				in.Delim('[')
				if out.ChildIDs == nil {
					if !in.IsDelim(']') {
						out.ChildIDs = make([]string, 0, 4)
					} else {
						out.ChildIDs = []string{}
					}
				} else {
					out.ChildIDs = (out.ChildIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.ChildIDs = append(out.ChildIDs, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]CommentTag, 0, 1)
					} else {
						out.Tags = []CommentTag{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v7 CommentTag
					(v7).UnmarshalEasyJSON(in)
					out.Tags = append(out.Tags, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "childCount":
			out.ChildCount = int(in.Int())
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "deletedAt":
			if in.IsNull() {
				in.Skip()
				out.DeletedAt = nil
			} else {
				if out.DeletedAt == nil {
					out.DeletedAt = new(Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DeletedAt).UnmarshalJSON(data))
				}
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
				if !in.IsDelim('}') {
					out.Extra = make(map[string]interface{})
				} else {
					out.Extra = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 interface{}
					if m, ok := v8.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v8.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v8 = in.Interface()
					}
					(out.Extra)[key] = v8
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
func easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral4(out *jwriter.Writer, in Comment) {
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
		const prefix string = ",\"ancestorIDs\":"
		out.RawString(prefix)
		if in.AncestorIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.AncestorIDs {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	if in.ParentID != "" {
		const prefix string = ",\"parentID\":"
		out.RawString(prefix)
		out.String(string(in.ParentID))
	}
	if in.ParentRevisionID != "" {
		const prefix string = ",\"parentRevisionID\":"
		out.RawString(prefix)
		out.String(string(in.ParentRevisionID))
	}
	{
		const prefix string = ",\"authorID\":"
		out.RawString(prefix)
		out.String(string(in.AuthorID))
	}
	{
		const prefix string = ",\"storyID\":"
		out.RawString(prefix)
		out.String(string(in.StoryID))
	}
	{
		const prefix string = ",\"revisions\":"
		out.RawString(prefix)
		if in.Revisions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Revisions {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"actionCounts\":"
		out.RawString(prefix)
		if in.ActionCounts == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v13First := true
			for v13Name, v13Value := range in.ActionCounts {
				if v13First {
					v13First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v13Name))
				out.RawByte(':')
				out.Int(int(v13Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"childIDs\":"
		out.RawString(prefix)
		if in.ChildIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.ChildIDs {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.Tags {
				if v16 > 0 {
					out.RawByte(',')
				}
				(v17).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"childCount\":"
		out.RawString(prefix)
		out.Int(int(in.ChildCount))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	if in.DeletedAt != nil {
		const prefix string = ",\"deletedAt\":"
		out.RawString(prefix)
		out.Raw((*in.DeletedAt).MarshalJSON())
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
			v18First := true
			for v18Name, v18Value := range in.Extra {
				if v18First {
					v18First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v18Name))
				out.RawByte(':')
				if m, ok := v18Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v18Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v18Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD09abad2EncodeGitlabComCoralprojectCoralImporterCommonCoral4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD09abad2DecodeGitlabComCoralprojectCoralImporterCommonCoral4(l, v)
}
