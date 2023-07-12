package Ncomplex

import . "calculator"

type Ncomplex struct {
	Real      Number
	Img       Number
	Precision int
}

func NewNcomplex(a Number, b Number) Ncomplex {
	return Ncomplex{a, b, 4}
}

func (stype Ncomplex) Convert(a Number) Number {
	stype.Real = a.To_float()
	stype.Img = a.To_float().Intmul(0)
	return stype
}

func (stype Ncomplex) To_float() Number {
	return stype.Real.To_float()
}

func (stype Ncomplex) To_integer() Number {
	return stype.Real.To_integer()
}

func (stype Ncomplex) Printvalue() {
	print("real: ")
	stype.Real.Printvalue()
	print("img: ")
	stype.Img.Printvalue()
}

func (stype Ncomplex) Intmul(a int) Number {
	stype.Real = stype.Real.Intmul(a)
	stype.Img = stype.Img.Intmul(a)
	return stype
}

func (stype Ncomplex) Bigger(a Number) bool {
	if stype.Img.Same(stype.Img.Intmul(0)) && a.(Ncomplex).Img.Same(a.(Ncomplex).Img.Intmul(0)) {
		if stype.Real.Bigger(a.(Ncomplex).Real) {
			return true
		} else {
			return false
		}
	}
	return false
}

func (stype Ncomplex) Same(a Number) bool {
	if stype.Real.Same(a.(Ncomplex).Real) && stype.Img.Same(a.(Ncomplex).Img) {
		return true
	}
	return false
}

func (stype Ncomplex) Getprecision() int {
	return stype.Precision
}

func (stype Ncomplex) Valueprint() {
	print("real: ")
	stype.Real.Printvalue()
	print("\nimg: ")
	stype.Img.Printvalue()
}

func (stype Ncomplex) Add(a Number) Number {
	n1 := stype.Real.Add(a.(Ncomplex).Real)
	n2 := stype.Img.Add(a.(Ncomplex).Img)
	stype.Real = n1
	stype.Img = n2
	if stype.Img.Same(stype.Img.Intmul(0)) {
		return stype.To_float()
	}
	return stype
}

func (stype Ncomplex) Sub(a Number) Number {
	n1 := stype.Real.Sub(a.(Ncomplex).Real)
	n2 := stype.Img.Sub(a.(Ncomplex).Img)
	stype.Real = n1
	stype.Img = n2
	if stype.Img.Same(stype.Img.Intmul(0)) {
		return stype.To_float()
	}
	return stype
}

func (stype Ncomplex) Mul(a Number) Number {
	n1 := stype.Real.Mul(a.(Ncomplex).Real).Sub(stype.Img.Mul(a.(Ncomplex).Img))
	n2 := stype.Real.Mul(a.(Ncomplex).Img).Add(stype.Img.Mul(a.(Ncomplex).Real))
	stype.Real = n1
	stype.Img = n2
	if stype.Img.Same(stype.Img.Intmul(0)) {
		return stype.To_float()
	}
	return stype
}

func (stype Ncomplex) Div(a Number) Number {
	dn := a.(Ncomplex).Real.Mul(a.(Ncomplex).Real).Add(a.(Ncomplex).Img.Mul(a.(Ncomplex).Img))
	unr := stype.Real.Mul(a.(Ncomplex).Real).Add(stype.Img.Mul(a.(Ncomplex).Img))
	uni := stype.Real.Mul(a.(Ncomplex).Img).Sub(stype.Img.Mul(a.(Ncomplex).Real))
	stype.Real = unr.Div(dn)
	stype.Img = uni.Div(dn)
	if stype.Img.Same(stype.Img.Intmul(0)) {
		return stype.To_float()
	}
	return stype
}
