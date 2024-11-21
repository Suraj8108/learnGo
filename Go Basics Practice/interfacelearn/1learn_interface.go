package interfacelearn

import "fmt"

func LearnInterface() {
	value1 := 1
	value2 := 1.1342
	extraceType(value1)
	extraceType(value2)

	sumValue := summation(1, 4)
	extraceType(sumValue)
}

func extraceType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer", value)
	case float64: 
		fmt.Println("Float")
	default:
		fmt.Println("Not the type")
	}
}

func summation(value1 interface{}, value2 interface{}) interface{}{
	intValue1, isInt1 := value1.(int);
	intValue2, isInt2 := value2.(int);

	if isInt1 && isInt2 {
		return intValue1 + intValue2;
	}

	return nil
}