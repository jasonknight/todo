package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"strings"
)

func make_cmd_here() *commander.Command {
	cmd_add := func(cmd *commander.Command, args []string) error {
		
		filename,_ := os.Getwd();
		filename += "/.todo";
		w, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		defer w.Close()
		if len(args) == 0 {
			return nil
		}
		_, err = fmt.Fprintf(w, " %s\n", strings.Join(args, " "))
		return err
	}

	return &commander.Command{
		Run:       cmd_add,
		UsageLine: "here [message]",
		Short:     "",
	}
}
