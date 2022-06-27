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