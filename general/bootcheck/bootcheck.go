package bootcheck

func EnvironmentOK(version string) bool {
	return version != ""
}
