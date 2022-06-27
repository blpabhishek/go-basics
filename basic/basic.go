package basic


func LCM(a int,b int) int {
	var max int
	if max = b; a>b {
		max = a
	}
	for(!(max%a==0 && max%b==0)) {
		max+=1
	}
	return max
}

func GCD(num1 int, num2 int ) int {
	for i := num2; i > 1; i-- {
		if(num2%i==0 && num1%i==0){
			return i;
		}
	}
	return 1;
}

func foo() func() int {
	a := 0
	return func () int{
		a+=1
		return a
	}
}

func IsEven(num int) bool  {
	return num%2==0
}

func Square(num int) int {
	return num*num
}

func Cube(num int) int {
	return Square(num) * num
}

func Sum(num int) int {
	var sum int
	for i := 0; i < num; i++ {
		sum+=i
	}
	return sum
}

func Factorial(num int) int {
	if(num==1){
		return 1;
	}
	return Factorial(num-1)* num
}

func Fibonacci() func() int {
	var first int = 0
	var second  int = 1
	return func () int {
		third:= first + second
		first = second
		second = third
		return third
	}
}

func TakeFibonacci(nth int) []int {
	gen := Fibonacci()
	terms:=make([]int, nth)
	for i := range terms {
		terms[i] = gen()
	}
	return terms
}

func FilterEven(nums []int) []int  {
	slice:= make([]int, 0)
	for _, v := range nums {
		if IsEven(v) {
			slice = append(slice,v)				
		}
	}
	return slice
}