package niall

/*
#include <stdio.h>
#include <stdarg.h>

// Startup/shutdown
void Niall_Init(void);
void Niall_Free(void);

// Niall's main functions
void Niall_Learn(char *Buffer);
void Niall_Reply(char *Buffer,int BufSize);

// Housekeeping functions
void Niall_NewDictionary(void);
void Niall_ListDictionary(void);
void Niall_SaveDictionary(char *file);
void Niall_LoadDictionary(char *file);
void Niall_CorrectSpelling(char *Original,char *Correct);

// Niall calls these functions (from C to Go)
extern void niall_go_print( char *msg );
extern void niall_go_warning( char *msg );
extern void niall_go_error( char *msg );
*/
import "C"

import (
	"bytes"
	"fmt"
)

var (
	WarningFunction = fmt.Print
	ErrorFunction = fmt.Print
	PrintFunction = fmt.Print
)

//export niall_go_warning
func niall_go_warning(msg *C.char) {
	WarningFunction("WARNING: " + C.GoString(msg))
}

//export niall_go_error
func niall_go_error(msg *C.char) {
	ErrorFunction("ERROR: " + C.GoString(msg))
}

//export niall_go_print
func niall_go_print(msg *C.char) {
	PrintFunction(C.GoString(msg))
}

// Create an empty *char buffer that is C.BUFSIZ large
// C.BUFSIZ is defined in stdio.h (was 1024 here)
func newCStringBuffer() (*C.char, _Ctype_int) {
	temp := make([]byte, 0, C.BUFSIZ)
	b := bytes.NewBuffer(temp)
	buf := C.CString(b.String())
	return buf, _Ctype_int(cap(temp))
}

// Make Niall learn from a string
func Learn(s string) {
	C.Niall_Learn(C.CString(s))
}

// Make Niall say something
func Talk() string {
	buf, size := newCStringBuffer()
	C.Niall_Reply(buf, size)
	return C.GoString(buf)
}

// This seeds the random number generator and clears the dictionary 
func Init() {
	C.Niall_Init()
}

// This frees the allocated memory
func Quit() {
	C.Niall_Free()
}

// Clear the dictionary
func NewDictionary() {
	C.Niall_NewDictionary()
}

// List the contents of the dictionary
func ListDictionary() {
	C.Niall_ListDictionary()
}

// Load the dictionary from file
func LoadDictionary(filename string) {
	C.Niall_LoadDictionary(C.CString(filename))
}

// Save the dictionary to file
func SaveDictionary(filename string) {
	C.Niall_SaveDictionary(C.CString(filename))
}

// Correct the spelling
func CorrectSpelling(original, correct string) {
	C.Niall_CorrectSpelling(C.CString(original), C.CString(correct))
}

