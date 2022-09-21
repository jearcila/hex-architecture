package environment

import (
	"strings"
)

const (
	EnvProduction = "production"
	EnvStaging    = "staging"
	EnvLocal      = "local"
)

func IsLocalEnvironment(env string) bool {
	return strings.Contains(env, EnvLocal)
}

func IsProdEnvironment(env string) bool {
	return strings.Contains(env, "prod")
}

func IsStagingEnvironment(env string) bool {
	return strings.Contains(env, EnvStaging)
}

func GetEnv(env string) string {
	if IsProdEnvironment(env) {
		return EnvProduction
	}

	return env
}

type Environment string

const (
	// EnvironmentProduction defines production environment value
	EnvironmentProduction Environment = "production"
	// EnvironmentProduction defines staging environment value
	EnvironmentStaging Environment = "staging"
	// EnvironmentLocal defines local environment value
	EnvironmentLocal Environment = "local"
)
