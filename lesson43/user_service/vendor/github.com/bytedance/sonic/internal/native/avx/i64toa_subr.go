// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__i64toa = 64
)

const (
    _stack__i64toa = 8
)

const (
    _size__i64toa = 2272
)

var (
    _pcsp__i64toa = [][2]uint32{
        {1, 0},
        {170, 8},
        {171, 0},
        {505, 8},
        {506, 0},
        {637, 8},
        {638, 0},
        {1101, 8},
        {1102, 0},
        {1238, 8},
        {1239, 0},
        {1540, 8},
        {1541, 0},
        {1901, 8},
        {1902, 0},
        {2268, 8},
        {2270, 0},
    }
)

var _cfunc_i64toa = []loader.CFunc{
    {"_i64toa_entry", 0,  _entry__i64toa, 0, nil},
    {"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
}
