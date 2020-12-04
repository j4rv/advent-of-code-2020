package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var passports []map[string]string
	rawPassports := strings.Split(data, "\n\n")
	for _, rp := range rawPassports {
		passport := make(map[string]string)
		rp = strings.ReplaceAll(rp, "\n", " ")
		rawKeyValues := strings.Split(rp, " ")
		for _, rkv := range rawKeyValues {
			keyValue := strings.Split(rkv, ":")
			passport[keyValue[0]] = keyValue[1]
		}
		passports = append(passports, passport)
	}

	// Part One
	var validCount1 int
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, pp := range passports {
		valid := true
		for _, requiredField := range requiredFields {
			if pp[requiredField] == "" {
				valid = false
				break
			}
		}
		if valid {
			validCount1++
		}
	}
	log.Println("Part One solution:", validCount1)

	// Part Two
	var validCount2 int
	for _, pp := range passports {
		if !validateIntRange(pp["byr"], 1920, 2002+1) {
			continue
		}
		if !validateIntRange(pp["iyr"], 2010, 2020+1) {
			continue
		}
		if !validateIntRange(pp["eyr"], 2020, 2030+1) {
			continue
		}
		if !validateHeight(pp["hgt"]) {
			continue
		}
		if !validateHairColor(pp["hcl"]) {
			continue
		}
		if !validateEyeColor(pp["ecl"]) {
			continue
		}
		if !validatePassportID(pp["pid"]) {
			continue
		}
		validCount2++
	}
	log.Println("Part Two solution:", validCount2)
}

// validateIntRange checks that value is a number between min and max, min inclusive, max exclusive
func validateIntRange(value string, min, max int64) bool {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false
	}
	return min <= val && val < max
}

/*
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
*/
func validateHeight(hgt string) bool {
	if len(hgt) < 2 {
		return false
	}

	switch hgt[len(hgt)-2:] {
	case "cm":
		if !validateIntRange(hgt[:len(hgt)-2], 150, 193+1) {
			return false
		}
	case "in":
		if !validateIntRange(hgt[:len(hgt)-2], 59, 76+1) {
			return false
		}
	default:
		return false
	}

	return true
}

var hairColorRgx *regexp.Regexp = regexp.MustCompile("^#[0-9a-z]{6}$")

func validateHairColor(hcl string) bool {
	return hairColorRgx.MatchString(hcl)
}

var validEyeColors []string = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validateEyeColor(ecl string) bool {
	for _, vec := range validEyeColors {
		if vec == ecl {
			return true
		}
	}
	return false
}

var passportIDRgx *regexp.Regexp = regexp.MustCompile("^[0-9]{9}$")

func validatePassportID(pid string) bool {
	return passportIDRgx.MatchString(pid)
}
