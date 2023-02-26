//go:build unit

package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateFilter_GetGenericFilter(t *testing.T) {
	notADate := "not a date"
	dateStr := "2020-09-19T23:10:35+02:00"
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		t.Error(err)
		return
	}
	dateStr2 := "2020-09-19T23:10:35.999+02:00"
	date2, err := time.Parse(time.RFC3339Nano, dateStr2)
	if err != nil {
		t.Error(err)
		return
	}

	date = date.UTC()
	date2 = date2.UTC()

	type fields struct {
		Eq        interface{}
		NotEq     interface{}
		Gte       interface{}
		NotGte    interface{}
		Gt        interface{}
		NotGt     interface{}
		Lte       interface{}
		NotLte    interface{}
		Lt        interface{}
		NotLt     interface{}
		In        interface{}
		NotIn     interface{}
		IsNull    bool
		IsNotNull bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GenericFilter
		wantErr bool
	}{
		{
			name:   "Eq case (pointer string)",
			fields: fields{Eq: &dateStr},
			want:   &GenericFilter{Eq: &date},
		},
		{
			name:   "Eq case (string)",
			fields: fields{Eq: dateStr},
			want:   &GenericFilter{Eq: &date},
		},
		{
			name:   "Eq case (pointer date)",
			fields: fields{Eq: &date},
			want:   &GenericFilter{Eq: &date},
		},
		{
			name:   "Eq case (date)",
			fields: fields{Eq: date},
			want:   &GenericFilter{Eq: &date},
		},
		{
			name:   "Eq case (nano case pointer string)",
			fields: fields{Eq: &dateStr2},
			want:   &GenericFilter{Eq: &date2},
		},
		{
			name:    "Eq not a date",
			fields:  fields{Eq: &notADate},
			wantErr: true,
		},
		{
			name:    "Eq not a date 2",
			fields:  fields{Eq: notADate},
			wantErr: true,
		},
		{
			name:    "Eq not a date 3",
			fields:  fields{Eq: 9},
			wantErr: true,
		},
		{
			name:   "Eq nil",
			fields: fields{Eq: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "NotEq case (pointer string)",
			fields: fields{NotEq: &dateStr},
			want:   &GenericFilter{NotEq: &date},
		},
		{
			name:   "NotEq case (string)",
			fields: fields{NotEq: dateStr},
			want:   &GenericFilter{NotEq: &date},
		},
		{
			name:   "NotEq case (pointer date)",
			fields: fields{NotEq: &date},
			want:   &GenericFilter{NotEq: &date},
		},
		{
			name:   "NotEq case (date)",
			fields: fields{NotEq: date},
			want:   &GenericFilter{NotEq: &date},
		},
		{
			name:   "NotEq case (nano case pointer string)",
			fields: fields{NotEq: &dateStr2},
			want:   &GenericFilter{NotEq: &date2},
		},
		{
			name:    "NotEq not a date",
			fields:  fields{NotEq: &notADate},
			wantErr: true,
		},
		{
			name:    "NotEq not a date 2",
			fields:  fields{NotEq: notADate},
			wantErr: true,
		},
		{
			name:    "NotEq not a date 3",
			fields:  fields{NotEq: 9},
			wantErr: true,
		},
		{
			name:   "NotEq nil",
			fields: fields{NotEq: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "Gte case (pointer string)",
			fields: fields{Gte: &dateStr},
			want:   &GenericFilter{Gte: &date},
		},
		{
			name:   "Gte case (string)",
			fields: fields{Gte: dateStr},
			want:   &GenericFilter{Gte: &date},
		},
		{
			name:   "Gte case (pointer date)",
			fields: fields{Gte: &date},
			want:   &GenericFilter{Gte: &date},
		},
		{
			name:   "Gte case (date)",
			fields: fields{Gte: date},
			want:   &GenericFilter{Gte: &date},
		},
		{
			name:   "Gte case (nano date pointer string)",
			fields: fields{Gte: &dateStr2},
			want:   &GenericFilter{Gte: &date2},
		},
		{
			name:    "Gte not a date",
			fields:  fields{Gte: &notADate},
			wantErr: true,
		},
		{
			name:    "Gte not a date 2",
			fields:  fields{Gte: notADate},
			wantErr: true,
		},
		{
			name:    "Gte not a date 3",
			fields:  fields{Gte: 9},
			wantErr: true,
		},
		{
			name:   "Gte nil",
			fields: fields{Gte: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "NotGte case (pointer string)",
			fields: fields{NotGte: &dateStr},
			want:   &GenericFilter{NotGte: &date},
		},
		{
			name:   "NotGte case (string)",
			fields: fields{NotGte: dateStr},
			want:   &GenericFilter{NotGte: &date},
		},
		{
			name:   "NotGte case (pointer date)",
			fields: fields{NotGte: &date},
			want:   &GenericFilter{NotGte: &date},
		},
		{
			name:   "NotGte case (date)",
			fields: fields{NotGte: date},
			want:   &GenericFilter{NotGte: &date},
		},
		{
			name:   "NotGte case (nano case pointer string)",
			fields: fields{NotGte: &dateStr2},
			want:   &GenericFilter{NotGte: &date2},
		},
		{
			name:    "NotGte not a date",
			fields:  fields{NotGte: &notADate},
			wantErr: true,
		},
		{
			name:    "NotGte not a date 2",
			fields:  fields{NotGte: notADate},
			wantErr: true,
		},
		{
			name:   "NotGte nil",
			fields: fields{NotGte: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "Gt case (pointer string)",
			fields: fields{Gt: &dateStr},
			want:   &GenericFilter{Gt: &date},
		},
		{
			name:   "Gt case (string)",
			fields: fields{Gt: dateStr},
			want:   &GenericFilter{Gt: &date},
		},
		{
			name:   "Gt case (pointer date)",
			fields: fields{Gt: &date},
			want:   &GenericFilter{Gt: &date},
		},
		{
			name:   "Gt case (date)",
			fields: fields{Gt: date},
			want:   &GenericFilter{Gt: &date},
		},
		{
			name:   "Gt case (nano case pointer string)",
			fields: fields{Gt: &dateStr2},
			want:   &GenericFilter{Gt: &date2},
		},
		{
			name:    "Gt not a date",
			fields:  fields{Gt: &notADate},
			wantErr: true,
		},
		{
			name:    "Gt not a date 2",
			fields:  fields{Gt: notADate},
			wantErr: true,
		},
		{
			name:    "Gt not a date 3",
			fields:  fields{Gt: 9},
			wantErr: true,
		},
		{
			name:   "Gt nil",
			fields: fields{Gt: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "NotGt case (pointer string)",
			fields: fields{NotGt: &dateStr},
			want:   &GenericFilter{NotGt: &date},
		},
		{
			name:   "NotGt case (string)",
			fields: fields{NotGt: dateStr},
			want:   &GenericFilter{NotGt: &date},
		},
		{
			name:   "NotGt case (pointer date)",
			fields: fields{NotGt: &date},
			want:   &GenericFilter{NotGt: &date},
		},
		{
			name:   "NotGt case (date)",
			fields: fields{NotGt: date},
			want:   &GenericFilter{NotGt: &date},
		},
		{
			name:   "NotGt case (nano case pointer string)",
			fields: fields{NotGt: &dateStr2},
			want:   &GenericFilter{NotGt: &date2},
		},
		{
			name:    "NotGt not a date",
			fields:  fields{NotGt: &notADate},
			wantErr: true,
		},
		{
			name:    "NotGt not a date 2",
			fields:  fields{NotGt: notADate},
			wantErr: true,
		},
		{
			name:    "NotGt not a date 3",
			fields:  fields{NotGt: 9},
			wantErr: true,
		},
		{
			name:   "NotGt nil",
			fields: fields{NotGt: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "Lte case (pointer string)",
			fields: fields{Lte: &dateStr},
			want:   &GenericFilter{Lte: &date},
		},
		{
			name:   "Lte case (string)",
			fields: fields{Lte: dateStr},
			want:   &GenericFilter{Lte: &date},
		},
		{
			name:   "Lte case (pointer date)",
			fields: fields{Lte: &date},
			want:   &GenericFilter{Lte: &date},
		},
		{
			name:   "Lte case (date)",
			fields: fields{Lte: date},
			want:   &GenericFilter{Lte: &date},
		},
		{
			name:   "Lte case (nano case pointer string)",
			fields: fields{Lte: &dateStr2},
			want:   &GenericFilter{Lte: &date2},
		},
		{
			name:    "Lte not a date",
			fields:  fields{Lte: &notADate},
			wantErr: true,
		},
		{
			name:    "Lte not a date 2",
			fields:  fields{Lte: notADate},
			wantErr: true,
		},
		{
			name:    "Lte not a date 3",
			fields:  fields{Lte: 9},
			wantErr: true,
		},
		{
			name:   "Lte nil",
			fields: fields{Lte: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "NotLte case (pointer string)",
			fields: fields{NotLte: &dateStr},
			want:   &GenericFilter{NotLte: &date},
		},
		{
			name:   "NotLte case (string)",
			fields: fields{NotLte: dateStr},
			want:   &GenericFilter{NotLte: &date},
		},
		{
			name:   "NotLte case (pointer date)",
			fields: fields{NotLte: &date},
			want:   &GenericFilter{NotLte: &date},
		},
		{
			name:   "NotLte case (date)",
			fields: fields{NotLte: date},
			want:   &GenericFilter{NotLte: &date},
		},
		{
			name:   "NotLte case (nano case pointer string)",
			fields: fields{NotLte: &dateStr2},
			want:   &GenericFilter{NotLte: &date2},
		},
		{
			name:    "NotLte not a date",
			fields:  fields{NotLte: &notADate},
			wantErr: true,
		},
		{
			name:    "NotLte not a date 2",
			fields:  fields{NotLte: notADate},
			wantErr: true,
		},
		{
			name:    "NotLte not a date 3",
			fields:  fields{NotLte: 9},
			wantErr: true,
		},
		{
			name:   "NotLte nil",
			fields: fields{NotLte: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "Lt case (pointer string)",
			fields: fields{Lt: &dateStr},
			want:   &GenericFilter{Lt: &date},
		},
		{
			name:   "Lt case (string)",
			fields: fields{Lt: dateStr},
			want:   &GenericFilter{Lt: &date},
		},
		{
			name:   "Lt case (pointer date)",
			fields: fields{Lt: &date},
			want:   &GenericFilter{Lt: &date},
		},
		{
			name:   "Lt case (date)",
			fields: fields{Lt: date},
			want:   &GenericFilter{Lt: &date},
		},
		{
			name:   "Lt case (nano case pointer string)",
			fields: fields{Lt: &dateStr2},
			want:   &GenericFilter{Lt: &date2},
		},
		{
			name:    "Lt not a date",
			fields:  fields{Lt: &notADate},
			wantErr: true,
		},
		{
			name:    "Lt not a date 2",
			fields:  fields{Lt: notADate},
			wantErr: true,
		},
		{
			name:    "Lt not a date 3",
			fields:  fields{Lt: 9},
			wantErr: true,
		},
		{
			name:   "Lt nil",
			fields: fields{Lt: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "NotLt case (pointer string)",
			fields: fields{NotLt: &dateStr},
			want:   &GenericFilter{NotLt: &date},
		},
		{
			name:   "NotLt case (string)",
			fields: fields{NotLt: dateStr},
			want:   &GenericFilter{NotLt: &date},
		},
		{
			name:   "NotLt case (pointer date)",
			fields: fields{NotLt: &date},
			want:   &GenericFilter{NotLt: &date},
		},
		{
			name:   "NotLt case (date)",
			fields: fields{NotLt: date},
			want:   &GenericFilter{NotLt: &date},
		},
		{
			name:   "NotLt case (nano case pointer string)",
			fields: fields{NotLt: &dateStr2},
			want:   &GenericFilter{NotLt: &date2},
		},
		{
			name:    "NotLt not a date",
			fields:  fields{NotLt: &notADate},
			wantErr: true,
		},
		{
			name:    "NotLt not a date 2",
			fields:  fields{NotLt: notADate},
			wantErr: true,
		},
		{
			name:    "NotLt not a date 3",
			fields:  fields{NotLt: 9},
			wantErr: true,
		},
		{
			name:   "NotLt nil",
			fields: fields{NotLt: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "In case (pointer string)",
			fields: fields{In: []*string{&dateStr}},
			want:   &GenericFilter{In: []*time.Time{&date}},
		},
		{
			name:   "In case (string)",
			fields: fields{In: []string{dateStr}},
			want:   &GenericFilter{In: []*time.Time{&date}},
		},
		{
			name:   "In case (pointer date)",
			fields: fields{In: []*time.Time{&date}},
			want:   &GenericFilter{In: []*time.Time{&date}},
		},
		{
			name:   "In case (date)",
			fields: fields{In: []time.Time{date}},
			want:   &GenericFilter{In: []*time.Time{&date}},
		},
		{
			name:   "In case (nano case pointer string)",
			fields: fields{In: []*string{&dateStr2}},
			want:   &GenericFilter{In: []*time.Time{&date2}},
		},
		{
			name:    "In not a date",
			fields:  fields{In: []*string{&notADate}},
			wantErr: true,
		},
		{
			name:    "In not a date 2",
			fields:  fields{In: []string{notADate}},
			wantErr: true,
		},
		{
			name:    "In not a date 3",
			fields:  fields{In: 9},
			wantErr: true,
		},
		{
			name:   "In nil",
			fields: fields{In: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "NotIn case (pointer string)",
			fields: fields{NotIn: []*string{&dateStr}},
			want:   &GenericFilter{NotIn: []*time.Time{&date}},
		},
		{
			name:   "NotIn case (string)",
			fields: fields{NotIn: []string{dateStr}},
			want:   &GenericFilter{NotIn: []*time.Time{&date}},
		},
		{
			name:   "NotIn case (pointer date)",
			fields: fields{NotIn: []*time.Time{&date}},
			want:   &GenericFilter{NotIn: []*time.Time{&date}},
		},
		{
			name:   "NotIn case (date)",
			fields: fields{NotIn: []time.Time{date}},
			want:   &GenericFilter{NotIn: []*time.Time{&date}},
		},
		{
			name:   "NotIn case (nano case pointer string)",
			fields: fields{NotIn: []*string{&dateStr2}},
			want:   &GenericFilter{NotIn: []*time.Time{&date2}},
		},
		{
			name:    "NotIn not a date",
			fields:  fields{NotIn: []string{notADate}},
			wantErr: true,
		},
		{
			name:    "NotIn not a date 2",
			fields:  fields{NotIn: []string{notADate}},
			wantErr: true,
		},
		{
			name:    "NotIn not a date 3",
			fields:  fields{NotIn: 9},
			wantErr: true,
		},
		{
			name:   "NotIn nil",
			fields: fields{NotIn: nil},
			want:   &GenericFilter{},
		},
		{
			name:   "IsNull case",
			fields: fields{IsNull: true},
			want:   &GenericFilter{IsNull: true},
		},
		{
			name:   "IsNotNull case",
			fields: fields{IsNotNull: true},
			want:   &GenericFilter{IsNotNull: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DateFilter{
				Eq:        tt.fields.Eq,
				NotEq:     tt.fields.NotEq,
				Gte:       tt.fields.Gte,
				NotGte:    tt.fields.NotGte,
				Gt:        tt.fields.Gt,
				NotGt:     tt.fields.NotGt,
				Lte:       tt.fields.Lte,
				NotLte:    tt.fields.NotLte,
				Lt:        tt.fields.Lt,
				NotLt:     tt.fields.NotLt,
				In:        tt.fields.In,
				NotIn:     tt.fields.NotIn,
				IsNull:    tt.fields.IsNull,
				IsNotNull: tt.fields.IsNotNull,
			}
			got, err := d.GetGenericFilter()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateFilter.GetGenericFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
