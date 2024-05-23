// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import (
	"time"

	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/os/gtime"
)

// Time converts `val` to time.Time.
func Time(val interface{}, format ...string) time.Time {
	// It's already this type.
	if len(format) == 0 {
		if v, ok := val.(time.Time); ok {
			return v
		}
	}
	if t := GTime(val, format...); t != nil {
		return t.Time
	}
	return time.Time{}
}

// Duration converts `val` to time.Duration.
// If `val` is string, then it uses time.ParseDuration to convert it.
// If `val` is numeric, then it converts `val` as nanoseconds.
func Duration(val interface{}) time.Duration {
	// It's already this type.
	if v, ok := val.(time.Duration); ok {
		return v
	}
	s := String(val)
	if !utils.IsNumeric(s) {
		d, _ := gtime.ParseDuration(s)
		return d
	}
	return time.Duration(Int64(val))
}

// GTime converts `val` to *gtime.Time.
// The parameter `format` can be used to specify the format of `val`.
// It returns the converted value that matched the first format of the formats slice.
// If no `format` given, it converts `val` using gtime.NewFromTimeStamp if `val` is numeric,
// or using gtime.StrToTime if `val` is string.
func GTime(val interface{}, format ...string) *gtime.Time {
	if val == nil {
		return nil
	}
	if v, ok := val.(iGTime); ok {
		return v.GTime(format...)
	}
	// It's already this type.
	if len(format) == 0 {
		if v, ok := val.(*gtime.Time); ok {
			return v
		}
		if t, ok := val.(time.Time); ok {
			return gtime.New(t)
		}
		if t, ok := val.(*time.Time); ok {
			return gtime.New(t)
		}
	}
	s := String(val)
	if len(s) == 0 {
		return gtime.New()
	}
	// Priority conversion using given format.
	if len(format) > 0 {
		for _, item := range format {
			t, err := gtime.StrToTimeFormat(s, item)
			if t != nil && err == nil {
				return t
			}
		}
		return nil
	}
	if utils.IsNumeric(s) {
		return gtime.NewFromTimeStamp(Int64(s))
	} else {
		t, _ := gtime.StrToTime(s)
		return t
	}
}
