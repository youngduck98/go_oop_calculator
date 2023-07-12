package rational

import . "calculator"

type Rational struct {
	Nt        Number //분자
	Dnt       Number //분모
	Precision int    //정밀도
}

func NewRational(a Number, b Number) Rational {
	return Rational{a, b, 3}
}

func (stype Rational) To_integer() Number {
	return stype.Nt.Div(stype.Dnt).To_integer()
}

func (stype Rational) To_float() Number {
	return stype.Nt.To_float().Div(stype.Dnt.To_float())
}

func (stype Rational) Intmul(n int) Number {
	stype.Nt = stype.Nt.Intmul(n)
	return stype
}

func (stype Rational) Printvalue() {
	stype.Nt.Printvalue()
	print("/")
	stype.Dnt.Printvalue()
}

func abbreviation(a Rational) Rational {
	one := a.Dnt.Div(a.Dnt).To_integer()
	factor := a.Dnt.Div(a.Dnt).To_integer()

	for a.Dnt.Bigger(factor) || a.Dnt.Same(factor) {
		if a.Nt == factor.Mul(a.Nt.Div(factor)) && a.Dnt == factor.Mul(a.Dnt.Div(factor)) {
			a.Nt = a.Nt.Div(factor)
			a.Dnt = a.Dnt.Div(factor)
			factor = one
		}
		factor = factor.Add(one).To_integer()
	}
	return a
}

func (stype Rational) Same(a Number) bool {
	if stype.Nt.Same(a.(Rational).Nt) && stype.Dnt.Same(a.(Rational).Dnt) {
		return true
	} else {
		return false
	}
}

func (stype Rational) Bigger(a Number) bool {
	zero := a.Sub(a).To_integer() //zero
	return stype.Sub(a).(Rational).Nt.Bigger(zero)
}

func (stype Rational) Convert(a Number) Number {
	stype.Dnt = a.Div(a).To_integer().Intmul(1000)
	stype.Nt = a.Intmul(1000).To_integer()
	return abbreviation(stype)
} // 정밀도 낮은 수 변환함수

func (stype Rational) Getprecision() int {
	return stype.Precision
}

func (stype Rational) Valueprint() {
	print(stype.Nt, "/", stype.Dnt)
}

func (stype Rational) Add(a Number) Number {
	stype.Nt = stype.Nt.Mul(a.(Rational).Dnt).Add(stype.Dnt.Mul(a.(Rational).Nt))
	stype.Dnt = stype.Dnt.Mul(a.(Rational).Dnt)
	stype = abbreviation(stype)
	one := stype.Dnt.Div(stype.Dnt)
	if stype.Dnt.Same(one) {
		return stype.To_integer()
	}
	return stype
}

func (stype Rational) Sub(a Number) Number {
	stype.Nt = stype.Nt.Mul(a.(Rational).Dnt).Sub(stype.Dnt.Mul(a.(Rational).Nt))
	stype.Dnt = stype.Dnt.Mul(a.(Rational).Dnt)
	stype = abbreviation(stype)
	one := stype.Dnt.Div(stype.Dnt)
	if stype.Dnt.Same(one) {
		return stype.To_integer()
	}
	return stype
}

func (stype Rational) Mul(a Number) Number {
	stype.Nt = stype.Nt.Mul(a.(Rational).Nt)
	stype.Dnt = stype.Dnt.Mul(a.(Rational).Dnt)
	stype = abbreviation(stype)
	one := stype.Dnt.Div(stype.Dnt)
	if stype.Dnt.Same(one) {
		return stype.To_integer()
	}
	return stype
}

func (stype Rational) Div(a Number) Number {
	stype.Dnt = stype.Dnt.Mul(a.(Rational).Nt)
	stype.Nt = stype.Nt.Mul(a.(Rational).Dnt)
	stype = abbreviation(stype)
	one := stype.Dnt.Div(stype.Dnt)
	if stype.Dnt.Same(one) {
		return stype.To_integer()
	}
	return stype
}
