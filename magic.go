package Test

type Config interface {
	Get() string
}

func GetInstance(profile string) Config {
	return &config{Profile: profile}
}
