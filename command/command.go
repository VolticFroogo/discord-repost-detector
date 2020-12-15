package command

import (
	"fmt"
	"github.com/VolticFroogo/discord-repost-detector/model"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Command struct {
	Name          string
	Description   string
	Format        string
	Example       string
	Triggers      []string
	Exec          func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
	RequiresAdmin bool
}

func (command *Command) NameWithAdmin() string {
	if command.RequiresAdmin {
		return command.Name + " (admin)"
	}

	return command.Name
}

func (command *Command) FormattedDescription() string {
	return fmt.Sprintf("%s\n\nFormat: %s %s %s\nExample: %s %s %s", command.Description,
		model.Ping, command.Triggers[0], command.Format,
		model.Ping, command.Triggers[0], command.Example)
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
			Format:        "",
			Example:       "",
			Exec:          help,
			RequiresAdmin: false,
		},
		{
			Name:        "Commands",
			Description: "List all available commands.",
			Triggers: []string{
				"commands",
				"list",
			},
			Format:        "",
			Example:       "",
			Exec:          list,
			RequiresAdmin: false,
		},
		{
			Name:        "Channel",
			Description: "Add or remove channels for the bot to detect reposts in.",
			Triggers: []string{
				"channel",
				"chan",
				"chat",
			},
			Format:        "[add/remove]",
			Example:       "add",
			Exec:          channel,
			RequiresAdmin: true,
		},
	}

	triggerMap = make(map[string]*Command)

	for i := range commands {
		for _, trigger := range commands[i].Triggers {
			triggerMap[trigger] = &commands[i]
		}
	}
}

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(strings.ToLower(m.Content), " ")

	if len(args) == 1 {
		help(s, m, args)
		return
	}

	if command, ok := triggerMap[args[1]]; ok {
		if command.RequiresAdmin {
			admin := hasAdministrator(s, m, args)
			if !admin {
				return
			}
		}

		command.Exec(s, m, args)
	} else {
		unknownCommand(s, m, args)
	}
}
