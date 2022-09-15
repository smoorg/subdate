package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Parses string in format HH:mm:SS and returns hour, minutes, and seconds as ints.
func parseTime(timeStr string) (int, int, int, error) {
    timeStrings := strings.Split(timeStr, ":")

    timeSize := len(timeStrings)
    if timeSize != 3 {
    	return 0, 0, 0, errors.New(fmt.Sprintln("Invalid time format", timeSize))
    }
    hour, err := parseTimePart(timeStrings[0])
    if err != nil {
    	return 0, 0, 0, err
    }

    minute, err := parseTimePart(timeStrings[1])
    if err != nil {
    	return 0, 0, 0, err
    }

    second, err := parseTimePart(timeStrings[2])
    if err != nil {
    	return 0, 0, 0, err
    }
 
    return hour, minute,second, nil
}

func parseTimePart(v string) (int, error) {
    res, err := strconv.Atoi(v)
    if err != nil {
    	return 0, errors.New(fmt.Sprintln("Invalid time format", v))
    }

    return res, nil
}

func parseIntoTwoDigits(v int) string {
    result := ""
    if v < 10 {
        result += "0"
    } 

    result += strconv.Itoa(v)
    return result
}
