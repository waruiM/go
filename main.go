package main

import (
	"fmt"
	"project-go/helper"
	"time"
	// "strconv"
	

)
// creating package level variables
var proj_name = "Gogo mike"
const proj_year= 2023
const proj_total=13
const conc_tick uint =30
var rem_tick uint=30

//// creating an empty list of maps
// var bookings = make([]map[string]string,0)

// creating a wait group

var bookings =make([]Userinf,0)

// creating a struct
type Userinf struct{
	userName string
	userEmail string
	numberOfTickets uint
}

func main()  {

	//calling the functions created 
	startUp()
	welsec()
	

	//creating an endless loop
	for{
		userName,user_mail,userTickets := userin()

		// calling and providing return values
		isValidname,isValidmail,isValidtick:=helper.Validation(userName,user_mail,userTickets,rem_tick)
		
		// The send mail command 
		go sendEmail(userName,userTickets,proj_name,user_mail)

		if isValidmail && isValidname && isValidtick {
			// fmt.Printf("User %v has purchased %v tickets for %v \n",userName,userTickets,proj_name)
			// fmt.Printf("Thank you %v for booking %v ticket. A payment confirmation has been sent to your email addres at %v\n",userName,proj_name,user_mail)
			


			rem_tick=rem_tick-userTickets
			
			////////// creating a user map
			// var userData = make(map[string]string)

			// userData["userName"] = userName
			// userData["userEmail"] = user_mail

			// //converting the uint into a string 
			// userData["numberOfTickets"]=strconv.FormatUint(uint64(userTickets),10)


			// bookings = append(bookings,userData)

			// fmt.Printf("the remaining tickets are: %v\n",rem_tick)

			// 	// looping the emails
			// mails := []string{}
			// for _,bookings := range bookings {
				
			// 	mails = append(mails,bookings["userEmail"])
			// }
			// fmt.Println(mails) 


		    ////creating a struct
			var userinf = Userinf{

				userName :userName,
				userEmail : user_mail,
				numberOfTickets:userTickets,
			}

			bookings= append(bookings,userinf)
			
			mails := []string{}
			for _,bookings := range bookings{
				mails = append(mails,bookings.userEmail)
			} 
			
	
			fmt.Printf("The whole slice %v\n",bookings)
			fmt.Printf("The email address of the bookes %v\n",mails)
			


			if rem_tick <= 0 {
				fmt.Println("The concert is officially sold out")
				break
			}
		}else{
			if !isValidmail{
				fmt.Println("Email provided is invalid")
			}
			if !isValidname{
				fmt.Println("User name provided is less than the set length")
			}
			if !isValidtick{
				fmt.Printf("Ticket number provided is invalid %v\n",userTickets)
			}
			fmt.Println("Please try again")
		}
		


	}
}
// creating a function
func startUp(){
	fmt.Println("The application has began and is running......")
	fmt.Printf("Generating tickets for %v concert \n", proj_name)
}
// second function
func welsec(){
	fmt.Printf("Welcome to %v this is the first of many other conserts for %v \n",proj_name,proj_year)
	fmt.Println("For",proj_year,"we expect to create more than",proj_total, "projects/conserts")
	fmt.Println("The total tickets available for the",proj_name,"concert are",conc_tick,"the remaining tickets are",rem_tick)
	fmt.Println("To get your ticket to",proj_name,"here")
}
// user input function
func userin()(string,string,uint){
	// get user input
	var userName string 
	var user_mail string
	var userTickets uint
	
	
	fmt.Printf("Enter your name:")
	fmt.Scan(&userName)

	fmt.Printf("Please provide your email address:")
	fmt.Scan(&user_mail)

	fmt.Printf("How many tickets do are you looking to purchase:")
	fmt.Scan(&userTickets)
	
	return userName,user_mail,userTickets
}
func sendEmail(userName string ,userTickets uint ,proj_name string ,user_mail string){
	time.Sleep(30*time.Second)
	fmt.Println("###########")
	fmt.Println("\n Please wait as we send your data .....")
	var tick =fmt.Sprintf("\nThank you, %v for purchasing %v,%v consert tickets \n Check your email address %v for payment receipts and tickets \n",userName,userTickets,proj_name,user_mail)
	fmt.Printf("\n Sending ticket to email %v......",user_mail)
	fmt.Printf(tick)
	fmt.Println("###########")

}
