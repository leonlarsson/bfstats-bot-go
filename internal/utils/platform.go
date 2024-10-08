package utils

func TRNPlatformNameToInt(platform string) int {
	switch platform {
	case "origin":
		return 0
	case "psn":
		return 1
	case "xbl":
		return 2
	default:
		return 0
	}
}

func PlatformIntToTRNName(platform int) string {
	switch platform {
	case 0:
		return "origin"
	case 1:
		return "psn"
	case 2:
		return "xbl"
	default:
		return "origin"
	}
}
