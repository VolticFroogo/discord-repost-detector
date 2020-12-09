package command

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Format      string
	Example     string
	Triggers    []string
	Exec        func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (err error)
}

var (
	commands   []Command
	triggerMap map[string]*Command
)

func RegisterCommands() {
	commands = []Command{
		{
			Name:        "Help",
			Description: "Get help on how to use Repost Detector.",
			Triggers: []string{
				"help",
			},
			Format:  "",
			Example: "",
			Exec:    help,
		},
		{
			Name:        "Commands",
			Description: "List all available commands.",
			Triggers: []string{
				"commands",
				"list",
			},
			Format:  "",
			Example: "",
			Exec:    list,
		},
	}

	triggerMap = make(map[string]*Command)

	for i := range commands {
		for _, trigger := range commands[i].Triggers {
			triggerMap[trigger] = &commands[i]
		}
	}
}

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	split := strings.Split(strings.ToLower(m.Content), " ")

	if len(split) == 1 {
		err = help(s, m, split)
		return
	}

	if command, ok := triggerMap[split[1]]; ok {
		err = command.Exec(s, m, split)
	} else {
		err = unknownCommand(s, m, split)
	}

	return
}
