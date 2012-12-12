
#include "runtime.h"
#include "cgocall.h"

void ·_Cerrno(void*, int32);

void
·_Cfunc_GoString(int8 *p, String s)
{
	s = runtime·gostring((byte*)p);
	FLUSH(&s);
}

void
·_Cfunc_GoStringN(int8 *p, int32 l, String s)
{
	s = runtime·gostringn((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_GoBytes(int8 *p, int32 l, Slice s)
{
	s = runtime·gobytes((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_CString(String s, int8 *p)
{
	p = runtime·cmalloc(s.len+1);
	runtime·memmove((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&p);
}

void _cgo_532b7e5247ed_Cfunc_Niall_Free(void*);

void
·_Cfunc_Niall_Free(struct{uint8 x[1];}p)
{
	runtime·cgocall(_cgo_532b7e5247ed_Cfunc_Niall_Free, &p);
}

void _cgo_532b7e5247ed_Cfunc_Niall_Init(void*);

void
·_Cfunc_Niall_Init(struct{uint8 x[1];}p)
{
	runtime·cgocall(_cgo_532b7e5247ed_Cfunc_Niall_Init, &p);
}

void _cgo_532b7e5247ed_Cfunc_Niall_Learn(void*);

void
·_Cfunc_Niall_Learn(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_532b7e5247ed_Cfunc_Niall_Learn, &p);
}

void _cgo_532b7e5247ed_Cfunc_Niall_ListDictionary(void*);

void
·_Cfunc_Niall_ListDictionary(struct{uint8 x[1];}p)
{
	runtime·cgocall(_cgo_532b7e5247ed_Cfunc_Niall_ListDictionary, &p);
}

void _cgo_532b7e5247ed_Cfunc_Niall_NewDictionary(void*);

void
·_Cfunc_Niall_NewDictionary(struct{uint8 x[1];}p)
{
	runtime·cgocall(_cgo_532b7e5247ed_Cfunc_Niall_NewDictionary, &p);
}

void _cgo_532b7e5247ed_Cfunc_Niall_Reply(void*);

void
·_Cfunc_Niall_Reply(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_532b7e5247ed_Cfunc_Niall_Reply, &p);
}

#pragma dynexport niall_go_warning niall_go_warning
extern void ·niall_go_warning();

#pragma textflag 7
void
_cgoexp_532b7e5247ed_niall_go_warning(void *a, int32 n)
{
	runtime·cgocallback(·niall_go_warning, a, n);
}
#pragma dynexport niall_go_error niall_go_error
extern void ·niall_go_error();

#pragma textflag 7
void
_cgoexp_532b7e5247ed_niall_go_error(void *a, int32 n)
{
	runtime·cgocallback(·niall_go_error, a, n);
}
#pragma dynexport niall_go_print niall_go_print
extern void ·niall_go_print();

#pragma textflag 7
void
_cgoexp_532b7e5247ed_niall_go_print(void *a, int32 n)
{
	runtime·cgocallback(·niall_go_print, a, n);
}
