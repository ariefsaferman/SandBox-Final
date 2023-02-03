package helpers

import "strconv"

func GenerateNewWalletNumber(id int) (int, error) {
	var idString string
	startingString := "777"
	if id < 10 {
		idString = "00" + strconv.Itoa(id)
	} else if id < 100 {
		idString = "0" + strconv.Itoa(id)
	} else {
		idString = strconv.Itoa(id)
	}
	resultString := startingString + idString
	resultInt, err := strconv.Atoi(resultString)
	if err != nil {
		return 0, err
	}

	return resultInt, nil
}