package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	MONGO_SCOPE             = "MONGODB"
	CREATE_CLIENT_ERROR     = "failed to create and initialize new client"
	VERIFY_CONNECTION_ERROR = "failed to verify that client can reach deployment"
	CONNECT_DB_SUCCESS      = "successfully established connection to server"
	DISCONNECT_DB_ERROR     = "failed to gracefully disconnect from server"
	DISCONNECT_DB_SUCCESS   = "successfully disconnected from server"

	USER_SCOPE             = "USER_OBJECT"
	PASSWORD_HASHING_ERROR = "failed to hash/salt provided plaintext password"
)

func NewError(scope, msg string, err error) error {
	format := "%s %s %s\n%s"
	class := color.RedString("[class:ERROR]")
	scope = color.BlackString("[scope:%s]", scope)
	errMsg := wrapErrMsgAtLength(err.Error(), 80)
	return fmt.Errorf(format, class, scope, msg, errMsg)
}

func LogInfo(scope, msg string) {
	format := "%s %s %s\n"
	class := color.YellowString("[class:INFO]")
	scope = color.BlackString("[scope:%s]", scope)
	fmt.Printf(format, class, scope, msg)
}

func LogSuccess(scope, msg string) {
	format := "%s %s %s\n"
	class := color.GreenString("[class:SUCCESS]")
	scope = color.BlackString("[scope:%s]", scope)
	fmt.Printf(format, class, scope, msg)
}

func wrapErrMsgAtLength(errMsg string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(errMsg))
	wrappedText := fmt.Sprintf("  → %s", words[0])
	remainingSpace := lineWidth - len(wrappedText)
	for _, currWord := range words[1:] {
		if len(currWord)+1 > remainingSpace {
			wrappedText += fmt.Sprintf("\n    %s", currWord)
			remainingSpace = lineWidth - len(currWord)
		} else {
			wrappedText += fmt.Sprintf(" %s", currWord)
			remainingSpace -= len(currWord) + 1
		}
	}
	return wrappedText
}
