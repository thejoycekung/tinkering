package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// printError is a util function that prints the given error and exits.
func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

// convertToSymbols takes an octal number and returns the permissions as a string.
func convertToSymbols(nums int64) (string, error) {
	if nums > 777 {
		return "", errors.New("number is too large for chmod")
	}

	// convert to binary
	bins := strconv.FormatInt(nums, 2)
	if len(bins) < 9 {
		// pad it with 0s
		padding := 9 - len(bins)
		for i := 0; i < padding; i++ {
			bins = "0" + bins
		}
	}

	perms := ""
	for i := 0; i < 9; i++ {
		switch {
		case bins[i] == '0':
			perms += "-"
		case i%3 == 0:
			perms += "r"
		case i%3 == 1:
			perms += "w"
		case i%3 == 2:
			perms += "x"
		default:
			return "", errors.New("unable to parse octal")
		}
	}
	return perms, nil
}

// convertToOctal takes a string of permissions and converts it to a number.
func convertToOctal(symbs string) (int, error) {
	if len(symbs) != 9 {
		return 0, errors.New("permissions should be 9 characters long")
	}

	var tempPerms string
	for _, ch := range symbs {
		switch ch {
		case 'r', 'w', 'x':
			tempPerms += "1"
		case '-':
			tempPerms += "0"
		default:
			return 0, errors.New("unable to parse symbols")
		}
	}

	binPerms, err := strconv.ParseInt(tempPerms, 2, 0)
	if err != nil {
		return 0, err
	}
	perms, err := strconv.Atoi(strconv.FormatInt(binPerms, 8))
	if err != nil {
		return 0, err
	}
	return perms, nil
}

func main() {
	if len(os.Args) != 2 {
		printError(errors.New("uhoh"))
	}

	if modNum, err := strconv.ParseInt(os.Args[1], 8, 0); err == nil {
		fmt.Println("Number", os.Args[1], "received! Converting to symbolic format ...")
		symb, err := convertToSymbols(modNum)
		if err != nil {
			printError(err)
		}
		fmt.Println("This number represents these permissions: ", symb)
	} else {
		fmt.Println("Symbol", os.Args[1], "received! Converting to octal format ...")
		octs, err := convertToOctal(os.Args[1])
		if err != nil {
			printError(err)
		}
		fmt.Println("These permissions can be represented as this octal number: ", octs)
	}
}
