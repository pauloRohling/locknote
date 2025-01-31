package environment

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

const profileEnv = "PROFILE"

var (
	env     = Config{}
	profile = "dev"
)

func Env() Config {
	return env
}

func Profile() string {
	return profile
}

func init() {
	showBanner()

	if envProfile := os.Getenv(profileEnv); envProfile != "" {
		profile = envProfile
	}

	fileName := "./env.yml"
	if profile != "dev" {
		fileName = fmt.Sprintf("./env-%s.yml", profile)
	}

	_ = cleanenv.ReadConfig(fileName, &env)
	_ = cleanenv.ReadConfig(".env", &env)

	if err := env.validateRequiredFields(); err != nil {
		panic("Unable to read env.yml")
	}
}

func showBanner() {
	if file, err := os.ReadFile("./banner.txt"); err == nil {
		fmt.Println(string(file))
	}
}
