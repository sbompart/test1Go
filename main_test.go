package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_GetType(t *testing.T) {
	var tests = []struct {
		caseName string
		strVal string
		want string
		err error
	}{
		{
			"GetType - String must contain 3 chars min",
			"A0",
			"",
			errors.New("GetType - String must contain 3 chars min"),
		},
		{
			"GetType - First leter is not N or A",
			"Z01",
			"",
			errors.New("GetType - First leter is not N or A"),
		},
		{
			"GetType - Success",
			"A0111",
			"11",
			nil,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s: %s", tt.caseName, tt.strVal)
		t.Run(testname, func(t *testing.T) {
			tlvMap, err := getType(tt.strVal)
			if !reflect.DeepEqual(tlvMap, tt.want) || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got %+v, want %+v and error is %+v", tlvMap, tt.want, err)
			}
		})
	}
}

func Test_GetLength(t *testing.T) {
	var tests = []struct {
		caseName string
		strVal string
		want, want2 string
		err error
	}{
		{
			"GetLength - String must contain 2 chars min",
			"0",
			"", "",
			errors.New("GetLength - String must contain 2 chars min"),
		},
		{
			"GetLength - Success",
			"11B2131",
			"11", "B2131",
			nil,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s: %s", tt.caseName, tt.strVal)
		t.Run(testname, func(t *testing.T) {
			tlvMap, tlvMap2, err := getLenght(tt.strVal)
			if !reflect.DeepEqual(tlvMap, tt.want) || !reflect.DeepEqual(tlvMap2, tt.want2) || !reflect.DeepEqual(tlvMap, tt.want) || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got %+v, want %+v, want2 %+v and error is %+v", tlvMap, tt.want, tlvMap2, err)
			}
		})
	}
}

func Test_GetValue(t *testing.T) {
	var tests = []struct {
		caseName string
		strVal string; lgt int
		want, want2 string
		err error
	}{
		{
			"GetValue - Error",
			"B12313", 9,
			"", "",
			errors.New("GetValue - The value contains fewer characters than necessary."),
		},
		{
			"GetValue - Error",
			"00N11", 2,
			"00", "N11",
			nil,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s: %s", tt.caseName, tt.strVal)
		t.Run(testname, func(t *testing.T) {
			tlvMap, tlvMap2, err := getValue(tt.strVal, tt.lgt)
			if !reflect.DeepEqual(tlvMap, tt.want) || !reflect.DeepEqual(tlvMap2, tt.want2) || !reflect.DeepEqual(tlvMap, tt.want) || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got %+v, want %+v, want2 %+v and error is %+v", tlvMap, tt.want, tlvMap2, err)
			}
		})
	}
}

func Test_LoopForStr(t *testing.T) {
	var tests = []struct {
		caseName string
		strVal string
		want []string
		err error
	}{
		{
			"LoopForStr - Empty string",
			"",
			nil,
			errors.New("Erorr empty string"),
		},
		{
			"LoopForStr - GetType error",
			"A0",
			nil,
			errors.New("GetType - String must contain 3 chars min"),
		},
		{
			"LoopForStr - GetLength error",
			"A020",
			nil,
			errors.New("GetLength - String must contain 2 chars min"),
		},
		{
			"LoopForStr - GetValue error",
			"A020300",
			nil,
			errors.New("GetValue - The value contains fewer characters than necessary."),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s: %s", tt.caseName, tt.strVal)
		t.Run(testname, func(t *testing.T) {
			tlvMap, err := loopForStr(tt.strVal)
			if !reflect.DeepEqual(tlvMap, tt.want) || !reflect.DeepEqual(tlvMap, tt.want) || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got %+v, want %+v and error is %+v", tlvMap, tt.want, err)
			}
		})
	}
}