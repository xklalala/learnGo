package utils

func Cal(n1 float64, n2 float64, operator byte) float64{
	switch operator{
		case '+':
			return n1 + n2
		case '-':
			return n1 -  n2
		case '*':
			return n1 * n2;
		case '/':
			if n2 == 0{
				return 0;
			}else{
				return n1 / n2
			}
		default:
			return 0
	}
}