package main
import "fmt"
	func findDivByThreeFive(a,b int)(sum, i int) {
		for i = 1; i < 1000; i++; {
			if i % a == 0 || i % b == 0 {
				sum+=i;
			}
		}
		return sum;
	}

	func main() {
		fmt.Println("%v\n", findDivByAandB(3,5));
	}

