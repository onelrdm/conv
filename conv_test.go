package conv

import (
	"testing"
)

type stringer struct{}

func (stringer) String() string {
	return "1"
}

func TestShouldInt64(t *testing.T) {
	tests := []struct {
		name    string
		arg     interface{}
		want    int64
		wantErr bool
	}{
		{name: "string", arg: "1", want: 1, wantErr: false},
		{name: "stringer", arg: stringer{}, want: 1, wantErr: false},
		{name: "bool", arg: true, want: 1, wantErr: false},
		{name: "bool", arg: false, want: 0, wantErr: false},
		{name: "uint", arg: uint(1), want: 1, wantErr: false},
		{name: "uint8", arg: uint8(1), want: 1, wantErr: false},
		{name: "uint16", arg: uint16(1), want: 1, wantErr: false},
		{name: "uint32", arg: uint32(1), want: 1, wantErr: false},
		{name: "uint64", arg: uint64(1), want: 1, wantErr: false},
		{name: "int", arg: 1, want: 1, wantErr: false},
		{name: "int8", arg: int8(1), want: 1, wantErr: false},
		{name: "int16", arg: int16(1), want: 1, wantErr: false},
		{name: "int32", arg: int32(1), want: 1, wantErr: false},
		{name: "int64", arg: int64(1), want: 1, wantErr: false},
		{name: "float32", arg: float32(1), want: 1, wantErr: false},
		{name: "float64", arg: float64(1), want: 1, wantErr: false},
		{name: "unknown", arg: struct{}{}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShouldInt64(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShouldInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShouldInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustInt64(t *testing.T) {
	tests := []struct {
		name      string
		arg       interface{}
		want      int64
		wantPanic bool
	}{
		{name: "int", arg: 1, want: 1},
		{name: "unknown", arg: struct{}{}, wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("MustInt64() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()
			if got := MustInt64(tt.arg); got != tt.want {
				t.Errorf("MustInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShouldFloat64(t *testing.T) {
	tests := []struct {
		name    string
		arg     interface{}
		want    float64
		wantErr bool
	}{
		{name: "string", arg: "1", want: 1, wantErr: false},
		{name: "stringer", arg: stringer{}, want: 1, wantErr: false},
		{name: "bool", arg: true, want: 1, wantErr: false},
		{name: "bool", arg: false, want: 0, wantErr: false},
		{name: "uint", arg: uint(1), want: 1, wantErr: false},
		{name: "uint8", arg: uint8(1), want: 1, wantErr: false},
		{name: "uint16", arg: uint16(1), want: 1, wantErr: false},
		{name: "uint32", arg: uint32(1), want: 1, wantErr: false},
		{name: "uint64", arg: uint64(1), want: 1, wantErr: false},
		{name: "int", arg: 1, want: 1, wantErr: false},
		{name: "int8", arg: int8(1), want: 1, wantErr: false},
		{name: "int16", arg: int16(1), want: 1, wantErr: false},
		{name: "int32", arg: int32(1), want: 1, wantErr: false},
		{name: "int64", arg: int64(1), want: 1, wantErr: false},
		{name: "float32", arg: float32(1), want: 1, wantErr: false},
		{name: "float64", arg: float64(1), want: 1, wantErr: false},
		{name: "unknown", arg: struct{}{}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShouldFloat64(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShouldFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShouldFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustFloat64(t *testing.T) {
	tests := []struct {
		name      string
		arg       interface{}
		want      float64
		wantPanic bool
	}{
		{name: "int", arg: 1, want: 1},
		{name: "unknown", arg: struct{}{}, wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("MustFloat64() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()
			if got := MustFloat64(tt.arg); got != tt.want {
				t.Errorf("MustFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShouldString(t *testing.T) {
	tests := []struct {
		name    string
		arg     interface{}
		want    string
		wantErr bool
	}{
		{name: "string", arg: "1", want: "1", wantErr: false},
		{name: "stringer", arg: stringer{}, want: "1", wantErr: false},
		{name: "[]byte", arg: []byte("1"), want: "1", wantErr: false},
		{name: "bool", arg: true, want: "true", wantErr: false},
		{name: "bool", arg: false, want: "false", wantErr: false},
		{name: "uint", arg: uint(1), want: "1", wantErr: false},
		{name: "uint8", arg: uint8(1), want: "1", wantErr: false},
		{name: "uint16", arg: uint16(1), want: "1", wantErr: false},
		{name: "uint32", arg: uint32(1), want: "1", wantErr: false},
		{name: "uint64", arg: uint64(1), want: "1", wantErr: false},
		{name: "int", arg: 1, want: "1", wantErr: false},
		{name: "int8", arg: int8(1), want: "1", wantErr: false},
		{name: "int16", arg: int16(1), want: "1", wantErr: false},
		{name: "int32", arg: int32(1), want: "1", wantErr: false},
		{name: "int64", arg: int64(1), want: "1", wantErr: false},
		{name: "float32", arg: float32(1), want: "1.000000", wantErr: false},
		{name: "float64", arg: float64(1), want: "1.000000", wantErr: false},
		{name: "unknown", arg: struct{}{}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShouldString(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShouldString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShouldString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustString(t *testing.T) {
	tests := []struct {
		name      string
		arg       interface{}
		want      string
		wantPanic bool
	}{
		{name: "int", arg: 1, want: "1"},
		{name: "unknown", arg: struct{}{}, wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("MustString() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()
			if got := MustString(tt.arg); got != tt.want {
				t.Errorf("MustString() = %v, want %v", got, tt.want)
			}
		})
	}
}
