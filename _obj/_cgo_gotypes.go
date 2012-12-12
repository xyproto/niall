// Created by cgo - DO NOT EDIT

package niall

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int) { *dst = syscall.Errno(x) }
type _Ctype_int int32

type _Ctype_char int8

type _Ctype_void [0]byte
const _Cconst_BUFSIZ = 0x2000

func _Cfunc_CString(string) *_Ctype_char
func _Cfunc_GoString(*_Ctype_char) string
func _Cfunc_Niall_Free()
func _Cfunc_Niall_Init()
func _Cfunc_Niall_Learn(*_Ctype_char)
func _Cfunc_Niall_ListDictionary()
func _Cfunc_Niall_NewDictionary()
func _Cfunc_Niall_Reply(*_Ctype_char, _Ctype_int)
