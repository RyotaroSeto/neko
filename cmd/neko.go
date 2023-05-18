package cmd

import (
	"strings"
)

type neko struct {
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
		"        `\"\"\"\"\"\"\"\"\"\"\"\"`\n",
	}, "\n"),
}

var voiceNeko = neko{
	Simple: strings.Join([]string{
		"            /\\___/\\",
		"           ( o   o )",
		"           (  =^=  )   ニャー",
		"           (        )",
		"           (         )",
		"        ___v__^__v___",
		"       /             \\",
		"      |               |",
		"      |               |",
		"       \\             /",
		"        `\"\"\"\"\"\"\"\"\"\"\"\"`\n",
	}, "\n"),
}

func NamedNeko(name string) string {
	simpleNeko = neko{
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
			"この猫の名前は" + name,
			"名前を呼んであげてね",
		}, "\n"),
	}

	return simpleNeko.Simple
}
