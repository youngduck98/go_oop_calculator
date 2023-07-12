package main

/*
사용방법:
입력 순서: RNF
입력 형식: cal.getdata(구조체 생성자(value))
integer 입력: cal.Getdata(NewInteger(2))
Float 입력: cal.Getdata(NewFloat(2.1))
Rational 입력: NewRational(NewInteger(1), NewInteger(2))
complex 입력: NewNcomplex(NewFloat(2.0), NewFloat(2.0))
op 입력: cal.Getop("*")
식이 알맞게 입력 되었다면 계산값 출력
식에 op가 남거나 부족 하다면 error출력
*/

import (
	//. "Ncomplex"
	. "calculator"
	//. "rational"
)

func main() {
	println("start!")
	cal := Calculator{}
	println("ready cal")
	cal.Getdata(NewFloat(1.2))
	cal.Getdata(NewInteger(2))
	//cal.Getdata(NewRational(NewInteger(1), NewInteger(3)))
	//cal.Getdata(NewNcomplex(NewFloat(2.0), NewFloat(2.0)))
	cal.Getop("-")
	//cal.Getop("*")
	//cal.Getop("+")
	cal.Answer()
}
