package helper

import "strings"

// user validation function

func Validation(userName string, user_mail string, userTickets uint,rem_tick uint) (bool,bool,bool){
	isValidname := len(userName) > 2
	isValidmail := strings.Contains(user_mail,"@")
	isValidtick := userTickets > 0 && userTickets <= rem_tick
	return isValidname,isValidmail,isValidtick

}
