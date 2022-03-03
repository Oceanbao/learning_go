package main

/*
* Duration - int164 nanosecond
* 1e10 == 1 sec or 1000 ms
* 1h = 1*60*60*10e11
*
 */

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	newT := t.Add(time.Hour * 1)
	fmt.Printf("Add 1H\n: %s\n", newT)

	newT = t.Add(time.Minute * 15)
	fmt.Printf("Add 15m\n: %s\n", newT)

	newT = t.Add(time.Second * 10)
	fmt.Printf("Add 10s\n: %s\n", newT)

	newT = t.Add(time.Millisecond * 100)
	fmt.Printf("Adding 100 millisecond\n: %s\n", newT)

	newT = t.Add(time.Nanosecond * 1000)
	fmt.Printf("Adding 1000 nanosecond\n: %s\n", newT)

	newT = t.AddDate(1, 2, 4) // YMD
	fmt.Printf("Add 1Y 2M 4D\n: %s\n", newT)

	//Parse YYYY-MM-DD
	timeT, _ := time.Parse("2006-01-02", "2020-01-29")
	fmt.Println(timeT)

	//Parse YY-MM-DD
	timeT, _ = time.Parse("06-01-02", "20-01-29")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD
	timeT, _ = time.Parse("2006-Jan-02", "2020-Jan-29")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD WeekDay HH:MM:SS
	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05", "2020-Jan-29 Wednesday 12:19:25")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset
	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05 PM MST -07:00", "2020-Jan-29 Wednesday 12:19:25 AM IST +05:30")
	fmt.Println(timeT)

	now = time.Now()

	//Format YYYY-MM-DD
	fmt.Printf("YYYY-MM-DD: %s\n", now.Format("2006-01-02"))

	//Format YY-MM-DD
	fmt.Printf("YY-MM-DD: %s\n", now.Format("06-01-02"))

	//Format YYYY-#{MonthName}-DD
	fmt.Printf("YYYY-#{MonthName}-DD: %s\n", now.Format("2006-Jan-02"))

	//Format HH:MM:SS
	fmt.Printf("HH:MM:SS: %s\n", now.Format("03:04:05"))

	//Format HH:MM:SS Millisecond
	fmt.Printf("HH:MM:SS Millisecond: %s\n", now.Format("03:04:05 .999"))

	//Format YYYY-#{MonthName}-DD WeekDay HH:MM:SS
	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS: %s\n", now.Format("2006-Jan-02 Monday 03:04:05"))

	//Format YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset
	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset: %s\n", now.Format("2006-Jan-02 Monday 03:04:05 PM MST -07:00"))

	currentTime := time.Now()
	oldTime := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	diff := currentTime.Sub(oldTime)
	fmt.Println("diff: ", diff)

	tNow := time.Now()

	//time.Time to Unix Timestamp
	tUnix := tNow.Unix()
	fmt.Printf("timeUnix %d\n", tUnix)

	//Unix Timestamp to time.Time
	timeT := time.Unix(tUnix, 0)
	fmt.Printf("time.Time: %s\n", timeT)

	now = time.Now()

	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("UTC Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Europe/Berlin")
	fmt.Printf("Berlin Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Dubai")
	fmt.Printf("Dubai Time: %s\n", now.In(loc))
}
