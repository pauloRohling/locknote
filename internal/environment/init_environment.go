package environment

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
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
	env.validateDefaultValues()

	zap.ReplaceGlobals(createLogger("default", env.Log.Level.ZapLevel()))
	if err := env.validateRequiredFields(); err != nil {
		zap.L().Fatal("Unable to read env.yml", zap.Error(err))
	}
}

func showBanner() {
	if file, err := os.ReadFile("./banner.txt"); err == nil {
		fmt.Println(string(file))
	}
}
