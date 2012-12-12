// Created by cgo - DO NOT EDIT

//line niall.go:1
package niall
//line niall.go:9

//line niall.go:8
import (
	"bytes"
	"fmt"
)
//line niall.go:14

//line niall.go:13
var (
	WarningFunction	= fmt.Print
	ErrorFunction	= fmt.Print
	PrintFunction	= fmt.Print
)
//line niall.go:21

//line niall.go:20
func niall_go_warning(msg *_Ctype_char) {
	WarningFunction("WARNING: " + _Cfunc_GoString(msg))
}
//line niall.go:26

//line niall.go:25
func niall_go_error(msg *_Ctype_char) {
	ErrorFunction("ERROR: " + _Cfunc_GoString(msg))
}
//line niall.go:31

//line niall.go:30
func niall_go_print(msg *_Ctype_char) {
	PrintFunction(_Cfunc_GoString(msg))
}
//line niall.go:37

//line niall.go:36
func newCStringBuffer() (*_Ctype_char, _Ctype_int) {
	temp := make([]byte, 0, _Cconst_BUFSIZ)
	b := bytes.NewBuffer(temp)
	buf := _Cfunc_CString(b.String())
	return buf, _Ctype_int(cap(temp))
}
//line niall.go:45

//line niall.go:44
func Learn(s string) {
//line niall.go:44
	_Cfunc_Niall_Learn(_Cfunc_CString(s))
//line niall.go:46
}
//line niall.go:50

//line niall.go:49
func Talk() string {
			buf, size := newCStringBuffer()
//line niall.go:50
	_Cfunc_Niall_Reply(buf, size)
//line niall.go:52
	return _Cfunc_GoString(buf)
}
//line niall.go:57

//line niall.go:56
func Init() {
//line niall.go:56
	_Cfunc_Niall_Init()
//line niall.go:58
}
//line niall.go:62

//line niall.go:61
func Quit() {
//line niall.go:61
	_Cfunc_Niall_Free()
//line niall.go:63
}
//line niall.go:67

//line niall.go:66
func NewDictionary() {
//line niall.go:66
	_Cfunc_Niall_NewDictionary()
//line niall.go:68
}
//line niall.go:72

//line niall.go:71
func ListDictionary() {
//line niall.go:71
	_Cfunc_Niall_ListDictionary()
//line niall.go:73
}
//line niall.go:76

//line niall.go:75
func LoadDictionary(filename string) {
//line niall.go:79

//line niall.go:78
}
//line niall.go:81

//line niall.go:80
func SaveDictionary(filename string) {
//line niall.go:84

//line niall.go:83
}
//line niall.go:86

//line niall.go:85
func CorrectSpelling(original, correct string) {
//line niall.go:89

//line niall.go:88
}
