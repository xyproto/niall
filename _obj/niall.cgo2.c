#line 3 "niall.go"
#include <stdio.h>
#include <stdarg.h>
#include "niall.h"



// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention "name".
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

#include <errno.h>
#include <string.h>

void
_cgo_532b7e5247ed_Cfunc_Niall_Free(void *v)
{
	struct {
		char unused;
	} __attribute__((__packed__)) *a = v;
	Niall_Free();
}

void
_cgo_532b7e5247ed_Cfunc_Niall_Init(void *v)
{
	struct {
		char unused;
	} __attribute__((__packed__)) *a = v;
	Niall_Init();
}

void
_cgo_532b7e5247ed_Cfunc_Niall_Learn(void *v)
{
	struct {
		char* p0;
	} __attribute__((__packed__)) *a = v;
	Niall_Learn((void*)a->p0);
}

void
_cgo_532b7e5247ed_Cfunc_Niall_ListDictionary(void *v)
{
	struct {
		char unused;
	} __attribute__((__packed__)) *a = v;
	Niall_ListDictionary();
}

void
_cgo_532b7e5247ed_Cfunc_Niall_NewDictionary(void *v)
{
	struct {
		char unused;
	} __attribute__((__packed__)) *a = v;
	Niall_NewDictionary();
}

void
_cgo_532b7e5247ed_Cfunc_Niall_Reply(void *v)
{
	struct {
		char* p0;
		int p1;
		char __pad12[4];
	} __attribute__((__packed__)) *a = v;
	Niall_Reply((void*)a->p0, a->p1);
}

