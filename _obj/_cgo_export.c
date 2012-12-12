/* Created by cgo - DO NOT EDIT. */
#include "_cgo_export.h"

extern void crosscall2(void (*fn)(void *, int), void *, int);

extern void _cgoexp_532b7e5247ed_niall_go_warning(void *, int);

void niall_go_warning(char* p0)
{
	struct {
		char* p0;
	} __attribute__((packed)) a;
	a.p0 = p0;
	crosscall2(_cgoexp_532b7e5247ed_niall_go_warning, &a, 8);
}
extern void _cgoexp_532b7e5247ed_niall_go_error(void *, int);

void niall_go_error(char* p0)
{
	struct {
		char* p0;
	} __attribute__((packed)) a;
	a.p0 = p0;
	crosscall2(_cgoexp_532b7e5247ed_niall_go_error, &a, 8);
}
extern void _cgoexp_532b7e5247ed_niall_go_print(void *, int);

void niall_go_print(char* p0)
{
	struct {
		char* p0;
	} __attribute__((packed)) a;
	a.p0 = p0;
	crosscall2(_cgoexp_532b7e5247ed_niall_go_print, &a, 8);
}
