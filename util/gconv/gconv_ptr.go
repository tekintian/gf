// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

// PtrAny creates and returns an interface{} pointer variable to this value.
func PtrAny(val interface{}) *interface{} {
	return &val
}

// PtrString creates and returns a string pointer variable to this value.
func PtrString(val interface{}) *string {
	v := String(val)
	return &v
}

// PtrBool creates and returns a bool pointer variable to this value.
func PtrBool(val interface{}) *bool {
	v := Bool(val)
	return &v
}

// PtrInt creates and returns an int pointer variable to this value.
func PtrInt(val interface{}) *int {
	v := Int(val)
	return &v
}

// PtrInt8 creates and returns an int8 pointer variable to this value.
func PtrInt8(val interface{}) *int8 {
	v := Int8(val)
	return &v
}

// PtrInt16 creates and returns an int16 pointer variable to this value.
func PtrInt16(val interface{}) *int16 {
	v := Int16(val)
	return &v
}

// PtrInt32 creates and returns an int32 pointer variable to this value.
func PtrInt32(val interface{}) *int32 {
	v := Int32(val)
	return &v
}

// PtrInt64 creates and returns an int64 pointer variable to this value.
func PtrInt64(val interface{}) *int64 {
	v := Int64(val)
	return &v
}

// PtrUint creates and returns an uint pointer variable to this value.
func PtrUint(val interface{}) *uint {
	v := Uint(val)
	return &v
}

// PtrUint8 creates and returns an uint8 pointer variable to this value.
func PtrUint8(val interface{}) *uint8 {
	v := Uint8(val)
	return &v
}

// PtrUint16 creates and returns an uint16 pointer variable to this value.
func PtrUint16(val interface{}) *uint16 {
	v := Uint16(val)
	return &v
}

// PtrUint32 creates and returns an uint32 pointer variable to this value.
func PtrUint32(val interface{}) *uint32 {
	v := Uint32(val)
	return &v
}

// PtrUint64 creates and returns an uint64 pointer variable to this value.
func PtrUint64(val interface{}) *uint64 {
	v := Uint64(val)
	return &v
}

// PtrFloat32 creates and returns a float32 pointer variable to this value.
func PtrFloat32(val interface{}) *float32 {
	v := Float32(val)
	return &v
}

// PtrFloat64 creates and returns a float64 pointer variable to this value.
func PtrFloat64(val interface{}) *float64 {
	v := Float64(val)
	return &v
}
