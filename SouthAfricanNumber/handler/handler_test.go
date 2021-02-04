package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var response []byte

type mockWriter struct {
	HttpHeaderResponse http.Header
	WriteResponseCode  int
	Error              error
	HeaderCode         int
}

func (m mockWriter) Header() http.Header {
	return m.HttpHeaderResponse
}

func (m mockWriter) Write(resp []byte) (int, error) {
	response = resp
	return m.WriteResponseCode, m.Error
}

func (m mockWriter) WriteHeader(statusCode int) {
	m.HttpHeaderResponse["status"] = []string{fmt.Sprintf("%v", statusCode)}
}

func TestCheck(t *testing.T) {
	tests := map[string]struct {
		writer             mockWriter
		expectedHeaderCode int
		number             string
		Row                *NumberManager.Row
	}{
		"EmptyNumber": {
			number: "",
			writer: mockWriter{
				HttpHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
			Row:                nil,
		},
		"NumberOK": {
			number: "27831234567",
			writer: mockWriter{
				HttpHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusOK,
			Row: &NumberManager.Row{
				Original: "27831234567",
				Changed:  "",
				Errors:   nil,
				Type:     NumberManager.ValidFirstAttempt,
			},
		},
		"NumberKO": {
			number: "278312345671",
			writer: mockWriter{
				HttpHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
			Row: &NumberManager.Row{
				Original: "278312345671",
				Changed:  "27831234567",
				Errors:   []string{NumberManager.ErrorCutExtraDigits},
				Type:     NumberManager.InvalidButFixable,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler.Check(tc.writer, tc.number)
			status, err := strconv.Atoi(tc.writer.HttpHeaderResponse["status"][0])
			require.Nil(t, err)
			assert.Equal(t, tc.expectedHeaderCode, status)
			if tc.Row == nil {
				return
			}
			r := NumberManager.Row{}
			err = json.Unmarshal(response, &r)
			require.Nil(t, err)
			assert.EqualValues(t, *tc.Row, r)
		})
	}

}

func TestShowNumbers(t *testing.T) {
	testMap := map[string]*NumberManager.Row{
		"123123": &NumberManager.Row{
			Original: "1234567",
			Changed:  "27831234567",
			Errors:   []string{NumberManager.ErrorMissingPartialPrefix},
			Type:     NumberManager.InvalidCritical,
		},
		"1231234": &NumberManager.Row{
			Original: "12341234567",
			Changed:  "27831234567",
			Errors:   []string{NumberManager.ErrorWrongPrefix},
			Type:     NumberManager.InvalidCritical,
		},
		"1231233": &NumberManager.Row{
			Original: "1234567",
			Changed:  "27831234567",
			Errors:   []string{NumberManager.ErrorMissingPartialPrefix},
			Type:     NumberManager.InvalidButFixable,
		},
		"12131233": &NumberManager.Row{
			Original: "1234567",
			Changed:  "27831234567",
			Errors:   []string{NumberManager.ErrorMissingPartialPrefix},
			Type:     NumberManager.ValidFirstAttempt,
		},
	}

	tests := map[string]struct {
		hasError           bool
		writer             mockWriter
		expectedHeaderCode int
		numberListMap      map[string]*NumberManager.Row
	}{
		"UnmarshalError": {
			numberListMap: nil,
			writer: mockWriter{
				HttpHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
		},
		"AllNumbers": {
			numberListMap: testMap,
			writer: mockWriter{
				HttpHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
		},
		"InvalidButFixable": {
			numberListMap: map[string]*NumberManager.Row{
				"123123": &NumberManager.Row{
					Original: "1234567",
					Changed:  "27831234567",
					Errors:   nil,
					Type:     NumberManager.InvalidButFixable,
				},
			},
			writer: mockWriter{
				HttpHeaderResponse: http.Header{},
				WriteResponseCode:  http.StatusOK,
				Error:              nil,
			},
			expectedHeaderCode: http.StatusBadRequest,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler.ShowNumbers(tc.writer, tc.numberListMap)
			returnedMapList := make(map[string]*NumberManager.Row)
			err := json.Unmarshal(response, &returnedMapList)
			require.Nil(t, err)
			assert.EqualValues(t, tc.numberListMap, returnedMapList)
		})
	}

}
