package utils

import "strings"

// CleanUserAvatar cleans the UserAvatar string and returns a valid URL
func CleanUserAvatar(avatar string) string {
	// If UserAvatar is falsy, return the HD default gravatar image
	if avatar == "" {
		return "assets/images/DefaultGravatar.png"
	}

	// If UserAvatar includes default gravatar string or weird string, return the HD default gravatar image
	if strings.Contains(avatar, "default-avatar-36") || strings.Contains(avatar, "None?") {
		return "assets/images/DefaultGravatar.png"
	}

	// If UserAvatar is a broken link, replace and return. This has been seen in XB images
	if strings.Contains(avatar, "-ssl-ssl") {
		return strings.ReplaceAll(avatar, "-ssl-ssl", "-ssl")
	}

	return avatar
}
