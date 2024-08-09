package utils

func GetAvatarImageURL(avatar *string) string {
	if avatar == nil || *avatar == "" {
		return "assets/images/DefaultGravatar.png"
	}
	return *avatar
}
