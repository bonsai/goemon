package agent

import (
	"fmt"
	"goemon/src/core/mail"
)

// AgentInfo defines the profile of an agent in the system
type AgentInfo struct {
	Name        string
	Email       string
	Description string
}

var (
	// AgentAccounts maps agent roles to their information
	AgentAccounts = map[string]AgentInfo{
		"user": {
			Name:        "User",
			Email:       "user@mail.local",
			Description: "Human user interface",
		},
		"concierge": {
			Name:        "Concierge",
			Email:       "concierge@mail.local",
			Description: "Main orchestrator and router",
		},
		"baker": {
			Name:        "Baker",
			Email:       "baker@mail.local",
			Description: "Image generation worker (SD)",
		},
		"watcher": {
			Name:        "Watcher",
			Email:       "watcher@mail.local",
			Description: "Image analysis worker (VLM)",
		},
		"singer": {
			Name:        "Singer",
			Email:       "singer@mail.local",
			Description: "Audio generation worker (TTS)",
		},
	}
)

// GetAgentEmail returns the email address for a given agent name
func GetAgentEmail(name string) string {
	if agent, ok := AgentAccounts[name]; ok {
		return agent.Email
	}
	return ""
}

// GetAttachmentDir returns the attachment storage path for an agent
func GetAttachmentDir(agentName string) string {
	return "data/attachments/" + agentName
}

// GetMailboxDir returns the mailbox folder path for an agent
func GetMailboxDir(agentName string, folder string) string {
	return "data/mailbox/" + agentName + "/" + folder
}

// MailAgent handles the agent-specific mail logic
type MailAgent struct {
	Info   AgentInfo
	Config mail.MailConfig
}

// NewMailAgent creates a new mail agent instance
func NewMailAgent(name string) (*MailAgent, error) {
	info, ok := AgentAccounts[name]
	if !ok {
		return nil, fmt.Errorf("agent %s not found", name)
	}

	config := mail.GetDefaultConfig()

	return &MailAgent{
		Info:   info,
		Config: config,
	}, nil
}
