// +build unit

package utils

import (
	"testing"
)

func TestCalculateMutationComplexity(t *testing.T) {
	type args struct {
		childComplexity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero child",
			args: args{
				childComplexity: 0,
			},
			want: 10,
		},
		{
			name: "child",
			args: args{
				childComplexity: 2,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMutationComplexity(tt.args.childComplexity); got != tt.want {
				t.Errorf("CalculateMutationComplexity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateQuerySimpleStructComplexity(t *testing.T) {
	type args struct {
		childComplexity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero child",
			args: args{
				childComplexity: 0,
			},
			want: 1,
		},
		{
			name: "child",
			args: args{
				childComplexity: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateQuerySimpleStructComplexity(tt.args.childComplexity); got != tt.want {
				t.Errorf("CalculateQuerySimpleStructComplexity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateQueryConnectionComplexity(t *testing.T) {
	stringStar := func(s string) *string { return &s }
	intStar := func(s int) *int { return &s }
	type args struct {
		childComplexity int
		after           *string
		before          *string
		first           *int
		last            *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero child and no pagination",
			args: args{
				childComplexity: 0,
			},
			want: 2,
		},
		{
			name: "child and before and last pagination",
			args: args{
				childComplexity: 5,
				before:          stringStar("fake"),
				last:            intStar(5),
			},
			want: 27,
		},
		{
			name: "child and before and last pagination (2)",
			args: args{
				childComplexity: 1,
				before:          stringStar("fake"),
				last:            intStar(5),
			},
			want: 7,
		},
		{
			name: "child and empty before and last pagination",
			args: args{
				childComplexity: 5,
				before:          stringStar(""),
				last:            intStar(5),
			},
			want: 52,
		},
		{
			name: "child and before and empty last pagination",
			args: args{
				childComplexity: 5,
				before:          stringStar("fake"),
				last:            nil,
			},
			want: 52,
		},
		{
			name: "child and after and first pagination",
			args: args{
				childComplexity: 5,
				after:           stringStar("fake"),
				first:           intStar(5),
			},
			want: 27,
		},
		{
			name: "child and after and first pagination (2)",
			args: args{
				childComplexity: 1,
				after:           stringStar("fake"),
				first:           intStar(5),
			},
			want: 7,
		},
		{
			name: "child and no after and first pagination",
			args: args{
				childComplexity: 5,
				after:           nil,
				first:           intStar(5),
			},
			want: 27,
		},
		{
			name: "child and full pagination (priority to before and last)",
			args: args{
				childComplexity: 5,
				before:          stringStar("fake"),
				after:           stringStar("fake"),
				last:            intStar(5),
				first:           intStar(25),
			},
			want: 27,
		},
		{
			name: "child without any pagination",
			args: args{
				childComplexity: 5,
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateQueryConnectionComplexity(tt.args.childComplexity, tt.args.after, tt.args.before, tt.args.first, tt.args.last); got != tt.want {
				t.Errorf("CalculateQueryConnectionComplexity() = %v, want %v", got, tt.want)
			}
		})
	}
}
