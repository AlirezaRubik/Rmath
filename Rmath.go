package Rmath
import (
	"fmt"
	"unsafe"
	"math/bits"
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
func Float64frombits(b uint64) float64 { return *(*float64)(unsafe.Pointer(&b)) }
func Float64bits(f float64) uint64     { return *(*uint64)(unsafe.Pointer(&f)) }

// Abs
func (math *RMath) Abs(input float64) float64 {
	return Float64frombits(Float64bits(input) &^ (1 << 63))
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
func (math *RMath) Reverse(input int) (sum int, count int, ave float64) {
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
	return Float64frombits(v)
}
func (math *RMath)log1p(x float64) float64 {
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
	case x < -1 || IsNaN(x): // includes -Inf
		return NaN()
	case x == -1:
		return math.Inf(-1)
	case IsInf(x, 1):
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
			iu = Float64bits(u)
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
			iu = Float64bits(u)
			k = int((iu >> 52) - 1023)
			c = 0
		}
		iu &= 0x000fffffffffffff
		if iu < 0x0006a09e667f3bcd { // mantissa of Sqrt(2)
			u = Float64frombits(iu | 0x3ff0000000000000) // normalize u
		} else {
			k++
			u = Float64frombits(iu | 0x3fe0000000000000) // normalize u/2
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


func archSin(x float64) float64 {

	panic("not implemented")
}
func (math *RMath) archAcos(x float64) float64 {
	panic("not implemented")
}
func trigReduce(x float64) (j uint64, z float64) {
	const PI4 = Pi / 4
	if x < PI4 {
		return 0, x
	}
	ix := Float64bits(x)
	exp := int(ix>>shift&mask) - bias - shift
	ix &^= mask << shift
	ix |= 1 << shift
	digit, bitshift := uint(exp+61)/64, uint(exp+61)%64
	z0 := (mPi4[digit] << bitshift) | (mPi4[digit+1] >> (64 - bitshift))
	z1 := (mPi4[digit+1] << bitshift) | (mPi4[digit+2] >> (64 - bitshift))
	z2 := (mPi4[digit+2] << bitshift) | (mPi4[digit+3] >> (64 - bitshift))
	z2hi, _ := bits.Mul64(z2, ix)
	z1hi, z1lo := bits.Mul64(z1, ix)
	z0lo := z0 * ix
	lo, c := bits.Add64(z1lo, z2hi, 0)
	hi, _ := bits.Add64(z0lo, z1hi, c)
	j = hi >> 61
	hi = hi<<3 | lo>>61
	lz := uint(bits.LeadingZeros64(hi))
	e := uint64(bias - (lz + 1))
	hi = (hi << (lz + 1)) | (lo >> (64 - (lz + 1)))
	hi >>= 64 - shift
	hi |= e << shift
	z = Float64frombits(hi)
	if j&1 == 1 {
		j++
		j &= 7
		z--
	}
	return j, z * PI4
}
func NaN() float64 { return Float64frombits(uvnan) }

func(math *RMath)Cos(input float64) float64 {

	return math.cos(input)
}

func IsNaN(f float64) (is bool) {
	return f != f
}
func IsInf(f float64, sign int) bool {
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
	case IsNaN(x) || IsInf(x, 0):
		return NaN()
	}

	// make argument positive
	sign := false
	x = math.Abs(x)

	var j uint64
	var y, z float64
	if x >= reduceThreshold {
		j, z = trigReduce(x)
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
		return archSin(x)
	}
	return sin(x)
}

func sin(x float64) float64 {
	const (
		PI4A = 7.85398125648498535156e-1  // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668e-8  // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645e-15 // 0x3ce8469898cc5170,
	)
	// special cases
	switch {
	case x == 0 || IsNaN(x):
		return x // return ±0 || NaN()
	case IsInf(x, 0):
		return NaN()
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
		j, z = trigReduce(x)
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
func (math*RMath)llog1p(x float64) float64 {
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
	case x < -1 || IsNaN(x): // includes -Inf
		return NaN()
	case x == -1:
		return math.Inf(-1)
	case IsInf(x, 1):
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
			iu = Float64bits(u)
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
			iu = Float64bits(u)
			k = int((iu >> 52) - 1023)
			c = 0
		}
		iu &= 0x000fffffffffffff
		if iu < 0x0006a09e667f3bcd { // mantissa of Sqrt(2)
			u = Float64frombits(iu | 0x3ff0000000000000) // normalize u
		} else {
			k++
			u = Float64frombits(iu | 0x3fe0000000000000) // normalize u/2
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
	case x < 1 || IsNaN(x):
		return NaN()
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
func xatan(x float64) float64 {
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

func satan(x float64) float64 {
	const (
		Morebits = 6.123233995736765886130e-17 
		Tan3pio8 = 2.41421356237309504880      
	)
	if x <= 0.66 {
		return xatan(x)
	}
	if x > Tan3pio8 {
		return Pi/2 - xatan(1/x) + Morebits
	}
	return Pi/4 + xatan((x-1)/(x+1)) + 0.5*Morebits
}

//arctan
func(math *RMath) Atan(x float64) float64 {

	return atan(x)
}
//arctan
func atan(x float64) float64 {
	if x == 0 {
		return x
	}
	if x > 0 {
		return satan(x)
	}
	return -satan(-x)
}
func archTan(x float64) float64 {
	panic("not implemented")
}
//tan
func (math *RMath) Tan(x float64) float64 {
	if haveArchTan {
		return archTan(x)
	}
	return tan(x)
}
func tan(x float64) float64 {
	const (
		PI4A = 7.85398125648498535156e-1  // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668e-8  // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645e-15 // 0x3ce8469898cc5170,
	)
	// special cases
	switch {
	case x == 0 || IsNaN(x):
		return x // return ±0 || NaN()
	case IsInf(x, 0):
		return NaN()
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
		j, z = trigReduce(x)
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
