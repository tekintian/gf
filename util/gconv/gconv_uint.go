// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import (
	"math"
	"strconv"

	"github.com/gogf/gf/v2/encoding/gbinary"
)

// Uint converts `val` to uint.
func Uint(val interface{}) uint {
	if val == nil {
		return 0
	}
	if v, ok := val.(uint); ok {
		return v
	}
	return uint(Uint64(val))
}

// Uint8 converts `val` to uint8.
func Uint8(val interface{}) uint8 {
	if val == nil {
		return 0
	}
	if v, ok := val.(uint8); ok {
		return v
	}
	return uint8(Uint64(val))
}

// Uint16 converts `val` to uint16.
func Uint16(val interface{}) uint16 {
	if val == nil {
		return 0
	}
	if v, ok := val.(uint16); ok {
		return v
	}
	return uint16(Uint64(val))
}

// Uint32 converts `val` to uint32.
func Uint32(val interface{}) uint32 {
	if val == nil {
		return 0
	}
	if v, ok := val.(uint32); ok {
		return v
	}
	return uint32(Uint64(val))
}

// Uint64 converts `val` to uint64.
func Uint64(val interface{}) uint64 {
	if val == nil {
		return 0
	}
	switch value := val.(type) {
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return gbinary.DecodeToUint64(value)
	default:
		if f, ok := value.(iUint64); ok {
			return f.Uint64()
		}
		s := String(value)
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseUint(s[2:], 16, 64); e == nil {
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseUint(s, 10, 64); e == nil {
			return v
		}
		// Float64
		if valueFloat64 := Float64(value); math.IsNaN(valueFloat64) {
			return 0
		} else {
			return uint64(valueFloat64)
		}
	}
}
