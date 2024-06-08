// Copyright 2024 Michael F. Collins, III
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of thi software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDER BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/mfcollins3/go-azure-functions/pkg/http"
	"github.com/mfcollins3/go-azure-functions/pkg/timer"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Request struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}

func (req Request) Get(key string, output interface{}) error {
	return mapstructure.Decode(req.Data[key], output)
}

func (req Request) GetJSON(key string, output interface{}) error {
	str, ok := req.Data[key].(string)
	if !ok {
		return fmt.Errorf("data with key \"%v\" was not a string")
	}

	return json.Unmarshal([]byte(str), output)
}

func (req Request) TimerInfo(key string) (info timer.Info, err error) {
	var dto timer.InfoDTO
	err = req.Get(key, &dto)
	if err != nil {
		return
	}

	var last, lastUpdated, next time.Time
	if len(dto.ScheduleStatus.Last) > 0 {
		last, err = time.Parse(time.RFC3339Nano, dto.ScheduleStatus.Last)
		if err != nil {
			return
		}
	}

	if len(dto.ScheduleStatus.Next) > 0 {
		next, err = time.Parse(time.RFC3339Nano, dto.ScheduleStatus.Next)
		if err != nil {
			return
		}
	}

	if len(dto.ScheduleStatus.LastUpdated) > 0 {
		lastUpdated, err = time.Parse(time.RFC3339Nano, dto.ScheduleStatus.LastUpdated)
		if err != nil {
			return
		}
	}

	info = timer.Info{
		Schedule: timer.Schedule{
			AdjustForDST: dto.Schedule.AdjustForDST,
		},
		ScheduleStatus: timer.ScheduleStatus{
			Last:        last,
			LastUpdated: lastUpdated,
			Next:        next,
		},
		IsPastDue: dto.IsPastDue,
	}
	return
}

func (req Request) HTTPRequest(key string) (
	httpRequest http.Request,
	err error,
) {
	var request http.Request
	err = req.Get(key, &request)
	return request, err
}
