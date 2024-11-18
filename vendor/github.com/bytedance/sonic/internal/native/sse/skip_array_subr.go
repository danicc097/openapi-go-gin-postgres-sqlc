// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__skip_array = 272
)

const (
    _stack__skip_array = 184
)

const (
    _size__skip_array = 14468
)

var (
    _pcsp__skip_array = [][2]uint32{
        {0x1, 0},
        {0x6, 8},
        {0x8, 16},
        {0xa, 24},
        {0xc, 32},
        {0xd, 40},
        {0x14, 48},
        {0x35ea, 184},
        {0x35eb, 48},
        {0x35ed, 40},
        {0x35ef, 32},
        {0x35f1, 24},
        {0x35f3, 16},
        {0x35f4, 8},
        {0x35f5, 0},
        {0x3884, 184},
    }
)

var _cfunc_skip_array = []loader.CFunc{
    {"_skip_array_entry", 0,  _entry__skip_array, 0, nil},
    {"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
}
