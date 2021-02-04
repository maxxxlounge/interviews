package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/maxxxlounge/interviews/SouthAfricanNumber/handler"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/numbermanager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var response []byte

type mockWriter struct {
	HTTPHeaderResponse http.Header
	WriteResponseCode  int
	Error              error
	HeaderCode         int
}

func (m mockWriter) Header() http.Header {
	return m.HTTPHeaderResponse
}

func (m mockWriter) Write(resp []byte) (int, error) {
	response = resp
	return m.WriteResponseCode, m.Error
}

func (m mockWriter) WriteHeader(statusCode int) {
	m.HTTPHeaderResponse["status"] = []string{fmt.Sprintf("%v", statusCode)}
}

func TestCheck(t *testing.T) {
	tests := map[string]struct {
		writer             mockWriter
		expectedHeaderCode int
		number             string
		Row                *numbermanager.Row
	}{
		"EmptyNumber": {
			number: "",
			writer: mockWriter{
				HTTPHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
			Row:                nil,
		},
		"NumberOK": {
			number: "27831234567",
			writer: mockWriter{
				HTTPHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusOK,
			Row: &numbermanager.Row{
				Original: "27831234567",
				Changed:  "",
				Errors:   nil,
				Type:     numbermanager.ValidFirstAttempt,
			},
		},
		"NumberKO": {
			number: "278312345671",
			writer: mockWriter{
				HTTPHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
			Row: &numbermanager.Row{
				Original: "278312345671",
				Changed:  "27831234567",
				Errors:   []string{numbermanager.ErrorCutExtraDigits},
				Type:     numbermanager.InvalidButFixable,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler.Check(tc.writer, tc.number)
			status, err := strconv.Atoi(tc.writer.HTTPHeaderResponse["status"][0])
			require.Nil(t, err)
			assert.Equal(t, tc.expectedHeaderCode, status)
			if tc.Row == nil {
				return
			}
			r := numbermanager.Row{}
			err = json.Unmarshal(response, &r)
			require.Nil(t, err)
			assert.EqualValues(t, *tc.Row, r)
		})
	}
}

func TestShowNumbers(t *testing.T) {
	testMap := map[string]*numbermanager.Row{
		"123123": {
			Original: "1234567",
			Changed:  "27831234567",
			Errors:   []string{numbermanager.ErrorMissingPartialPrefix},
			Type:     numbermanager.InvalidCritical,
		},
		"1231234": {
			Original: "12341234567",
			Changed:  "27831234567",
			Errors:   []string{numbermanager.ErrorWrongPrefix},
			Type:     numbermanager.InvalidCritical,
		},
		"1231233": {
			Original: "1234567",
			Changed:  "27831234567",
			Errors:   []string{numbermanager.ErrorMissingPartialPrefix},
			Type:     numbermanager.InvalidButFixable,
		},
		"12131233": {
			Original: "1234567",
			Changed:  "27831234567",
			Errors:   []string{numbermanager.ErrorMissingPartialPrefix},
			Type:     numbermanager.ValidFirstAttempt,
		},
	}

	tests := map[string]struct {
		hasError           bool
		writer             mockWriter
		expectedHeaderCode int
		numberListMap      map[string]*numbermanager.Row
	}{
		"UnmarshalError": {
			numberListMap: nil,
			writer: mockWriter{
				HTTPHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
		},
		"AllNumbers": {
			numberListMap: testMap,
			writer: mockWriter{
				HTTPHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
		},
		"InvalidButFixable": {
			numberListMap: map[string]*numbermanager.Row{
				"123123": {
					Original: "1234567",
					Changed:  "27831234567",
					Errors:   nil,
					Type:     numbermanager.InvalidButFixable,
				},
			},
			writer: mockWriter{
				HTTPHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler.ShowNumbers(tc.writer, tc.numberListMap)
			returnedMapList := make(map[string]*numbermanager.Row)
			err := json.Unmarshal(response, &returnedMapList)
			require.Nil(t, err)
			assert.EqualValues(t, tc.numberListMap, returnedMapList)
		})
	}
}
