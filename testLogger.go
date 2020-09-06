package cryptoGo

import "fmt"

const LogEnabled = false

func LogStatement(format string, a ...interface{}) (n int, err error){
	if LogEnabled {
		return fmt.Printf(format, a ...)
	}
	return 0, nil
}

func LogError(errToLog error) (n int, err error){
	if LogEnabled {
		return fmt.Println(errToLog)
	}
	return 0, nil
}
