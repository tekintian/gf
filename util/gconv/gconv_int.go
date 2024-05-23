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

// Int converts `val` to int.
func Int(val interface{}) int {
	if val == nil {
		return 0
	}
	if v, ok := val.(int); ok {
		return v
	}
	return int(Int64(val))
}

// Int8 converts `val` to int8.
func Int8(val interface{}) int8 {
	if val == nil {
		return 0
	}
	if v, ok := val.(int8); ok {
		return v
	}
	return int8(Int64(val))
}

// Int16 converts `val` to int16.
func Int16(val interface{}) int16 {
	if val == nil {
		return 0
	}
	if v, ok := val.(int16); ok {
		return v
	}
	return int16(Int64(val))
}

// Int32 converts `val` to int32.
func Int32(val interface{}) int32 {
	if val == nil {
		return 0
	}
	if v, ok := val.(int32); ok {
		return v
	}
	return int32(Int64(val))
}

// Int64 converts `val` to int64.
func Int64(val interface{}) int64 {
	if val == nil {
		return 0
	}
	switch value := val.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return gbinary.DecodeToInt64(value)
	default:
		if f, ok := value.(iInt64); ok {
			return f.Int64()
		}
		var (
			s       = String(value)
			isMinus = false
		)
		if len(s) > 0 {
			if s[0] == '-' {
				isMinus = true
				s = s[1:]
			} else if s[0] == '+' {
				s = s[1:]
			}
		}
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseInt(s[2:], 16, 64); e == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseInt(s, 10, 64); e == nil {
			if isMinus {
				return -v
			}
			return v
		}
		// Float64
		if valueInt64 := Float64(value); math.IsNaN(valueInt64) {
			return 0
		} else {
			return int64(valueInt64)
		}
	}
}
