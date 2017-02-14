package models

/*
for Day/night cycle
and which day it is

How to track and update. All here.

Need a save that runs an update on each

No weather(yet), just:
Twilight 0400 to 0600
Day 0600 to 2000
Twilight 2000 to 2200
Night 2200 to 0400
https://www.timeanddate.com/sun/usa/albuquerque

*/

var daycount = 1      //first day
var currenttime = 900 //24 time, starts at 900 on the first day

//GetTime returns time
func GetTime() (int, string) {
	solar := "Day"
	switch {
	case currenttime < 400:
		solar = "Night"
	case currenttime < 600:
		solar = "Twilight"
	case currenttime < 2000:
		solar = "Day"
	case currenttime < 2200:
		solar = "Twilight"
	case currenttime < 2400:
		solar = "Night"
	}
	return currenttime, solar
}

//CalendarCheck returns day count
func CalendarCheck() int {
	return daycount
}

//UpdateTime updates the time per input
func UpdateTime(dt int) {
	currenttime = currenttime + dt
	if currenttime >= 2400 {
		currenttime = currenttime - 2400
		daycount = daycount + 1
	}
}

//SetTime sets the time to the time provided
func SetTime(ct int) {
	currenttime = ct
}

//UpdateDay updates the day per input
func UpdateDay(dd int) {
	daycount = daycount + dd
}
