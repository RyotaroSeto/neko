package cmd

import "strings"

type neko struct {
	Move   []string
	Simple string
}

var simpleNeko = neko{
	Simple: strings.Join([]string{
		"            /\\___/\\",
		"           ( o   o )",
		"           (  =^=  )",
		"           (        )",
		"           (         )",
		"        ___v__^__v___",
		"       /             \\",
		"      |               |",
		"      |               |",
		"       \\             /",
		"        `\"\"\"\"\"\"\"\"\"\"\"\"`",
	}, "\n"),
}
