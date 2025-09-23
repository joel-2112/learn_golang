package main

import "fmt"

func main() {
	//if statement
	//exeute a block of code if a specified condition is true
	x := 10
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}
	//nested if statement
	score := 95
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
	//loops
	//for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("iteration:", i)
	}
	//nested for loop
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d * %d =%d\t", i, j, i*j)
		}
		fmt.Println("row", i, "completed")
	}

	//while loop (using for loop)
	count := 1
	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}
	//infinite loop
	// for {
	// 	fmt.Println("This loop will run forever")
	// }
	//break statement
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break // Exit the loop when i is 6
		}
		fmt.Println("Number:", i)
	}
	//continue statement
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd Number:", i)
	}

	//goto statement
	i := 1
Here:
	if i <= 5 {
		fmt.Println("i:", i)
		i++
		goto Here 
	}
	fmt.Println("Loop ended")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			goto Skip
		}
		fmt.Println("i:", i)
	}
	Skip:
		fmt.Println("Skipped when i was 3")

	//switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
	//defer statement
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")

}
