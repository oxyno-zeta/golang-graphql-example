//go:build unit

package utils

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalString(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name  string
		input any
		want  string
	}{
		{"plain text", "Fleet Paris – été", "Fleet Paris – été"},
		{"whitespace", " \n\t ", " \n\t "},
		{"script", "before<script>alert(1)</script>after", "beforeafter"},
		{"event handler", `<img src=x onerror=alert(1)>Fleet`, `<img src="x">Fleet`},
		{"javascript link", `<a href="javascript:alert(1)">Fleet</a>`, "Fleet"},
		{"formatting", "<b>Fleet</b>", "<b>Fleet</b>"},
		{"encoded markup", "&lt;script&gt;alert(1)&lt;/script&gt;", "&lt;script&gt;alert(1)&lt;/script&gt;"},
		{"ampersand", "R&D", "R&amp;D"},
		{"empty", "", ""},
		{"gqlgen coercion", 42, "42"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalString(tt.input)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}

	_, err := UnmarshalString(map[string]any{"text": "value"})
	require.Error(t, err)
}

func TestMarshalString(t *testing.T) {
	input := "A & B <example> \"quoted\"\n\\"
	var buf bytes.Buffer
	MarshalString(input).MarshalGQL(&buf)
	var got string
	require.NoError(t, json.Unmarshal(buf.Bytes(), &got))
	require.Equal(t, input, got)
}
