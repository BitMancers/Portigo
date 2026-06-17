package utils

import (
	"math/rand"
	"os/exec"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func IsZFS() bool {
	cmd := exec.Command("zfs", "list")
	return cmd.Run() == nil
}

func MatchExtension(filename, extension string) bool {
	split := strings.Split(filename, ".")
	extPos := len(split)
	actualExtension := split[extPos-1]
	return actualExtension == extension
}
