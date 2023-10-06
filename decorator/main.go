package main

import (
	"Sergaziyev_ass2/decorator/university"
	"fmt"
)

//in this code, options such as news subscription,
//access to online courses and library access are dynamically added
//to the student's basic account without changing its original BasicStudentAccount class

func main() {
	// input of student name
	fmt.Print("enter student name: ")
	var studentName string
	fmt.Scanln(&studentName)

	// making of base account with 100 points (points in this code it is like currency thanks
	//to which you open new courses and so on, like in-game currency but inside the programming
	//course (to make it more interesting to learn)
	basicStudentAccount := university.NewStudentAccount(studentName, "Baha@example.com", 100)
	fmt.Println("account replenished with $100.")

	fmt.Println("choose options")
	fmt.Println("1. withNewsSubscription: News subscription (cost 20).")
	fmt.Println("2. withOnlineCourses: Access to online courses.")
	fmt.Println("3. withLibraryAccess: Library access.")

	var choice int
	fmt.Print("Enter the number of the selected option: ")
	fmt.Scanln(&choice)

	var studentAccount university.StudentAccount

	switch choice {
	case 1:
		studentAccount = university.NewNewsSubscriptionDecorator(basicStudentAccount, 20)
	case 2:

		fmt.Println("List of available courses:")
		fmt.Println("1. Web Technologies (20 points).")
		fmt.Println("2. Operating Systems (20 points).")
		fmt.Println("3. Software Design Patterns (20 points).")

		var courseChoice int
		fmt.Print("Enter the number of the selected course: ")
		fmt.Scanln(&courseChoice)

		studentAccount = university.NewOnlineCoursesDecorator(basicStudentAccount, courseChoice)
	case 3:
		studentAccount = university.NewLibraryAccessDecorator(basicStudentAccount, 30)
	default:
		fmt.Println("Incorrect option selection.")
		return
	}

	fmt.Println("Account description:", studentAccount.GetDescription())
	fmt.Println("Account balance:", studentAccount.GetBalance())
}
