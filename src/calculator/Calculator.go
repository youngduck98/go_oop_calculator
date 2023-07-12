package calculator

type Calculator struct {
	Calstack Stack //계산기 내부 stack
}

func (cal *Calculator) Getdata(num Number) {
	cal.Calstack.Push(num) //stack에 넣어줌
}

func (cal *Calculator) Answer() {
	if cal.Calstack.size == 1 {
		print("answer: ")
		cal.Calstack.Peek().Printvalue()
	} else {
		println("error: need more op")
	}
} // 답 출력

func (cal *Calculator) Getop(op string) {
	if cal.Calstack.size >= 2 {
		a := cal.Calstack.Pop()
		b := cal.Calstack.Pop()
		if a.Getprecision() > b.Getprecision() {
			b = a.Convert(b)
		} else if a.Getprecision() < b.Getprecision() {
			a = b.Convert(a)
		}
		a.Printvalue()
		print(" ", op, " ")
		b.Printvalue()
		print(" = ")
		if op == "+" {
			cal.Calstack.Push(a.Add(b)) //더하기
			cal.Calstack.Peek().Printvalue()
			println()
		} else if op == "-" {
			cal.Calstack.Push(a.Sub(b)) //빼기
			cal.Calstack.Peek().Printvalue()
			println()
		} else if op == "*" {
			cal.Calstack.Push(a.Mul(b)) //곱하기
			cal.Calstack.Peek().Printvalue()
			println()
		} else if op == "/" {
			cal.Calstack.Push(a.Div(b)) //나누기
			cal.Calstack.Peek().Printvalue()
			println()
		}
	} else {
		println("error: need more value before op")
	}
}
