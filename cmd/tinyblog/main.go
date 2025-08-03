package main

import (
	"github.com/thoseJanes/tinyblog/internal/pkg/log"
	"github.com/thoseJanes/tinyblog/internal/tinyblog"
)

func main() {
	command := tinyblog.NewTinyBlogCommand()
	if err := command.Execute(); err != nil {
		log.Fatalw("Error in command", "err", err)
	}
}
