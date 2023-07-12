package calculator

type Float struct {
	Value     float32
	Precision int
}

func NewFloat(a float32) Float {
	return Float{a, 2}
}

func (stype Float) Printvalue() {
	println(stype.Value)
}

func (stype Float) To_integer() Number {
	a := Integer{0, 1}
	a.Value = int(stype.Value)
	return a
}

func (stype Float) To_float() Number {
	a := Float{0, 2}
	a.Value = float32(stype.Value)
	return a
}

func (stype Float) Intmul(n int) Number {
	stype.Value = stype.Value * float32(n)
	return stype
}

func (stype Float) Number(a interface{}) Number {
	return Float{a.(float32), 2}
}

func (stype Float) Convert(a Number) Number {
	return a.To_float()
}

func (stype Float) Getvalue() interface{} {
	return stype.Value
}

func (stype Float) Same(a Number) bool {
	if a.(Float).Value == stype.Value {
		return true
	} else {
		return false
	}
}

func (stype Float) Bigger(a Number) bool {
	if stype.Value > a.(Float).Value {
		return true
	} else {
		return false
	}
}

func (stype Float) Getprecision() int {
	return stype.Precision
}

func (stype Float) Valueprint() {
	print(stype.Value)
}

func (stype Float) Add(a Number) Number {
	stype.Value = stype.Value + a.(Float).Value
	return stype
}

func (stype Float) Sub(a Number) Number {
	stype.Value = stype.Value - a.(Float).Value
	return stype
}

func (stype Float) Mul(a Number) Number {
	stype.Value = (stype.Value) * (a.(Float).Value)
	return stype
}

func (stype Float) Div(a Number) Number {
	stype.Value = (stype.Value) / (a.(Float).Value)
	return stype
}
