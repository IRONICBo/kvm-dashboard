package utils

import "github.com/melbahja/goph"

func GetCommandOutput(conn *goph.Client, name string, args ...string) (string, error) {
	cmd, err := conn.Command(name, args...)
	if err != nil {
		return "", err
	}

	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// err = cmd.Run()
	// if err != nil {
	// 	return "", err
	// }

	return string(stdout), nil
}
