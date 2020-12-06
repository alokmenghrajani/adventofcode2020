package day04

import (
	"github.com/alokmenghrajani/adventofcode2020/utils"
	"regexp"
	"strings"
)

func Part1(input string) uint {
	passports := strings.Split(input, "\n\n")
	valid := uint(0)
	for _, passport := range passports {
		// easier to remove newlines than deal with multiline regexp :P
		passport = strings.ReplaceAll(passport, "\n", " ")
		current := map[string]bool{}
		re := regexp.MustCompile(`(...):([^ ]+)`)
		matches := re.FindAllStringSubmatch(passport, -1)
		for _, match := range matches {
			if match[1] == "cid" {
				continue
			}
			current[match[1]] = true
		}
		if len(current) == 7 {
			valid++
		}
	}
	return valid
}

func Part2(input string) uint {
	passports := strings.Split(input, "\n\n")
	valid := uint(0)
	for _, passport := range passports {
		// easier to remove newlines than deal with multiline regexp :P
		passport = strings.ReplaceAll(passport, "\n", " ")
		current := map[string]bool{}
		re := regexp.MustCompile(`(...):([^ ]+)`)
		matches := re.FindAllStringSubmatch(passport, -1)
		for _, match := range matches {
			key := match[1]
			val := match[2]
			if key == "cid" {
				continue
			}
			if key == "byr" {
				t := utils.Atoi(val, 0)
				if t < 1920 || t > 2002 {
					continue
				}
			}
			if key == "iyr" {
				t := utils.Atoi(val, 0)
				if t < 2010 || t > 2020 {
					continue
				}
			}
			if key == "eyr" {
				t := utils.Atoi(val, 0)
				if t < 2020 || t > 2030 {
					continue
				}
			}
			if key == "hgt" {
				if strings.HasSuffix(val, "cm") {
					t := utils.Atoi(val[:len(val)-2], 0)
					if t < 150 || t > 193 {
						continue
					}
				} else if strings.HasSuffix(val, "in") {
					t := utils.Atoi(val[:len(val)-2], 0)
					if t < 59 || t > 76 {
						continue
					}
				} else {
					continue
				}
			}
			if key == "hcl" {
				re2 := regexp.MustCompile(`^#[0-9a-f]{6}$`)
				if !re2.MatchString(val) {
					continue
				}
			}
			if key == "ecl" {
				switch val {
				case "amb",
					"blu",
					"brn",
					"gry",
					"grn",
					"hzl",
					"oth":
				default:
					continue
				}
			}
			if key == "pid" {
				re2 := regexp.MustCompile(`^[0-9]{9}$`)
				if !re2.MatchString(val) {
					continue
				}
			}

			current[key] = true
		}
		if len(current) == 7 {
			valid++
		}
	}
	return valid
}
