package calculator

type Integer struct {
	Value     int
	Precision int
}

func NewInteger(a int) Integer {
	return Integer{a, 1}
}

func (stype Integer) Printvalue() {
	print(stype.Value)
}

func (stype Integer) To_integer() Number {
	return stype
}

func (stype Integer) To_float() Number {
	a := Float{0, 2}
	a.Value = float32(stype.Value)
	return a
}

func (stype Integer) Intmul(n int) Number {
	stype.Value = stype.Value * n
	return stype
}

func (stype Integer) Convert(a Number) Number {
	return a.To_integer()
}

func (stype Integer) Number(a interface{}) Number {
	return Integer{a.(int), 1}
}

func (stype Integer) Same(a Number) bool {
	if a.(Integer).Value == stype.Value {
		return true
	} else {
		return false
	}
}

func (stype Integer) Bigger(a Number) bool {
	return stype.Value > a.(Integer).Value
}

func (stype Integer) Getprecision() int {
	return stype.Precision
}

func (stype Integer) Valueprint() {
	print(stype.Value)
}

func (stype Integer) Getvalue() interface{} {
	return stype.Value
}

func (stype Integer) Add(a Number) Number {
	stype.Value = stype.Value + a.(Integer).Value
	return stype
}

func (stype Integer) Sub(a Number) Number {
	stype.Value = stype.Value - a.(Integer).Value
	return stype
}

func (stype Integer) Mul(a Number) Number {
	stype.Value = stype.Value * a.(Integer).Value
	return stype
}

func (stype Integer) Div(a Number) Number {
	stype.Value = stype.Value / a.(Integer).Value
	return stype
}
