package main

import (
	"reflect"
	"testing"
)

func TestCommand_Delimit(t *testing.T) {
	type fields struct {
		d string
		f int
		s bool
	}
	tests := []struct {
		name       string
		fields     fields
		configFile []byte
		want       []string
	}{
		{"First test", fields{d: " "}, []byte("Основные объявления флагов"), []string{"Основные", "объявления", "флагов"}},
		{"Second test", fields{d: ")"}, []byte("Основные)объявления)флагов"), []string{"Основные", "объявления", "флагов"}},
		{"Third test", fields{d: "."}, []byte("Основные.объявления флагов"), []string{"Основные", "объявления флагов"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				d: tt.fields.d,
				f: tt.fields.f,
				s: tt.fields.s,
			}
			if got := c.Delimit(tt.configFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_Fields(t *testing.T) {
	type fields struct {
		d string
		f int
		s bool
	}

	tests := []struct {
		name        string
		fields      fields
		configLines []string
		want        []string
	}{
		{"First test", fields{f: 1}, []string{"Основные", "объявления", "флагов"}, []string{"Основные", "объявления", "флагов"}},
		{"Second test", fields{f: 2}, []string{"Основные объявления", "флагов понятие"}, []string{"объявления", "понятие"}},
		{"Second test", fields{f: 10}, []string{"Основные объявления", "флагов понятие"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				d: tt.fields.d,
				f: tt.fields.f,
				s: tt.fields.s,
			}
			if got := c.Fields(tt.configLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_Separated(t *testing.T) {
	type fields struct {
		d string
		f int
		s bool
	}
	tests := []struct {
		name       string
		fields     fields
		configFile []byte
		want       []string
	}{
		{"First test", fields{d: " ", s: true}, []byte("Основные флагов\nосновные"), []string{"Основные флагов"}},
		{"Second test", fields{d: ")", s: true}, []byte("Основные флагов\nосновные)\nда"), []string{"основные)"}},
		{"Third test", fields{d: " ", s: true}, []byte("Основныефлагов\nосновные"), []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				d: tt.fields.d,
				f: tt.fields.f,
				s: tt.fields.s,
			}
			if got := c.Separated(tt.configFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Separated() = %v, want %v", got, tt.want)
			}
		})
	}
}
