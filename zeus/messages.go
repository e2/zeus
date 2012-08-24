package zeus

import (
	"strings"
	"strconv"
	"errors"
)

func parsePidMessage(msg string) (int, string, error) {
	parts := strings.SplitN(msg, ":", 3)
	if parts[0] != "P" {
		return -1, "", errors.New("Wrong message type!")
	}

	identifier := parts[2]
	pid, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, "", err
	}

	return pid, identifier, nil
}

func createActionMessage(action string) (string) {
	return "A:" + action
}

func parseActionResponseMessage(msg string) (string, error) {
	parts := strings.SplitN(msg, ":", 2)
	if parts[0] != "R" {
		return "", errors.New("Wrong message type!")
	}
	return parts[1], nil
}

func createSpawnSlaveMessage(identifier string) (string) {
	return "S:" + identifier
}

func createSpawnCommandMessage(identifier string) (string) {
	return "C:" + identifier
}