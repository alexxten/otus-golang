package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = defineEnv(env)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
		return EnvdirReturnCode
	}
	return 0
}

func defineEnv(env Environment) []string {
	initialEnv := os.Environ()
	newEnv := make(Environment)
	for _, v := range initialEnv {
		envar := strings.SplitN(v, "=", 2)
		if x, found := env[envar[0]]; found {
			if !x.NeedRemove {
				newEnv[envar[0]] = EnvValue{Value: x.Value}
			}
		} else {
			newEnv[envar[0]] = EnvValue{Value: envar[1]}
		}
	}
	for key, v := range env {
		if _, found := newEnv[key]; !found {
			newEnv[key] = v
		}
	}

	newEnvStr := make([]string, 0)
	for key, v := range newEnv {
		newEnvStr = append(newEnvStr, key+"="+v.Value)
	}

	return newEnvStr
}
