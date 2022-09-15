package main

import (
	"flag"
	"fmt"
	"log"
)

func main () {
    timeFlag1 := "";
    timeFlag2 := "";

    flag.StringVar(&timeFlag1, "d1", "", "Provide time1 i.e. 12:24:36")
    flag.StringVar(&timeFlag2, "d2", "", "Provide time2 i.e. 10:23:45")

    flag.Parse()

    h1, m1, s1, err := parseTime(timeFlag1)
    if err != nil {
        log.Fatalln(err.Error())
    }

    h2, m2, s2, err := parseTime(timeFlag2)
    if err != nil {
        log.Fatalln(err.Error())
    }

    valueH := Abs(h2 - h1)
    valueM := Abs(m2 - m1)
    valueS := Abs(s2 - s1)

    resultM := parseIntoTwoDigits(valueM)
    resultS := parseIntoTwoDigits(valueS)

    fmt.Printf("%d:%s:%s \n", valueH, resultM, resultS)
}

