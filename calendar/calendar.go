// Copyright 2019 Michael DOUBEZ
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package calendar

import (
	"fmt"
	"net/http"
	//"github.com/prometheus/client_golang/prometheus"
)

// A calendar is able to expose metrics and a base page
// describing the calendar
type Calendar interface {
	// Display info about calendar
	Handle(http.ResponseWriter)
}

type CalendarList struct {
	calendar map[string]*Calendar
}

// NewCalendarList return an empty collection of calendars to be exported
func NewCalendarList() *CalendarList {
	return &CalendarList{
		calendar: map[string]*Calendar{},
	}
}

// Handle Calendar list http request on calendar path
func (calendarList *CalendarList) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, ok := r.URL.Query()["name"]
		if !ok {
			w.Write([]byte(`<html><head><title>Calendar Exporter</title></head><body>`))
			for name := range calendarList.calendar {
				w.Write([]byte(name))
				w.Write([]byte("<br />\n"))
			}
			w.Write([]byte(`</body></head>`))
			return
		}

		if calendar, ok := calendarList.calendar[name[0]]; !ok {
			http.Error(w, fmt.Sprintf("unknown calendar %v\n", calendar), http.StatusNotFound)
		} else {
			(*calendar).Handle(w)
		}
	})
}
