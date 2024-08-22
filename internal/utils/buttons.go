package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

func BuildErrorStatsLinkButton(statsLink string) discord.ButtonComponent {
	return discord.ButtonComponent{
		Style: discord.ButtonStyleLink,
		Label: Ternary(strings.HasPrefix(statsLink, "https://battlefieldtracker.com"), "Tracker Network", "Game Tools").(string),
		URL:   statsLink,
	}
}

func BuildInviteButton(loc localization.LanguageLocalizer) discord.ButtonComponent {
	return discord.ButtonComponent{
		Style: discord.ButtonStyleLink,
		Label: loc.Translate("messages/button_add_bot"),
		URL:   fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=%s", os.Getenv("BOT_ID")),
	}
}

func BuildWebsiteButton(loc localization.LanguageLocalizer) discord.ButtonComponent {
	return discord.ButtonComponent{
		Style: discord.ButtonStyleLink,
		Label: loc.Translate("messages/button_website"),
		URL:   "https://battlefieldstats.co",
	}
}

func BuildFullStatsButton(statsLink string, loc localization.LanguageLocalizer) discord.ButtonComponent {
	return discord.ButtonComponent{
		Style:    discord.ButtonStyleLink,
		Label:    loc.Translate("messages/button_full_stats"),
		URL:      statsLink,
		Disabled: statsLink == "",
	}
}

func BuildDeleteButton(loc localization.LanguageLocalizer) discord.ButtonComponent {
	return discord.ButtonComponent{
		Style:    discord.ButtonStyleDanger,
		Label:    loc.Translate("messages/button_delete"),
		CustomID: "delete-stats",
	}
}

func BuildShowImageButton(loc localization.LanguageLocalizer) discord.ButtonComponent {
	return discord.ButtonComponent{
		Style:    discord.ButtonStylePrimary,
		Label:    loc.Translate("messages/button_show_image"),
		CustomID: "show-image",
	}
}
