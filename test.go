package main

import "fmt"

func main()  {
	count:=0
	fmt.Println("Welcome to The Game")

	fmt.Printf("Enter Your Name:")
	var name string
	fmt.Scan(&name)

	fmt.Printf(" Hello %v Welcome to The Game \n",name)

	var age uint
	fmt.Printf(" Enter your Age :")
	fmt.Scan(&age)

	if age<15 {
		fmt.Println("You are not allowed to play the game")
		return
	}else {
		fmt.Printf("Enjoy the game \n")
	}
	fmt.Println()

	fmt.Printf("Here is your First Question \n")
	fmt.Printf("What is the Squar of 5 ?\n" )
	var ans1 uint
	fmt.Scan(&ans1)
	if ans1==25 {
		count++
		
	}

	fmt.Println()

	fmt.Printf("Here is your Second  Question \n")
	fmt.Printf("What is the Capital of IND ?\n" )
	var ans2 string
	fmt.Scan(&ans2)
	if ans2=="Delhi" {
		count++
		
	}

	fmt.Println()

	fmt.Printf("Here is your Third Question \n")
	fmt.Printf("What is the squar of your agg?\n" )
	var ans3 uint
	fmt.Scan(&ans3)
	if ans3== age*age {
		count++
		
	}

	fmt.Println()

	fmt.Printf("Here is your Fourth Question \n")
	fmt.Printf("What is the capital of Guj ?\n" )
	var ans4 string
	fmt.Scan(&ans4)
	if ans4 == "AHD" {
		count++
		
	}

	

	fmt.Printf("Here is your last Question \n")
	fmt.Printf("What is the Squar of 2 ?\n" )
	var ans6 uint
	fmt.Scan(&ans6)
	if ans6 == 4 {
		count++
		fmt.Println("Correct")
		
	}

	
	var Score int
	Score = count*100/5
	fmt.Printf("Your Score is : %v%%,",Score)
	
	
}


