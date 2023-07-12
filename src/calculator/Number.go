package calculator

type Number interface {
	Add(Number) Number     //같은 구현의 두 형을 더함
	Sub(Number) Number     //같은 구현의 두 형을 뺌
	Mul(Number) Number     //같은 구현의 두 형을 곱함
	Div(Number) Number     //같은 구현의 두 형을 나눔
	Convert(Number) Number //한 형의 구현을 리시버의 형의 구현으로 바꿈
	Same(Number) bool      //같은지 다른지 확인함
	Bigger(Number) bool    // 리시버 보다 작은지 확인
	Intmul(int) Number     // 해당 수 만큼 배가시킴
	To_integer() Number    //해당 number를 Integer로 변환해주는 함수
	To_float() Number      // 해당 number를 Float로 변환해주는 함수
	Getprecision() int     //해당 num의 정밀도를 출력함(ex. float는 int보다 정밀도가 높습니다.)
	Printvalue()           // test용, value출력
}
