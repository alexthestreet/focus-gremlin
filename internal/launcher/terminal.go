package launcher

import (
	"fmt"
	"os/exec"
)

func BuildCommand(terminalCommand []string, appCommand []string) (*exec.Cmd, error) {
	if len(terminalCommand) == 0 {
		return nil, fmt.Errorf("terminal command is empty")
	}

	args := append(append([]string(nil), terminalCommand...), appCommand...)
	return exec.Command(args[0], args[1:]...), nil
}

func Launch(terminalCommand []string, appCommand []string) error {
	cmd, err := Start(terminalCommand, appCommand)
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func Start(terminalCommand []string, appCommand []string) (*exec.Cmd, error) {
	cmd, err := BuildCommand(terminalCommand, appCommand)
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}
