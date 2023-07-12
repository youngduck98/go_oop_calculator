# go_oop_calculator
새로운 type(유리수, 복소수 같은)에 대해 비 파괴적 확장이 가능한 interface기반 calculator

구성요소:(src파일에 들어있음)

calculator: 계산기 패키지(number, integer, float, stack, calculator 포함)

calculator_expandable: main 패키지

ncomplex: 복소수 패키지

rational: 유리수 패키지


간략 리포트:
1. interface Number를 만듬(number.go참조)
2. 해당 자료형을 기반으로 한list based stack을 만듬
- push: stack에 값 추가
- pop: stack pop
- peek: stack 맨 위의 값 가져오기(pop 없이)
3. calculator 만듬
  
-: func (cal *Calculator) getdata(num number)

//num을 계산기의 stack에 추가(형과 구현이 분리되어 있어 가능)

-: func (cal *Calculator) Answer()

//해당 계산기에서 계산한 값을 출력

-: func (cal *Calculator) getop(op string)

// op가 들어오면 2개를 스택에서 pop하여 두 정밀도를 비교함

// 이후 한 쪽의 정밀도가 다른 한쪽보다 크다면 작은 쪽의 형을 큰 쪽의 형으로 동적으로 변환함

ex) if a.Getprecision() > b.Getprecision() {b = a.Convert(b)}}

4. number인터페이스의 구현을 통해 각 type 구조체 연결(duck typing)

사고의 흐름:
1. number로 각 type을 묶고
2. 해당 number를 base로 stack과 calculator를 만들고
3. 안의 각 함수(add, sub 등등)가 number의 구현에 따라 dynamic binding되게 하자
4. number에서 각 구현의 값이 필요할 때는 type assertion을 활용하자
5. 가장 높은 정확도로 계산되어야 하므로 높은 정밀도로의 형 변환을 추가하자
ex: (integer -> float or float -> rational)
정밀도 순서: integer->float->rational->complex

어려웠던 점이나 미비했던 점:
1. 동적 바인딩으로 각 case마다 함수를 분리 하려 함
-> go는 오버로딩이 없고 리시버만 존재한다는 것을 깨달음
-> 각 리시버당 같은 이름의 함수를 하나 밖에 생성하지 못한다는 것을 깨달음
-> 이때 언어를 바꿨어야 했으나 같은 자료형의 계산만 할 수 있도록 계산기를 짜면 된다는 무리한 생각을 함
-> 결국 계산 전에 두 수의 정밀도를 비교하는 과정이 생기게 됨.(타입 check와 다를게 없다 느껴짐.)
ex) if a.Getprecision() > b.Getprecision() {b = a.Convert(b)}} 
-> 지금 생각하면 원하는 기능이 있는(오버라이딩이 있는) 언어로 작업 했어야 했음.

참고: getprecision, convert: number interface에 있는 함수
(a 와 b가 계산된다면 계산 전에 a,b의 정밀도를 검사해서 a의 정밀도가 위면 a.convert(b)로 b를 a의 형으로 바꿈)
(convert함수는 interface number에 있는 함수임. 매개변수의 형을 리시버의 형으로 바꿈)

2. 라이브러리를 통한 확인
-> go는 비코딩적 라이브러리를 지원하지 않는다는 것을 깨달음
-> 이때 언어를 바꿨어야 했으나 결국 시간만 낭비하다 실패함
-> program1, program2를 분리하지 못하고 하나의 프로그램만 내고 대신 스크린 샷으로 import만으로 비파괴적 확장이 가능하다는 것만 보임
-> 단지 이것으로는 속임수를 썻는지 확인이 힘들기 때문에 채점자에게 불친절한 방식이었음

3. Rational, Ncomplex의 멤버 초기화
-> Number라는 인터페이스에 기초하여 만들어졌음(rational, ncomplex의 멤버 모두 Number)
-> 그렇기에 rational의 멤버로 integer가 아닌 float를 넣어도 컴파일러가 알지 못하고 계산 중에나 알게 되는 문제 있음 
-> 이는 Ncomplex도 마찬가지임
-> 해당 코드는 interface가 number밖에 없었기에 뒤에 만드는 사람이 두 형식을 모를 수 밖에 없다고 생각하여 이렇게 짰지만 아쉬움이 남음.

4. Rational과 Ncomplex의 멤버 값을 설정하는 부분
->기본적으로 integer와 float의 구현를 가지는 Number를 멤버로 가져서 원하는 값을 가지는 멤버변수를 만들기 힘들었음
(뒤에 만드는 애들은 integer와 float의 구조를 모른다고 생각 하기 때문)
-> 결국 to_float과 to_integer, Intmul 따위를 활용해서 엄청 복잡하게 만들어야 했음
ex) 새로 만드는 파일(Rational.go , Ncomplex.go) 에서 1000을 가지는 integer를 만드는 방법: 
-> a.Dnt.Div(a.Dnt).To_integer()).Intmul(1000): rational a를 rational a로 나눠서 그걸 인티져로 바꾸고 거기 1000을 곱한다)
-> 개인적으로 더 좋은 방법이 있을 것 같아 아쉬움. 
-> 특히 intmul이 부자연스럽다고 느낌. 뒤에 rational하고 complex를 만든다는 생각이 없었다면 number에 해당 함수를 추가 할 수 있었을까?
하는 느낌을 주는 함수임. 

참고: to_integer, to_float, intmul(다 number에 있는 함수들)

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
