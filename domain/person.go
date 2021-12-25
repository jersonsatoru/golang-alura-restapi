package domain

type Person struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	BackgroundStory string `json:"background_story"`
}

func NewPerson(name, backgroundStory string) *Person {
	return &Person{
		Name:            name,
		BackgroundStory: backgroundStory,
	}
}
