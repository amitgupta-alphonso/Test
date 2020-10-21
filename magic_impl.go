package Test

type config struct {
	Profile string
}

func (c *config) Get() string {
	return c.Profile
}
