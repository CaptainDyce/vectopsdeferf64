package vectopsdeferf64

import (
	fs "github.com/CaptainDyce/f64supp"
	is "github.com/CaptainDyce/intsupp"
)

type VectOp struct {
	prev *VectOp         // nil if root operation
	op   func([]float64) // nil if root operation
}

func (s VectOp) build(op func([]float64)) VectOp {
	if s.op == nil {
		// root
		return VectOp{nil, op}
	}
	return VectOp{&s, op}
}

var root = VectOp{nil, nil}

///////////////////////////
// bootstrap...
///////////////////////////

func Op() VectOp {
	return root
}

///////////////////////////
// operations...
///////////////////////////

func (s VectOp) Apply(f fs.IndexedFunc) VectOp {
	return s.build(func(v []float64) {
		fs.Apply(v, f)
	})
}

func (s VectOp) ApplyOp(f fs.Operator) VectOp {
	return s.build(func(v []float64) {
		fs.ApplyOp(v, f)
	})
}

func (s VectOp) ApplyOpi(f fs.IndexedOperator) VectOp {
	return s.build(func(v []float64) {
		fs.ApplyOpi(v, f)
	})
}

func (s VectOp) Ident() VectOp {
	return s.Apply(fs.CoerceInt)
}

func (s VectOp) Setl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Setl(v, value)
	})
}

func (s VectOp) Setv(v1 []float64) VectOp {
	return s.build(func(v []float64) {
		fs.Setv(v, v1)
	})
}

func (s VectOp) SetMaskl(value float64, p is.Predicate) VectOp {
	return s.build(func(v []float64) {
		fs.SetMaskl(v, value, p)
	})
}

func (s VectOp) SetMaskv(v1 []float64, p is.Predicate) VectOp {
	return s.build(func(v []float64) {
		fs.SetMaskv(v, v1, p)
	})
}

/////////////////
// Plus
/////////////////
// -> s'[i] = s[i] + v[i]
func (s VectOp) Plusv(v1 []float64) VectOp {
	return s.build(func(v []float64) {
		fs.Plusv(v, v1)
	})
}

// -> s'[i] = s[i] + value
func (s VectOp) Plusl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Plusl(v, value)
	})
}

// -> s'[i] = s[i] + o(i)
func (s VectOp) PlusOp(o fs.IndexedFunc) VectOp {
	return s.build(func(v []float64) {
		fs.PlusOp(v, o)
	})
}

// -> s'[i] = s[i] + o(i, s[i])
func (s VectOp) PlusOpi(o fs.IndexedOperator) VectOp {
	return s.build(func(v []float64) {
		fs.PlusOpi(v, o)
	})
}

/////////////////
// Minus
/////////////////
// -> s'[i] = s[i] - v[i]
func (s VectOp) Minusv(v1 []float64) VectOp {
	return s.build(func(v []float64) {
		fs.Minusv(v, v1)
	})
}

// -> s'[i] = s[i] - value
func (s VectOp) Minusl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Minusl(v, value)
	})
}

// -> s'[i] = s[i] - o(i)
func (s VectOp) MinusOp(o fs.IndexedFunc) VectOp {
	return s.build(func(v []float64) {
		fs.MinusOp(v, o)
	})
}

// -> s'[i] = s[i] - o(i, s[i])
func (s VectOp) MinusOpi(o fs.IndexedOperator) VectOp {
	return s.build(func(v []float64) {
		fs.MinusOpi(v, o)
	})
}

/////////////////
// Times
/////////////////
// -> s'[i] = s[i] * v[i]
func (s VectOp) Timesv(v1 []float64) VectOp {
	return s.build(func(v []float64) {
		fs.Timesv(v, v1)
	})
}

// -> s'[i] = s[i] * value
func (s VectOp) Timesl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Timesl(v, value)
	})
}

// -> s'[i] = s[i] * o(i)
func (s VectOp) TimesOp(o fs.IndexedFunc) VectOp {
	return s.build(func(v []float64) {
		fs.TimesOp(v, o)
	})
}

// -> s'[i] = s[i] * o(i, s[i])
func (s VectOp) TimesOpi(o fs.IndexedOperator) VectOp {
	return s.build(func(v []float64) {
		fs.TimesOpi(v, o)
	})
}

/////////////////
// Div
/////////////////
// -> s'[i] = s[i] / v[i]
func (s VectOp) Divv(v1 []float64) VectOp {
	return s.build(func(v []float64) {
		fs.Divv(v, v1)
	})
}

// -> s'[i] = s[i] / value
func (s VectOp) Divl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Divl(v, value)
	})
}

// -> s'[i] = s[i] / o(i)
func (s VectOp) DivOp(o fs.IndexedFunc) VectOp {
	return s.build(func(v []float64) {
		fs.DivOp(v, o)
	})
}

// -> s'[i] = s[i] / o(i, s[i])
func (s VectOp) DivOpi(o fs.IndexedOperator) VectOp {
	return s.build(func(v []float64) {
		fs.DivOpi(v, o)
	})
}

/////////////////
// Pow
/////////////////
// -> s'[i] = s[i] ^ v[i] (as in e.g. 2^3...)
func (s VectOp) Powv(v1 []float64) VectOp {
	return s.build(func(v []float64) {
		fs.Powv(v, v1)
	})
}

// -> s'[i] = s[i] ^ value
func (s VectOp) Powl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Powl(v, value)
	})
}

// -> s'[i] = s[i] ^ o(i)
func (s VectOp) PowOp(o fs.IndexedFunc) VectOp {
	return s.build(func(v []float64) {
		fs.PowOp(v, o)
	})
}

// -> s'[i] = s[i] ^ o(i, s[i])
func (s VectOp) PowOpi(o fs.IndexedOperator) VectOp {
	return s.build(func(v []float64) {
		fs.PowOpi(v, o)
	})
}

/////////////////
// Max
/////////////////
// -> s'[i] = max(s[i], v[i])
func (s VectOp) Maxv(v1 []float64) VectOp {
	return s.build(func(v []float64) {
		fs.Maxv(v, v1)
	})
}

// -> s'[i] = max(s[i], value)
func (s VectOp) Maxl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Maxl(v, value)
	})
}

// -> s'[i] = max(s[i], o(i))
func (s VectOp) MaxOp(o fs.IndexedFunc) VectOp {
	return s.build(func(v []float64) {
		fs.MaxOp(v, o)
	})
}

// -> s'[i] = max(s[i], o(i, s[i]))
func (s VectOp) MaxOpi(o fs.IndexedOperator) VectOp {
	return s.build(func(v []float64) {
		fs.MaxOpi(v, o)
	})
}

/////////////////
// Min
/////////////////
// -> s'[i] = min(s[i], v[i])
func (s VectOp) Minv(v1 []float64) VectOp {
	return s.build(func(v []float64) {
		fs.Minv(v, v1)
	})
}

// -> s'[i] = min(s[i], value)
func (s VectOp) Minl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Minl(v, value)
	})
}

// -> s'[i] = min(s[i], o(i))
func (s VectOp) MinOp(o fs.IndexedFunc) VectOp {
	return s.build(func(v []float64) {
		fs.MinOp(v, o)
	})
}

// -> s'[i] = min(s[i], o(i, s[i]))
func (s VectOp) MinOpi(o fs.IndexedOperator) VectOp {
	return s.build(func(v []float64) {
		fs.MinOpi(v, o)
	})
}

///////////////////////////
// misc...
///////////////////////////

func (s VectOp) Rev() VectOp {
	return s.build(func(v []float64) {
		fs.Rev(v)
	})
}

// -> s'[i] = -s[i]
func (s VectOp) Neg() VectOp {
	return s.build(func(v []float64) {
		fs.Negv(v)
	})
}

// -> s'[i] = abs(s[i])
func (s VectOp) Abs() VectOp {
	return s.build(func(v []float64) {
		fs.Abs(v)
	})
}

// -> s'[i] = value / s[i]
func (s VectOp) Idivl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Idivl(v, value)
	})
}

// -> s'[i] = 1 / s[i]
func (s VectOp) Inv() VectOp {
	return s.build(func(v []float64) {
		fs.Idivl(v, 1)
	})
}

// -> s'[i] = log(s[i])
func (s VectOp) Log() VectOp {
	return s.build(func(v []float64) {
		fs.Log(v)
	})
}

// -> s'[i] = exp(s[i])
func (s VectOp) Exp() VectOp {
	return s.build(func(v []float64) {
		fs.Exp(v)
	})
}

// -> s'[i] = value ^ s[i]
func (s VectOp) Expl(value float64) VectOp {
	return s.build(func(v []float64) {
		fs.Expl(v, value)
	})
}

/////////////////
// Terminals
/////////////////

// applies the complete operation chain to v
func (s VectOp) On(v []float64) []float64 {
	if s.prev != nil {
		s.prev.On(v)
	}
	s.op(v)
	return v
}

// applies the complete operation chain to a new [size] array
func (s VectOp) OnSize(size int) []float64 {
	v := make([]float64, size)
	return s.On(v)
}

// applies the complete operation chain to a new [size] <value>-filled array
func (s VectOp) OnConst(value float64, size int) []float64 {
	v := make([]float64, size)
	fs.Setl(v, value)
	return s.On(v)
}

// applies the complete operation chain to a new [size] identity array
func (s VectOp) OnIdent(size int) []float64 {
	v := make([]float64, size)
	fs.Ident(v)
	return s.On(v)
}
