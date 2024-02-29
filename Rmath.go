package Rmath
//developed by Alireza Alizadeh Aghdam
//website:https://www.rubikcomputer.ir
//AlirezaRubik
import (
	"fmt"
	"unsafe"
)
type RMath struct {
	
}
const haveArchTan = false
var _tanP = [...]float64{
	-1.30936939181383777646e4, // 0xc0c992d8d24f3f38
	1.15351664838587416140e6,  // 0x413199eca5fc9ddd
	-1.79565251976484877988e7, // 0xc1711fead3299176
}
var _tanQ = [...]float64{
	1.00000000000000000000e0,
	1.36812963470692954678e4,  // 0x40cab8a5eeb36572
	-1.32089234440210967447e6, // 0xc13427bc582abc96
	2.50083801823357915839e7,  // 0x4177d98fc2ead8ef
	-5.38695755929454629881e7, // 0xc189afe03cbe5a31
}
const haveArchLog1p = false

var mPi4 = [...]uint64{
	0x0000000000000001,
	0x45f306dc9c882a53,
	0xf84eafa3ea69bb81,
	0xb6c52b3278872083,
	0xfca2c757bd778ac3,
	0x6e48dc74849ba5c0,
	0x0c925dd413a32439,
	0xfc3bd63962534e7d,
	0xd1046bea5d768909,
	0xd338e04d68befc82,
	0x7323ac7306a673e9,
	0x3908bf177bf25076,
	0x3ff12fffbc0b301f,
	0xde5e2316b414da3e,
	0xda6cfd9e4f96136e,
	0x9e8c7ecd3cbfd45a,
	0xea4f758fd7cbe2f6,
	0x7a0e73ef14a525d4,
	0xd7f6bf623f1aba10,
	0xac06608df8f6d757,
}
var _sin = [...]float64{
	1.58962301576546568060e-10, // 0x3de5d8fd1fd19ccd
	-2.50507477628578072866e-8, // 0xbe5ae5e5a9291f5d
	2.75573136213857245213e-6,  // 0x3ec71de3567d48a1
	-1.98412698295895385996e-4, // 0xbf2a01a019bfdf03
	8.33333333332211858878e-3,  // 0x3f8111111110f7d0
	-1.66666666666666307295e-1, // 0xbfc5555555555548
}

// cos coefficients
var _cos = [...]float64{
	-1.13585365213876817300e-11, // 0xbda8fa49a0861a9b
	2.08757008419747316778e-9,   // 0x3e21ee9d7b4e3f05
	-2.75573141792967388112e-7,  // 0xbe927e4f7eac4bc6
	2.48015872888517045348e-5,   // 0x3efa01a019c844f5
	-1.38888888888730564116e-3,  // 0xbf56c16c16c14f91
	4.16666666666665929218e-2,   // 0x3fa555555555554b
}

const (
	haveArchAcos = false
	haveArchSin = false
	reduceThreshold = 1 << 29
	haveArchAsinh = false
	uvnan    = 0x7FF8000000000001
	uvinf    = 0x7FF0000000000000
	uvneginf = 0xFFF0000000000000
	uvone    = 0x3FF0000000000000
	mask     = 0x7FF
	shift    = 64 - 11 - 1
	bias     = 1023
	signMask = 1 << 63
	fracMask = 1<<shift - 1
	MaxFloat32             = 0x1p127 * (1 + (1 - 0x1p-23)) 
	SmallestNonzeroFloat32 = 0x1p-126 * 0x1p-23            
	E   = 2.71828182845904523536028747135266249775724709369995957496696763 
	Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 
	Phi = 1.61803398874989484820458683436563811772030917980576286213544862 
	Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339
	Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
	Log2E  = 1 / Ln2
	Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // https://oeis.org/A002392
	Log10E = 1 / Ln10
	MaxFloat64             = 0x1p1023 * (1 + (1 - 0x1p-52)) // 1.79769313486231570814527423731704356798070e+308
	SmallestNonzeroFloat64 = 0x1p-1022 * 0x1p-52            // 4.9406564584124654417656879286822137236505980e-324
)
const (
	Sqrt2M1     = 4.142135623730950488017e-01  // Sqrt(2)-1 = 0x3fda827999fcef34
	Sqrt2HalfM1 = -2.928932188134524755992e-01 // Sqrt(2)/2-1 = 0xbfd2bec333018866
	Small       = 1.0 / (1 << 29)              // 2**-29 = 0x3e20000000000000
	Tiny        = 1.0 / (1 << 54)              // 2**-54
	Two53       = 1 << 53                      // 2**53
	Ln2Hi       = 6.93147180369123816490e-01   // 3fe62e42fee00000
	Ln2Lo       = 1.90821492927058770002e-10   // 3dea39ef35793c76
	Lp1         = 6.666666666666735130e-01     // 3FE5555555555593
	Lp2         = 3.999999999940941908e-01     // 3FD999999997FA04
	Lp3         = 2.857142874366239149e-01     // 3FD2492494229359
	Lp4         = 2.222219843214978396e-01     // 3FCC71C51D8E78AF
	Lp5         = 1.818357216161805012e-01     // 3FC7466496CB03DE
	Lp6         = 1.531383769920937332e-01     // 3FC39A09D078C69F
	Lp7         = 1.479819860511658591e-01     // 3FC2F112DF3E5244
)
const m0 = 0x5555555555555555 // 01010101 ...
const m1 = 0x3333333333333333 // 00110011 ...
const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...
const m3 = 0x00ff00ff00ff00ff // etc.
const m4 = 0x0000ffff0000ffff
const ntz8tab = "" +
	"\x08\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x05\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x06\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x05\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x07\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x05\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x06\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x05\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00"

const pop8tab = "" +
	"\x00\x01\x01\x02\x01\x02\x02\x03\x01\x02\x02\x03\x02\x03\x03\x04" +
	"\x01\x02\x02\x03\x02\x03\x03\x04\x02\x03\x03\x04\x03\x04\x04\x05" +
	"\x01\x02\x02\x03\x02\x03\x03\x04\x02\x03\x03\x04\x03\x04\x04\x05" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x01\x02\x02\x03\x02\x03\x03\x04\x02\x03\x03\x04\x03\x04\x04\x05" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x03\x04\x04\x05\x04\x05\x05\x06\x04\x05\x05\x06\x05\x06\x06\x07" +
	"\x01\x02\x02\x03\x02\x03\x03\x04\x02\x03\x03\x04\x03\x04\x04\x05" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x03\x04\x04\x05\x04\x05\x05\x06\x04\x05\x05\x06\x05\x06\x06\x07" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x03\x04\x04\x05\x04\x05\x05\x06\x04\x05\x05\x06\x05\x06\x06\x07" +
	"\x03\x04\x04\x05\x04\x05\x05\x06\x04\x05\x05\x06\x05\x06\x06\x07" +
	"\x04\x05\x05\x06\x05\x06\x06\x07\x05\x06\x06\x07\x06\x07\x07\x08"

const rev8tab = "" +
	"\x00\x80\x40\xc0\x20\xa0\x60\xe0\x10\x90\x50\xd0\x30\xb0\x70\xf0" +
	"\x08\x88\x48\xc8\x28\xa8\x68\xe8\x18\x98\x58\xd8\x38\xb8\x78\xf8" +
	"\x04\x84\x44\xc4\x24\xa4\x64\xe4\x14\x94\x54\xd4\x34\xb4\x74\xf4" +
	"\x0c\x8c\x4c\xcc\x2c\xac\x6c\xec\x1c\x9c\x5c\xdc\x3c\xbc\x7c\xfc" +
	"\x02\x82\x42\xc2\x22\xa2\x62\xe2\x12\x92\x52\xd2\x32\xb2\x72\xf2" +
	"\x0a\x8a\x4a\xca\x2a\xaa\x6a\xea\x1a\x9a\x5a\xda\x3a\xba\x7a\xfa" +
	"\x06\x86\x46\xc6\x26\xa6\x66\xe6\x16\x96\x56\xd6\x36\xb6\x76\xf6" +
	"\x0e\x8e\x4e\xce\x2e\xae\x6e\xee\x1e\x9e\x5e\xde\x3e\xbe\x7e\xfe" +
	"\x01\x81\x41\xc1\x21\xa1\x61\xe1\x11\x91\x51\xd1\x31\xb1\x71\xf1" +
	"\x09\x89\x49\xc9\x29\xa9\x69\xe9\x19\x99\x59\xd9\x39\xb9\x79\xf9" +
	"\x05\x85\x45\xc5\x25\xa5\x65\xe5\x15\x95\x55\xd5\x35\xb5\x75\xf5" +
	"\x0d\x8d\x4d\xcd\x2d\xad\x6d\xed\x1d\x9d\x5d\xdd\x3d\xbd\x7d\xfd" +
	"\x03\x83\x43\xc3\x23\xa3\x63\xe3\x13\x93\x53\xd3\x33\xb3\x73\xf3" +
	"\x0b\x8b\x4b\xcb\x2b\xab\x6b\xeb\x1b\x9b\x5b\xdb\x3b\xbb\x7b\xfb" +
	"\x07\x87\x47\xc7\x27\xa7\x67\xe7\x17\x97\x57\xd7\x37\xb7\x77\xf7" +
	"\x0f\x8f\x4f\xcf\x2f\xaf\x6f\xef\x1f\x9f\x5f\xdf\x3f\xbf\x7f\xff"

const len8tab = "" +
	"\x00\x01\x02\x02\x03\x03\x03\x03\x04\x04\x04\x04\x04\x04\x04\x04" +
	"\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05" +
	"\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06" +
	"\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08"

const uintSize = 32 << (^uint(0) >> 63) // 32 or 64
// ArbitraryType is here for the purposes of documentation only and is not actually
// part of the unsafe package. It represents the type of an arbitrary Go expression.
type ArbitraryType int

// IntegerType is here for the purposes of documentation only and is not actually
// part of the unsafe package. It represents any arbitrary integer type.
type IntegerType int

// Pointer represents a pointer to an arbitrary type. There are four special operations
// available for type Pointer that are not available for other types:
//   - A pointer value of any type can be converted to a Pointer.
//   - A Pointer can be converted to a pointer value of any type.
//   - A uintptr can be converted to a Pointer.
//   - A Pointer can be converted to a uintptr.
//
// Pointer therefore allows a program to defeat the type system and read and write
// arbitrary memory. It should be used with extreme care.
//
// The following patterns involving Pointer are valid.
// Code not using these patterns is likely to be invalid today
// or to become invalid in the future.
// Even the valid patterns below come with important caveats.
//
// Running "go vet" can help find uses of Pointer that do not conform to these patterns,
// but silence from "go vet" is not a guarantee that the code is valid.
//
// (1) Conversion of a *T1 to Pointer to *T2.
//
// Provided that T2 is no larger than T1 and that the two share an equivalent
// memory layout, this conversion allows reinterpreting data of one type as
// data of another type. An example is the implementation of
// math.Float64bits:
//
//	func Float64bits(f float64) uint64 {
//		return *(*uint64)(unsafe.Pointer(&f))
//	}
//
// (2) Conversion of a Pointer to a uintptr (but not back to Pointer).
//
// Converting a Pointer to a uintptr produces the memory address of the value
// pointed at, as an integer. The usual use for such a uintptr is to print it.
//
// Conversion of a uintptr back to Pointer is not valid in general.
//
// A uintptr is an integer, not a reference.
// Converting a Pointer to a uintptr creates an integer value
// with no pointer semantics.
// Even if a uintptr holds the address of some object,
// the garbage collector will not update that uintptr's value
// if the object moves, nor will that uintptr keep the object
// from being reclaimed.
//
// The remaining patterns enumerate the only valid conversions
// from uintptr to Pointer.
//
// (3) Conversion of a Pointer to a uintptr and back, with arithmetic.
//
// If p points into an allocated object, it can be advanced through the object
// by conversion to uintptr, addition of an offset, and conversion back to Pointer.
//
//	p = unsafe.Pointer(uintptr(p) + offset)
//
// The most common use of this pattern is to access fields in a struct
// or elements of an array:
//
//	// equivalent to f := unsafe.Pointer(&s.f)
//	f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))
//
//	// equivalent to e := unsafe.Pointer(&x[i])
//	e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))
//
// It is valid both to add and to subtract offsets from a pointer in this way.
// It is also valid to use &^ to round pointers, usually for alignment.
// In all cases, the result must continue to point into the original allocated object.
//
// Unlike in C, it is not valid to advance a pointer just beyond the end of
// its original allocation:
//
//	// INVALID: end points outside allocated space.
//	var s thing
//	end = unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s))
//
//	// INVALID: end points outside allocated space.
//	b := make([]byte, n)
//	end = unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n))
//
// Note that both conversions must appear in the same expression, with only
// the intervening arithmetic between them:
//
//	// INVALID: uintptr cannot be stored in variable
//	// before conversion back to Pointer.
//	u := uintptr(p)
//	p = unsafe.Pointer(u + offset)
//
// Note that the pointer must point into an allocated object, so it may not be nil.
//
//	// INVALID: conversion of nil pointer
//	u := unsafe.Pointer(nil)
//	p := unsafe.Pointer(uintptr(u) + offset)
//
// (4) Conversion of a Pointer to a uintptr when calling syscall.Syscall.
//
// The Syscall functions in package syscall pass their uintptr arguments directly
// to the operating system, which then may, depending on the details of the call,
// reinterpret some of them as pointers.
// That is, the system call implementation is implicitly converting certain arguments
// back from uintptr to pointer.
//
// If a pointer argument must be converted to uintptr for use as an argument,
// that conversion must appear in the call expression itself:
//
//	syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))
//
// The compiler handles a Pointer converted to a uintptr in the argument list of
// a call to a function implemented in assembly by arranging that the referenced
// allocated object, if any, is retained and not moved until the call completes,
// even though from the types alone it would appear that the object is no longer
// needed during the call.
//
// For the compiler to recognize this pattern,
// the conversion must appear in the argument list:
//
//	// INVALID: uintptr cannot be stored in variable
//	// before implicit conversion back to Pointer during system call.
//	u := uintptr(unsafe.Pointer(p))
//	syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n))
//
// (5) Conversion of the result of reflect.Value.Pointer or reflect.Value.UnsafeAddr
// from uintptr to Pointer.
//
// Package reflect's Value methods named Pointer and UnsafeAddr return type uintptr
// instead of unsafe.Pointer to keep callers from changing the result to an arbitrary
// type without first importing "unsafe". However, this means that the result is
// fragile and must be converted to Pointer immediately after making the call,
// in the same expression:
//
//	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
//
// As in the cases above, it is invalid to store the result before the conversion:
//
//	// INVALID: uintptr cannot be stored in variable
//	// before conversion back to Pointer.
//	u := reflect.ValueOf(new(int)).Pointer()
//	p := (*int)(unsafe.Pointer(u))
//
// (6) Conversion of a reflect.SliceHeader or reflect.StringHeader Data field to or from Pointer.
//
// As in the previous case, the reflect data structures SliceHeader and StringHeader
// declare the field Data as a uintptr to keep callers from changing the result to
// an arbitrary type without first importing "unsafe". However, this means that
// SliceHeader and StringHeader are only valid when interpreting the content
// of an actual slice or string value.
//
//	var s string
//	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
//	hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
//	hdr.Len = n
//
// In this usage hdr.Data is really an alternate way to refer to the underlying
// pointer in the string header, not a uintptr variable itself.
//
// In general, reflect.SliceHeader and reflect.StringHeader should be used
// only as *reflect.SliceHeader and *reflect.StringHeader pointing at actual
// slices or strings, never as plain structs.
// A program should not declare or allocate variables of these struct types.
//
//	// INVALID: a directly-declared header will not hold Data as a reference.
//	var hdr reflect.StringHeader
//	hdr.Data = uintptr(unsafe.Pointer(p))
//	hdr.Len = n
//	s := *(*string)(unsafe.Pointer(&hdr)) // p possibly already lost
type Pointer *ArbitraryType

// Sizeof takes an expression x of any type and returns the size in bytes
// of a hypothetical variable v as if v was declared via var v = x.
// The size does not include any memory possibly referenced by x.
// For instance, if x is a slice, Sizeof returns the size of the slice
// descriptor, not the size of the memory referenced by the slice.
// For a struct, the size includes any padding introduced by field alignment.
// The return value of Sizeof is a Go constant if the type of the argument x
// does not have variable size.
// (A type has variable size if it is a type parameter or if it is an array
// or struct type with elements of variable size).
func Sizeof(x ArbitraryType) uintptr

// Offsetof returns the offset within the struct of the field represented by x,
// which must be of the form structValue.field. In other words, it returns the
// number of bytes between the start of the struct and the start of the field.
// The return value of Offsetof is a Go constant if the type of the argument x
// does not have variable size.
// (See the description of [Sizeof] for a definition of variable sized types.)
func Offsetof(x ArbitraryType) uintptr

// Alignof takes an expression x of any type and returns the required alignment
// of a hypothetical variable v as if v was declared via var v = x.
// It is the largest value m such that the address of v is always zero mod m.
// It is the same as the value returned by reflect.TypeOf(x).Align().
// As a special case, if a variable s is of struct type and f is a field
// within that struct, then Alignof(s.f) will return the required alignment
// of a field of that type within a struct. This case is the same as the
// value returned by reflect.TypeOf(s.f).FieldAlign().
// The return value of Alignof is a Go constant if the type of the argument
// does not have variable size.
// (See the description of [Sizeof] for a definition of variable sized types.)
func Alignof(x ArbitraryType) uintptr

// The function Add adds len to ptr and returns the updated pointer
// Pointer(uintptr(ptr) + uintptr(len)).
// The len argument must be of integer type or an untyped constant.
// A constant len argument must be representable by a value of type int;
// if it is an untyped constant it is given type int.
// The rules for valid uses of Pointer still apply.
func Add(ptr Pointer, len IntegerType) Pointer

// The function Slice returns a slice whose underlying array starts at ptr
// and whose length and capacity are len.
// Slice(ptr, len) is equivalent to
//
//	(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
//
// except that, as a special case, if ptr is nil and len is zero,
// Slice returns nil.
//
// The len argument must be of integer type or an untyped constant.
// A constant len argument must be non-negative and representable by a value of type int;
// if it is an untyped constant it is given type int.
// At run time, if len is negative, or if ptr is nil and len is not zero,
// a run-time panic occurs.
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType

// SliceData returns a pointer to the underlying array of the argument
// slice.
//   - If cap(slice) > 0, SliceData returns &slice[:1][0].
//   - If slice == nil, SliceData returns nil.
//   - Otherwise, SliceData returns a non-nil pointer to an
//     unspecified memory address.
func SliceData(slice []ArbitraryType) *ArbitraryType

// String returns a string value whose underlying bytes
// start at ptr and whose length is len.
//
// The len argument must be of integer type or an untyped constant.
// A constant len argument must be non-negative and representable by a value of type int;
// if it is an untyped constant it is given type int.
// At run time, if len is negative, or if ptr is nil and len is not zero,
// a run-time panic occurs.
//
// Since Go strings are immutable, the bytes passed to String
// must not be modified afterwards.
func String(ptr *byte, len IntegerType) string

// StringData returns a pointer to the underlying bytes of str.
// For an empty string the return value is unspecified, and may be nil.
//
// Since Go strings are immutable, the bytes returned by StringData
// must not be modified.
func StringData(str string) *byte

// UintSize is the size of a uint in bits.
const UintSize = uintSize

// --- LeadingZeros ---

// LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
func (math *RMath) LeadingZeros(x uint) int { return UintSize - math.Len(x) }

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func (math *RMath)LeadingZeros8(x uint8) int { return 8 - math.Len8(x) }

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func (math *RMath)LeadingZeros16(x uint16) int { return 16 - math.Len16(x) }

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func (math *RMath) LeadingZeros32(x uint32) int { return 32 - math.Len32(x) }

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func   (math *RMath)LeadingZeros64(x uint64) int { return 64 - math.Len64(x) }

// --- TrailingZeros ---

// See http://supertech.csail.mit.edu/papers/debruijn.pdf
const deBruijn32 = 0x077CB531

var deBruijn32tab = [32]byte{
	0, 1, 28, 2, 29, 14, 24, 3, 30, 22, 20, 15, 25, 17, 4, 8,
	31, 27, 13, 23, 21, 19, 16, 7, 26, 12, 18, 6, 11, 5, 10, 9,
}

const deBruijn64 = 0x03f79d71b4ca8b09

var deBruijn64tab = [64]byte{
	0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
	62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
	63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
	54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
}

// TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.
func  (math *RMath) TrailingZeros(x uint) int {
	if UintSize == 32 {
		return math.TrailingZeros32(uint32(x))
	}
	return math.TrailingZeros64(uint64(x))
}

// TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
func   (math *RMath) TrailingZeros8(x uint8) int {
	return int(ntz8tab[x])
}

// TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.
func  (math *RMath) TrailingZeros16(x uint16) int {
	if x == 0 {
		return 16
	}
	// see comment in TrailingZeros64
	return int(deBruijn32tab[uint32(x&-x)*deBruijn32>>(32-5)])
}

// TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
func  (math *RMath) TrailingZeros32(x uint32) int {
	if x == 0 {
		return 32
	}
	// see comment in TrailingZeros64
	return int(deBruijn32tab[(x&-x)*deBruijn32>>(32-5)])
}

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
func  (math *RMath) TrailingZeros64(x uint64) int {
	if x == 0 {
		return 64
	}
	return int(deBruijn64tab[(x&-x)*deBruijn64>>(64-6)])
}

// --- OnesCount ---



// OnesCount returns the number of one bits ("population count") in x.
func (math *RMath)OnesCount(x uint) int {
	if UintSize == 32 {
		return math.OnesCount32(uint32(x))
	}
	return math.OnesCount64(uint64(x))
}

// OnesCount8 returns the number of one bits ("population count") in x.
func (math *RMath)OnesCount8(x uint8) int {
	return int(pop8tab[x])
}

// OnesCount16 returns the number of one bits ("population count") in x.
func (math *RMath)OnesCount16(x uint16) int {
	return int(pop8tab[x>>8] + pop8tab[x&0xff])
}

// OnesCount32 returns the number of one bits ("population count") in x.
func (math *RMath) OnesCount32(x uint32) int {
	return int(pop8tab[x>>24] + pop8tab[x>>16&0xff] + pop8tab[x>>8&0xff] + pop8tab[x&0xff])
}

// OnesCount64 returns the number of one bits ("population count") in x.
func (math *RMath) OnesCount64(x uint64) int {
	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}

// --- RotateLeft ---

// RotateLeft returns the value of x rotated left by (k mod UintSize) bits.
// To rotate x right by k bits, call RotateLeft(x, -k).
// This function's execution time does not depend on the inputs.
func (math *RMath)RotateLeft(x uint, k int) uint {
	if UintSize == 32 {
		return uint(math.RotateLeft32(uint32(x), k))
	}
	return uint(math.RotateLeft64(uint64(x), k))
}

// RotateLeft8 returns the value of x rotated left by (k mod 8) bits.
// To rotate x right by k bits, call RotateLeft8(x, -k).
// This function's execution time does not depend on the inputs.
func (math *RMath) RotateLeft8(x uint8, k int) uint8 {
	const n = 8
	s := uint(k) & (n - 1)
	return x<<s | x>>(n-s)
}

// RotateLeft16 returns the value of x rotated left by (k mod 16) bits.
// To rotate x right by k bits, call RotateLeft16(x, -k).
// This function's execution time does not depend on the inputs.
func (math *RMath) RotateLeft16(x uint16, k int) uint16 {
	const n = 16
	s := uint(k) & (n - 1)
	return x<<s | x>>(n-s)
}

// RotateLeft32 returns the value of x rotated left by (k mod 32) bits.
// To rotate x right by k bits, call RotateLeft32(x, -k).
// This function's execution time does not depend on the inputs.
func (math *RMath)RotateLeft32(x uint32, k int) uint32 {
	const n = 32
	s := uint(k) & (n - 1)
	return x<<s | x>>(n-s)
}

// RotateLeft64 returns the value of x rotated left by (k mod 64) bits.
// To rotate x right by k bits, call RotateLeft64(x, -k).
// This function's execution time does not depend on the inputs.
func (math *RMath) RotateLeft64(x uint64, k int) uint64 {
	const n = 64
	s := uint(k) & (n - 1)
	return x<<s | x>>(n-s)
}

// --- Reverse ---
// Reverse returns the value of x with its bits in reversed order.
func (math *RMath)Reverse(x uint) uint {
	if UintSize == 32 {
		return uint(math.Reverse32(uint32(x)))
	}
	return uint(math.Reverse64(uint64(x)))
}

// Reverse8 returns the value of x with its bits in reversed order.
func (math *RMath) Reverse8(x uint8) uint8 {
	return rev8tab[x]
}
// Reverse16 returns the value of x with its bits in reversed order.
func (math *RMath)Reverse16(x uint16) uint16 {
	return uint16(rev8tab[x>>8]) | uint16(rev8tab[x&0xff])<<8
}

// Reverse32 returns the value of x with its bits in reversed order.
func (math *RMath)Reverse32(x uint32) uint32 {
	const m = 1<<32 - 1
	x = x>>1&(m0&m) | x&(m0&m)<<1
	x = x>>2&(m1&m) | x&(m1&m)<<2
	x = x>>4&(m2&m) | x&(m2&m)<<4
	return math.ReverseBytes32(x)
}

// Reverse64 returns the value of x with its bits in reversed order.
func (math *RMath)Reverse64(x uint64) uint64 {
	const m = 1<<64 - 1
	x = x>>1&(m0&m) | x&(m0&m)<<1
	x = x>>2&(m1&m) | x&(m1&m)<<2
	x = x>>4&(m2&m) | x&(m2&m)<<4
	return math.ReverseBytes64(x)
}

// --- ReverseBytes ---

// ReverseBytes returns the value of x with its bytes in reversed order.
//
// This function's execution time does not depend on the inputs.
func (math *RMath) ReverseBytes(x uint) uint {
	if UintSize == 32 {
		return uint(math.ReverseBytes32(uint32(x)))
	}
	return uint(math.ReverseBytes64(uint64(x)))
}

// ReverseBytes16 returns the value of x with its bytes in reversed order.
//
// This function's execution time does not depend on the inputs.
func (math *RMath) ReverseBytes16(x uint16) uint16 {
	return x>>8 | x<<8
}

// ReverseBytes32 returns the value of x with its bytes in reversed order.
//
// This function's execution time does not depend on the inputs.
func (math *RMath) ReverseBytes32(x uint32) uint32 {
	const m = 1<<32 - 1
	x = x>>8&(m3&m) | x&(m3&m)<<8
	return x>>16 | x<<16
}

// ReverseBytes64 returns the value of x with its bytes in reversed order.
//
// This function's execution time does not depend on the inputs.
func (math *RMath) ReverseBytes64(x uint64) uint64 {
	const m = 1<<64 - 1
	x = x>>8&(m3&m) | x&(m3&m)<<8
	x = x>>16&(m4&m) | x&(m4&m)<<16
	return x>>32 | x<<32
}

// --- Len ---

// Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func (math *RMath) Len(x uint) int {
	if UintSize == 32 {
		return math.Len32(uint32(x))
	}
	return math.Len64(uint64(x))
}

// Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func (math *RMath) Len8(x uint8) int {
	return int(len8tab[x])
}

// Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func (math *RMath) Len16(x uint16) (n int) {
	if x >= 1<<8 {
		x >>= 8
		n = 8
	}
	return n + int(len8tab[x])
}

// Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func (math *RMath) Len32(x uint32) (n int) {
	if x >= 1<<16 {
		x >>= 16
		n = 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return n + int(len8tab[x])
}

// Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func   (math *RMath) Len64(x uint64) (n int) {
	if x >= 1<<32 {
		x >>= 32
		n = 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return n + int(len8tab[x])
}
// Add returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func (math *RMath) Add(x, y, carry uint) (sum, carryOut uint) {
	if UintSize == 32 {
		s32, c32 := math.Add32(uint32(x), uint32(y), uint32(carry))
		return uint(s32), uint(c32)
	}
	s64, c64 := math. Add64(uint64(x), uint64(y), uint64(carry))
	return uint(s64), uint(c64)
}

// Add32 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
// This function's execution time does not depend on the inputs.
func (math *RMath)Add32(x, y, carry uint32) (sum, carryOut uint32) {
	sum64 := uint64(x) + uint64(y) + uint64(carry)
	sum = uint32(sum64)
	carryOut = uint32(sum64 >> 32)
	return
}

// Add64 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
// This function's execution time does not depend on the inputs.
func (math *RMath)Add64(x, y, carry uint64) (sum, carryOut uint64) {
	sum = x + y + carry
	// The sum will overflow if both top bits are set (x & y) or if one of them
	// is (x | y), and a carry from the lower place happened. If such a carry
	// happens, the top bit will be 1 + 0 + 1 = 0 (&^ sum).
	carryOut = ((x & y) | ((x | y) &^ sum)) >> 63
	return
}

// --- Subtract with borrow ---

// Sub returns the difference of x, y and borrow: diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
// This function's execution time does not depend on the inputs.
func (math *RMath) Sub(x, y, borrow uint) (diff, borrowOut uint) {
	if UintSize == 32 {
		d32, b32 := math.Sub32(uint32(x), uint32(y), uint32(borrow))
		return uint(d32), uint(b32)
	}
	d64, b64 := math.Sub64(uint64(x), uint64(y), uint64(borrow))
	return uint(d64), uint(b64)
}
// Sub32 returns the difference of x, y and borrow, diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func (math *RMath)Sub32(x, y, borrow uint32) (diff, borrowOut uint32) {
	diff = x - y - borrow
	// The difference will underflow if the top bit of x is not set and the top
	// bit of y is set (^x & y) or if they are the same (^(x ^ y)) and a borrow
	// from the lower place happens. If that borrow happens, the result will be
	// 1 - 1 - 1 = 0 - 0 - 1 = 1 (& diff).
	borrowOut = ((^x & y) | (^(x ^ y) & diff)) >> 31
	return
}
// Sub64 returns the difference of x, y and borrow: diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
// This function's execution time does not depend on the inputs.
func (math *RMath) Sub64(x, y, borrow uint64) (diff, borrowOut uint64) {
	diff = x - y - borrow
	// See Sub32 for the bit logic.
	borrowOut = ((^x & y) | (^(x ^ y) & diff)) >> 63
	return
}

// Mul returns the full-width product of x and y: (hi, lo) = x * y
// with the product bits' upper half returned in hi and the lower
// half returned in lo.
// This function's execution time does not depend on the inputs.
func (math *RMath) Mul(x, y uint) (hi, lo uint) {
	if UintSize == 32 {
		h, l := Mul32(uint32(x), uint32(y))
		return uint(h), uint(l)
	}
	h, l := math.Mul64(uint64(x), uint64(y))
	return uint(h), uint(l)
}

func Mul32(x, y uint32) (hi, lo uint32) {
	tmp := uint64(x) * uint64(y)
	hi, lo = uint32(tmp>>32), uint32(tmp)
	return
}

func (math *RMath) Div(hi, lo, y uint) (quo, rem uint) {
	if UintSize == 32 {
		q, r := math.Div32(uint32(hi), uint32(lo), uint32(y))
		return uint(q), uint(r)
	}
	q, r := math.Div64(uint64(hi), uint64(lo), uint64(y))
	return uint(q), uint(r)
}

// Div32 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
// half in parameter hi and the lower half in parameter lo.
// Div32 panics for y == 0 (division by zero) or y <= hi (quotient overflow).
func (math *RMath) Div32(hi, lo, y uint32) (quo, rem uint32) {
	if y != 0 && y <= hi {
		panic(overflowError)
	}
	z := uint64(hi)<<32 | uint64(lo)
	quo, rem = uint32(z/uint64(y)), uint32(z%uint64(y))
	return
}

// Div64 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
// half in parameter hi and the lower half in parameter lo.
// Div64 panics for y == 0 (division by zero) or y <= hi (quotient overflow).
func (math *RMath) Div64(hi, lo, y uint64) (quo, rem uint64) {
	if y == 0 {
		panic(divideError)
	}
	if y <= hi {
		panic(overflowError)
	}
	// If high part is zero, we can directly return the results.
	if hi == 0 {
		return lo / y, lo % y
	}

	s := uint(math.LeadingZeros64(y))
	y <<= s

	const (
		two32  = 1 << 32
		mask32 = two32 - 1
	)
	yn1 := y >> 32
	yn0 := y & mask32
	un32 := hi<<s | lo>>(64-s)
	un10 := lo << s
	un1 := un10 >> 32
	un0 := un10 & mask32
	q1 := un32 / yn1
	rhat := un32 - q1*yn1

	for q1 >= two32 || q1*yn0 > two32*rhat+un1 {
		q1--
		rhat += yn1
		if rhat >= two32 {
			break
		}
	}

	un21 := un32*two32 + un1 - q1*y
	q0 := un21 / yn1
	rhat = un21 - q0*yn1

	for q0 >= two32 || q0*yn0 > two32*rhat+un0 {
		q0--
		rhat += yn1
		if rhat >= two32 {
			break
		}
	}

	return q1*two32 + q0, (un21*two32 + un0 - q0*y) >> s
}

// Rem returns the remainder of (hi, lo) divided by y. Rem panics for
// y == 0 (division by zero) but, unlike Div, it doesn't panic on a
// quotient overflow.
func (math *RMath) Rem(hi, lo, y uint) uint {
	if UintSize == 32 {
		return uint(math.Rem32(uint32(hi), uint32(lo), uint32(y)))
	}
	return uint(math.Rem64(uint64(hi), uint64(lo), uint64(y)))
}

// Rem32 returns the remainder of (hi, lo) divided by y. Rem32 panics
// for y == 0 (division by zero) but, unlike Div32, it doesn't panic
// on a quotient overflow.
func (math *RMath) Rem32(hi, lo, y uint32) uint32 {
	return uint32((uint64(hi)<<32 | uint64(lo)) % uint64(y))
}

// Rem64 returns the remainder of (hi, lo) divided by y. Rem64 panics
// for y == 0 (division by zero) but, unlike Div64, it doesn't panic
// on a quotient overflow.
func (math *RMath) Rem64(hi, lo, y uint64) uint64 {
	// We scale down hi so that hi < y, then use Div64 to compute the
	// rem with the guarantee that it won't panic on quotient overflow.
	// Given that
	//   hi ≡ hi%y    (mod y)
	// we have
	//   hi<<64 + lo ≡ (hi%y)<<64 + lo    (mod y)
	_, rem := math.Div64(hi%y, lo, y)
	return rem
}

//go:linkname overflowError runtime.overflowError
var overflowError error

//go:linkname divideError runtime.divideError
var divideError error



// find even or odd
func (math *RMath) IsEvenOdd(input int) bool {
	// 1-true->even
	//0-false->odd
	if input%2 == 0 {
		return true
	} else {
		return false
	}
}

// find fibo
func (math *RMath) Fibo(input int) {
	A := 0
	B := 1
	var C int
	if input == 1 {
		fmt.Println("0")
		return
	}
	if input == 2 {
		fmt.Println("0,1")
		return
	}
	fmt.Print("0,1")
	for i := 3; i <= input; i++ {
		C = A + B
		fmt.Print(",", C)
		A = B
		B = C
	}
	return
}

// multi sum for int numbers
func (math *RMath) Isumer(input ...int) int {
	var sum int = 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}

// multi sum for float numbers
func (math *RMath) Fsumer(input ...float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}

// multi minus  int numbers
func (math *RMath) Iminus(input ...int) int {
	var minus int = 0
	for i := 0; i < len(input); i++ {
		minus -= input[i]
	}
	return minus
}

// multi minus for float numbers
func (math *RMath) Fminus(input ...float64) float64 {
	var minus float64 = 0
	for i := 0; i < len(input); i++ {
		minus -= input[i]
	}
	return minus
}

// multi Multiplication  for float numbers
func (math *RMath) Fmulti(input ...float64) float64 {
	var Multiplication float64 = 1
	for i := 0; i < len(input); i++ {
		Multiplication *= input[i]
	}
	return Multiplication
}

// multi Multiplication  for Int numbers
func (math *RMath) Imulti(input ...int) int {
	var Multiplication int = 1
	for i := 0; i < len(input); i++ {
		Multiplication *= input[i]
	}
	return Multiplication
}

// find prime number
func (math *RMath) IsPrime(input int) bool {
	//1- prime
	//2-not prime
	var count = 0
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			count++
		}
	}
	if count <= 2 {
		return true
	} else {
		return false
	}
}

// pow
func (math *RMath) Pow(base float64, exponent int) float64 {
	result := 1.0
	if exponent > 0 {
		for i := 0; i < exponent; i++ {
			result *= base
		}
	} else if exponent < 0 {
		for i := 0; i > exponent; i-- {
			result /= base
		}
	}
	return result
}

// fact
func (math *RMath) Fact(input int) int {
	var result = 1
	for i := 1; i <= input; i++ {
		result *= i
	}
	return result
}

// mod in int
func (math *RMath) Imod(num int, num2 int) int {
	return num % num2
}
func (math *RMath)Float64frombits(b uint64) float64 { return *(*float64)(unsafe.Pointer(&b)) }
func (math *RMath)Float64bits(f float64) uint64     { return *(*uint64)(unsafe.Pointer(&f)) }

// Abs
func (math *RMath) Abs(input float64) float64 {
	return math.Float64frombits(math.Float64bits(input) &^ (1 << 63))
}

// sqrt
func (math *RMath) Sqrt(number float64) float64 {
	const accuracy = 0.00001
	guess := number / 2
	for i := 0; i < 1000; i++ {
		guess = 0.5 * (guess + number/guess)
		diff := guess*guess - number
		if diff < 0 {
			diff = -diff
		}
		if diff < accuracy {
			break
		}
	}
	return guess
}

// find cqrt
func (math *RMath) Cqrt(input float64) float64 {
	const accuracy = 0.00001
	if input == 0 {
		return 0
	}
	guess := 1.0
	for i := 0; i < 1000; i++ {
		newGuess := (2*guess*guess*guess + input) / (3 * guess * guess)
		diff := newGuess - guess
		if diff < 0 {
			diff = -diff
		}
		if diff < accuracy {
			break
		}
		guess = newGuess
	}
	return guess
}

// log
func (math *RMath) Ln(input float64) float64 {
	if input <= 0 {
		return -1
	}
	//if your input wrong ->output=-1
	accuracy := 0.00000001
	result := 0.0
	term := (input - 1) / (input + 1)
	power := term
	n := 1

	for power > accuracy {
		result += power / float64(n)
		power *= term * term
		n += 2
	}

	return 2 * result
}

// log 10
func (math *RMath) Log10(input float64) float64 {
	return math.Ln(input) / math.Ln(10)
}

// log 2
func (math *RMath) Log2(input float64) float64 {
	return math.Ln(input) / math.Ln(2)
}

// Ceil
func (math *RMath) Ceil(input float64) int {
	if input >= 0 {
		return int(input + 0.999999999)
	}
	return int(input)
}

// floor
func (math *RMath) Floor(input float64) int {
	if input >= 0 {
		return int(input)
	}
	return int(input - 0.999999999)
}

// big number in slice for float Numbers
// float-slice-big-number
func (math *RMath) FSBigNum(input []float64) float64 {
	num := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > num {
			num = input[i]
		}
	}
	return num
}

// big number in slice for int Numbers
// int-slice-big-number
func (math *RMath) ISBigNum(input []int) int {
	num := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > num {
			num = input[i]
		}
	}
	return num
}

// little number in slice for float Numbers
// float -slice-little-number
func (math *RMath) FSLittleNum(input []float64) float64 {
	num := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < num {
			num = input[i]
		}
	}
	return num
}

// little number in slice for int Numbers
// int -slice-little-number
func (math *RMath) ISLittleNum(input []int) int {
	num := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < num {
			num = input[i]
		}
	}
	return num
}

// reverse of mum + sum + count of nums +avarage
func (math *RMath) Reverser(input int) (sum int, count int, ave float64) {
	var reverse int
	sum = 0
	for i := 0; input > 0; i++ {
		count++
		reverse = input % 10
		fmt.Print(reverse)
		sum += reverse
		input /= 10
	}
	fmt.Printf("\n")
	return sum, count, (float64(sum) / float64(count))
}

// Round
func (math *RMath) Round(input float64) int {
	if input >= 0 {
		return int(input + 0.5)
	}
	return int(input - 0.5)
}
func (math *RMath) FindSlice(input interface{}, slice []interface{}) (bool, int, interface{}) {
	for i := 0; i < len(slice); i++ {
		if input == slice[i] {
			return true, i, slice[i]
		}
	}

	return false, 0, nil
}
////////////////////////////////////////////////////////////////////////////////////
func (math *RMath)Inf(sign int) float64 {
	var v uint64
	if sign >= 0 {
		v = uvinf
	} else {
		v = uvneginf
	}
	return math.Float64frombits(v)
}
func (math *RMath)log1p(x float64) float64 {


	// special cases
	switch {
	case x < -1 || math.IsNaN(x): // includes -Inf
		return math.NaN()
	case x == -1:
		return math.Inf(-1)
	case math.IsInf(x, 1):
		return math.Inf(1)
	}

	absx := math.Abs(x)

	var f float64
	var iu uint64
	k := 1
	if absx < Sqrt2M1 { //  |x| < Sqrt(2)-1
		if absx < Small { // |x| < 2**-29
			if absx < Tiny { // |x| < 2**-54
				return x
			}
			return x - x*x*0.5
		}
		if x > Sqrt2HalfM1 { // Sqrt(2)/2-1 < x
			// (Sqrt(2)/2-1) < x < (Sqrt(2)-1)
			k = 0
			f = x
			iu = 1
		}
	}
	var c float64
	if k != 0 {
		var u float64
		if absx < Two53 { // 1<<53
			u = 1.0 + x
			iu = math.Float64bits(u)
			k = int((iu >> 52) - 1023)
			// correction term
			if k > 0 {
				c = 1.0 - (u - x)
			} else {
				c = x - (u - 1.0)
			}
			c /= u
		} else {
			u = x
			iu = math.Float64bits(u)
			k = int((iu >> 52) - 1023)
			c = 0
		}
		iu &= 0x000fffffffffffff
		if iu < 0x0006a09e667f3bcd { // mantissa of Sqrt(2)
			u = math.Float64frombits(iu | 0x3ff0000000000000) // normalize u
		} else {
			k++
			u = math.Float64frombits(iu | 0x3fe0000000000000) // normalize u/2
			iu = (0x0010000000000000 - iu) >> 2
		}
		f = u - 1.0 // Sqrt(2)/2 < u < Sqrt(2)
	}
	hfsq := 0.5 * f * f
	var s, R, z float64
	if iu == 0 { // |f| < 2**-20
		if f == 0 {
			if k == 0 {
				return 0
			}
			c += float64(k) * Ln2Lo
			return float64(k)*Ln2Hi + c
		}
		R = hfsq * (1.0 - 0.66666666666666666*f) // avoid division
		if k == 0 {
			return f - R
		}
		return float64(k)*Ln2Hi - ((R - (float64(k)*Ln2Lo + c)) - f)
	}
	s = f / (2.0 + f)
	z = s * s
	R = z * (Lp1 + z*(Lp2+z*(Lp3+z*(Lp4+z*(Lp5+z*(Lp6+z*Lp7))))))
	if k == 0 {
		return f - (hfsq - s*(hfsq+R))
	}
	return float64(k)*Ln2Hi - ((hfsq - (s*(hfsq+R) + (float64(k)*Ln2Lo + c))) - f)
}

func (math *RMath) Mul64(x, y uint64) (hi, lo uint64) {
	const mask32 = 1<<32 - 1
	x0 := x & mask32
	x1 := x >> 32
	y0 := y & mask32
	y1 := y >> 32
	w0 := x0 * y0
	t := x1*y0 + w0>>32
	w1 := t & mask32
	w2 := t >> 32
	w1 += x0 * y1
	hi = x1*y1 + w2 + w1>>32
	lo = x * y
	return
}

func (math *RMath) archSin(x float64) float64 {

	panic("not implemented")
}
func (math *RMath) archAcos(x float64) float64 {
	panic("not implemented")
}
func (math *RMath)trigReduce(x float64) (j uint64, z float64) {
	const PI4 = Pi / 4
	if x < PI4 {
		return 0, x
	}
	ix := math.Float64bits(x)
	exp := int(ix>>shift&mask) - bias - shift
	ix &^= mask << shift
	ix |= 1 << shift
	digit, bitshift := uint(exp+61)/64, uint(exp+61)%64
	z0 := (mPi4[digit] << bitshift) | (mPi4[digit+1] >> (64 - bitshift))
	z1 := (mPi4[digit+1] << bitshift) | (mPi4[digit+2] >> (64 - bitshift))
	z2 := (mPi4[digit+2] << bitshift) | (mPi4[digit+3] >> (64 - bitshift))
	z2hi, _ := math.Mul64(z2, ix)
	z1hi, z1lo := math.Mul64(z1, ix)
	z0lo := z0 * ix
	lo, c :=math.Add64(z1lo, z2hi, 0)
	hi, _ := math.Add64(z0lo, z1hi, c)
	j = hi >> 61
	hi = hi<<3 | lo>>61
	lz := uint(math.LeadingZeros64(hi))
	e := uint64(bias - (lz + 1))
	hi = (hi << (lz + 1)) | (lo >> (64 - (lz + 1)))
	hi >>= 64 - shift
	hi |= e << shift
	z = math.Float64frombits(hi)
	if j&1 == 1 {
		j++
		j &= 7
		z--
	}
	return j, z * PI4
}
func (math *RMath) NaN() float64 { return math.Float64frombits(uvnan) }

func(math *RMath)Cos(input float64) float64 {

	return math.cos(input)
}

func (math *RMath)IsNaN(f float64) (is bool) {
	return f != f
}
func (math *RMath) IsInf(f float64, sign int) bool {
	return sign >= 0 && f > MaxFloat64 || sign <= 0 && f < -MaxFloat64
}
func (math *RMath)cos(x float64) float64 {
	const (
		PI4A = 7.85398125648498535156e-1  // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668e-8  // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645e-15 // 0x3ce8469898cc5170,
	)
	// special cases
	switch {
	case math.IsNaN(x) || math.IsInf(x, 0):
		return math.NaN()
	}

	// make argument positive
	sign := false
	x = math.Abs(x)

	var j uint64
	var y, z float64
	if x >= reduceThreshold {
		j, z = math.trigReduce(x)
	} else {
		j = uint64(x * (4 / Pi)) // integer part of x/(Pi/4), as integer for tests on the phase angle
		y = float64(j)           // integer part of x/(Pi/4), as float

		// map zeros to origin
		if j&1 == 1 {
			j++
			y++
		}
		j &= 7                               // octant modulo 2Pi radians (360 degrees)
		z = ((x - y*PI4A) - y*PI4B) - y*PI4C // Extended precision modular arithmetic
	}

	if j > 3 {
		j -= 4
		sign = !sign
	}
	if j > 1 {
		sign = !sign
	}

	zz := z * z
	if j == 1 || j == 2 {
		y = z + z*zz*((((((_sin[0]*zz)+_sin[1])*zz+_sin[2])*zz+_sin[3])*zz+_sin[4])*zz+_sin[5])
	} else {
		y = 1.0 - 0.5*zz + zz*zz*((((((_cos[0]*zz)+_cos[1])*zz+_cos[2])*zz+_cos[3])*zz+_cos[4])*zz+_cos[5])
	}
	if sign {
		y = -y
	}
	return y
}

// Sin returns the sine of the radian argument x.
//
// Special cases are:
//
//	Sin(±0) = ±0
//	Sin(±Inf) = NaN
//	Sin(NaN) = NaN
func (math *RMath) Sin(x float64) float64 {
	if haveArchSin {
		return math.archSin(x)
	}
	return math.sin(x)
}

func (math *RMath)sin(x float64) float64 {
	const (
		PI4A = 7.85398125648498535156e-1  // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668e-8  // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645e-15 // 0x3ce8469898cc5170,
	)
	// special cases
	switch {
	case x == 0 || math.IsNaN(x):
		return x // return ±0 || NaN()
	case math.IsInf(x, 0):
		return math.NaN()
	}

	// make argument positive but save the sign
	sign := false
	if x < 0 {
		x = -x
		sign = true
	}

	var j uint64
	var y, z float64
	if x >= reduceThreshold {
		j, z = math.trigReduce(x)
	} else {
		j = uint64(x * (4 / Pi)) // integer part of x/(Pi/4), as integer for tests on the phase angle
		y = float64(j)           // integer part of x/(Pi/4), as float

		// map zeros to origin
		if j&1 == 1 {
			j++
			y++
		}
		j &= 7                               // octant modulo 2Pi radians (360 degrees)
		z = ((x - y*PI4A) - y*PI4B) - y*PI4C // Extended precision modular arithmetic
	}
	// reflect in x axis
	if j > 3 {
		sign = !sign
		j -= 4
	}
	zz := z * z
	if j == 1 || j == 2 {
		y = 1.0 - 0.5*zz + zz*zz*((((((_cos[0]*zz)+_cos[1])*zz+_cos[2])*zz+_cos[3])*zz+_cos[4])*zz+_cos[5])
	} else {
		y = z + z*zz*((((((_sin[0]*zz)+_sin[1])*zz+_sin[2])*zz+_sin[3])*zz+_sin[4])*zz+_sin[5])
	}
	if sign {
		y = -y
	}
	return y
}
func (math *RMath)llog1p(x float64) float64 {
	const (
		Sqrt2M1     = 4.142135623730950488017e-01  // Sqrt(2)-1 = 0x3fda827999fcef34
		Sqrt2HalfM1 = -2.928932188134524755992e-01 // Sqrt(2)/2-1 = 0xbfd2bec333018866
		Small       = 1.0 / (1 << 29)              // 2**-29 = 0x3e20000000000000
		Tiny        = 1.0 / (1 << 54)              // 2**-54
		Two53       = 1 << 53                      // 2**53
		Ln2Hi       = 6.93147180369123816490e-01   // 3fe62e42fee00000
		Ln2Lo       = 1.90821492927058770002e-10   // 3dea39ef35793c76
		Lp1         = 6.666666666666735130e-01     // 3FE5555555555593
		Lp2         = 3.999999999940941908e-01     // 3FD999999997FA04
		Lp3         = 2.857142874366239149e-01     // 3FD2492494229359
		Lp4         = 2.222219843214978396e-01     // 3FCC71C51D8E78AF
		Lp5         = 1.818357216161805012e-01     // 3FC7466496CB03DE
		Lp6         = 1.531383769920937332e-01     // 3FC39A09D078C69F
		Lp7         = 1.479819860511658591e-01     // 3FC2F112DF3E5244
	)

	// special cases
	switch {
	case x < -1 || math.IsNaN(x): // includes -Inf
		return math.NaN()
	case x == -1:
		return math.Inf(-1)
	case math.IsInf(x, 1):
		return math.Inf(1)
	}

	absx := math.Abs(x)

	var f float64
	var iu uint64
	k := 1
	if absx < Sqrt2M1 { //  |x| < Sqrt(2)-1
		if absx < Small { // |x| < 2**-29
			if absx < Tiny { // |x| < 2**-54
				return x
			}
			return x - x*x*0.5
		}
		if x > Sqrt2HalfM1 { // Sqrt(2)/2-1 < x
			// (Sqrt(2)/2-1) < x < (Sqrt(2)-1)
			k = 0
			f = x
			iu = 1
		}
	}
	var c float64
	if k != 0 {
		var u float64
		if absx < Two53 { // 1<<53
			u = 1.0 + x
			iu = math.Float64bits(u)
			k = int((iu >> 52) - 1023)
			// correction term
			if k > 0 {
				c = 1.0 - (u - x)
			} else {
				c = x - (u - 1.0)
			}
			c /= u
		} else {
			u = x
			iu = math.Float64bits(u)
			k = int((iu >> 52) - 1023)
			c = 0
		}
		iu &= 0x000fffffffffffff
		if iu < 0x0006a09e667f3bcd { // mantissa of Sqrt(2)
			u = math.Float64frombits(iu | 0x3ff0000000000000) // normalize u
		} else {
			k++
			u = math.Float64frombits(iu | 0x3fe0000000000000) // normalize u/2
			iu = (0x0010000000000000 - iu) >> 2
		}
		f = u - 1.0 // Sqrt(2)/2 < u < Sqrt(2)
	}
	hfsq := 0.5 * f * f
	var s, R, z float64
	if iu == 0 { // |f| < 2**-20
		if f == 0 {
			if k == 0 {
				return 0
			}
			c += float64(k) * Ln2Lo
			return float64(k)*Ln2Hi + c
		}
		R = hfsq * (1.0 - 0.66666666666666666*f) // avoid division
		if k == 0 {
			return f - R
		}
		return float64(k)*Ln2Hi - ((R - (float64(k)*Ln2Lo + c)) - f)
	}
	s = f / (2.0 + f)
	z = s * s
	R = z * (Lp1 + z*(Lp2+z*(Lp3+z*(Lp4+z*(Lp5+z*(Lp6+z*Lp7))))))
	if k == 0 {
		return f - (hfsq - s*(hfsq+R))
	}
	return float64(k)*Ln2Hi - ((hfsq - (s*(hfsq+R) + (float64(k)*Ln2Lo + c))) - f)
}

func(math *RMath) Acosh(x float64) float64 {
	return math.acosh(x)
}
func (math *RMath) archLog1p(x float64) float64 {
	panic("not implemented")
}
func (math *RMath) Log1p(x float64) float64 {

	return math.log1p(x)
}
func(math *RMath) acosh(x float64) float64 {
	const Large = 1 << 28 // 2**28
	// first case is special case
	switch {
	case x < 1 || math.IsNaN(x):
		return math.NaN()
	case x == 1:
		return 0
	case x >= Large:
		return math.Ln(x) + Ln2 // x > 2**28
	case x > 2:
		return math.Ln(2*x - 1/(x+math.Sqrt(x*x-1))) // 2**28 > x > 2
	}
	t := x - 1
	return math.Log1p(t + math.Sqrt(2*t+t*t)) // 2 >= x > 1
}
func (math *RMath) xatan(x float64) float64 {
	const (
		P0 = -8.750608600031904122785e-01
		P1 = -1.615753718733365076637e+01
		P2 = -7.500855792314704667340e+01
		P3 = -1.228866684490136173410e+02
		P4 = -6.485021904942025371773e+01
		Q0 = +2.485846490142306297962e+01
		Q1 = +1.650270098316988542046e+02
		Q2 = +4.328810604912902668951e+02
		Q3 = +4.853903996359136964868e+02
		Q4 = +1.945506571482613964425e+02
	)
	z := x * x
	z = z * ((((P0*z+P1)*z+P2)*z+P3)*z + P4) / (((((z+Q0)*z+Q1)*z+Q2)*z+Q3)*z + Q4)
	z = x*z + x
	return z
}

func (math *RMath)satan(x float64) float64 {
	const (
		Morebits = 6.123233995736765886130e-17 
		Tan3pio8 = 2.41421356237309504880      
	)
	if x <= 0.66 {
		return math.xatan(x)
	}
	if x > Tan3pio8 {
		return Pi/2 - math.xatan(1/x) + Morebits
	}
	return Pi/4 + math.xatan((x-1)/(x+1)) + 0.5*Morebits
}

//arctan
func(math *RMath) Atan(x float64) float64 {

	return math.atan(x)
}
//arctan
func (math *RMath) atan(x float64) float64 {
	if x == 0 {
	  return x
	}
	if x > 0 {
	  return math.satan(x)
	}
	return -(math.satan(-x))
  }
func (math *RMath) archTan(x float64) float64 {
	panic("not implemented")
}
//tan
func (math *RMath) Tan(x float64) float64 {
	if haveArchTan {
		return math.archTan(x)
	}
	return math.tan(x)
}
func (math *RMath) tan(x float64) float64 {
	const (
		PI4A = 7.85398125648498535156e-1  // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668e-8  // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645e-15 // 0x3ce8469898cc5170,
	)
	// special cases
	switch {
	case x == 0 || math.IsNaN(x):
		return x // return ±0 || NaN()
	case math.IsInf(x, 0):
		return math.NaN()
	}

	// make argument positive but save the sign
	sign := false
	if x < 0 {
		x = -x
		sign = true
	}
	var j uint64
	var y, z float64
	if x >= reduceThreshold {
		j, z = math.trigReduce(x)
	} else {
		j = uint64(x * (4 / Pi)) // integer part of x/(Pi/4), as integer for tests on the phase angle
		y = float64(j)           // integer part of x/(Pi/4), as float

		/* map zeros and singularities to origin */
		if j&1 == 1 {
			j++
			y++
		}

		z = ((x - y*PI4A) - y*PI4B) - y*PI4C
	}
	zz := z * z

	if zz > 1e-14 {
		y = z + z*(zz*(((_tanP[0]*zz)+_tanP[1])*zz+_tanP[2])/((((zz+_tanQ[1])*zz+_tanQ[2])*zz+_tanQ[3])*zz+_tanQ[4]))
	} else {
		y = z
	}
	if j&2 == 2 {
		y = -1 / y
	}
	if sign {
		y = -y
	}
	
	return y
}
//cot
func(math *RMath)Cot(input float64) float64 {
	sine := 1 / input 
	cosine := 1 / input
	cotangent := cosine / sine 
	return cotangent
}
