package model

type YuqueBookStacks struct {
	Data []data `mapstructure:"data"`
}

type data struct {
	Id    int    `mapstructure:"id"`
	Name  string `mapstructure:"name"`
	Books []book `mapstructure:"books"`
}

type book struct {
	Id          int       `mapstructure:"id"`
	Type        string    `mapstructure:"type"`
	Slug        string    `mapstructure:"slug"`
	Name        string    `mapstructure:"name"`
	Description string    `mapstructure:"description"`
	Summary     []summary `mapstructure:"summary"`
	User        user      `mapstructure:"user"`
}

type summary struct {
	Id    int    `mapstructure:"id"`
	Type  string `mapstructure:"type"`
	Title string `mapstructure:"title"`
	Slug  string `mapstructure:"slug"`
}

type user struct {
	Id    int    `mapstructure:"id"`
	Name  string `mapstructure:"name"`
	Login string `mapstructure:"login"`
}
