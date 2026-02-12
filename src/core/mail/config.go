package mail

type MailConfig struct {
	Server   string
	Port     int
	User     string
	Password string
}

func GetDefaultConfig() MailConfig {
	return MailConfig{
		Server: "localhost",
		Port:   993,
		User:   "user@mail.local",
		Password: "password",
	}
}
