package domain

const (
	EnvironmentProduction  = "production"
	EnvironmentDevelopment = "development"
	EnvironmentTesting     = "testing"
)

// IsProduction checks if the environment is production
func IsProduction(env string) bool {
	return env == EnvironmentProduction
}

// IsDevelopment checks if the environment is development
func IsDevelopment(env string) bool {
	return env == EnvironmentDevelopment
}

// IsTesting checks if the environment is testing
func IsTesting(env string) bool {
	return env == EnvironmentTesting
}
