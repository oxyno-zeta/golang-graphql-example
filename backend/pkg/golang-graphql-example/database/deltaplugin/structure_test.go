package deltaplugin

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestNanoDateTime_MarshalJSON(t *testing.T) {
	forgedDate, err := time.Parse(time.RFC3339Nano, "2009-11-10T23:00:00.999999Z")
	if err != nil {
		t.Error(err)
	}
	forgedDate2, err := time.Parse(time.RFC3339Nano, "2009-11-10T23:00:00Z")
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name    string
		d       NanoDateTime
		want    string
		wantErr bool
	}{
		{
			name: "should be ok with nanoseconds",
			d:    NanoDateTime(forgedDate),
			want: "\"2009-11-10T23:00:00.999999Z\"",
		},
		{
			name: "should be ok without nanoseconds",
			d:    NanoDateTime(forgedDate2),
			want: "\"2009-11-10T23:00:00Z\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("NanoDateTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("NanoDateTime.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestNanoDateTime_UnmarshalJSON(t *testing.T) {
	forgedDate, err := time.Parse(time.RFC3339Nano, "2009-11-10T23:00:00.999999Z")
	if err != nil {
		t.Error(err)
	}
	forgedDate2, err := time.Parse(time.RFC3339Nano, "2009-11-10T23:00:00Z")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    NanoDateTime
		wantErr bool
	}{
		{
			name: "empty case",
			args: args{data: []byte("\"\"")},
			want: NanoDateTime{},
		},
		{
			name: "null case",
			args: args{data: []byte("null")},
			want: NanoDateTime{},
		},
		{
			name: "should be ok with nanoseconds",
			args: args{data: []byte("\"2009-11-10T23:00:00.999999Z\"")},
			want: NanoDateTime(forgedDate),
		},
		{
			name: "should be ok without nanoseconds",
			args: args{data: []byte("\"2009-11-10T23:00:00Z\"")},
			want: NanoDateTime(forgedDate2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NanoDateTime{}
			err := json.Unmarshal(tt.args.data, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("NanoDateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NanoDateTime.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
