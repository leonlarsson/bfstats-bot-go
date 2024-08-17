package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

// OverviewSegment returns the overview segment choice for the stats command.
func OverviewSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/overview_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/overview_name"),
		Value: "overview",
	}
}

// WeaponsSegment returns the weapons segment choice for the stats command.
func WeaponsSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/weapons_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/weapons_name"),
		Value: "weapons",
	}
}

// VehiclesSegment returns the vehicles segment choice for the stats command.
func VehiclesSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/vehicles_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/vehicles_name"),
		Value: "vehicles",
	}
}

// ClassesSegment returns the classes segment choice for the stats command.
func ClassesSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/classes_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/classes_name"),
		Value: "classes",
	}
}

// GadgetsSegment returns the gadgets segment choice for the stats command.
func GadgetsSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/gadgets_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/gadgets_name"),
		Value: "gadgets",
	}
}

// MapsSegment returns the maps segment choice for the stats command.
func MapsSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/maps_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/maps_name"),
		Value: "maps",
	}
}

// ModesSegment returns the modes segment choice for the stats command.
func ModesSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/modes_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/modes_name"),
		Value: "modes",
	}
}

// MatchesSegment returns the matches segment choice for the stats command.
func MatchesSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/matches_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/matches_name"),
		Value: "matches",
	}
}

// FirestormSegment returns the firestorm segment choice for the stats command.
func FirestormSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/firestorm_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/firestorm_name"),
		Value: "firestorm",
	}
}

// HazardZoneSegment returns the hazard zone segment choice for the stats command.
func HazardZoneSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/hazardzone_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/hazardzone_name"),
		Value: "hazardzone",
	}
}

// ProgressionSegment returns the progression segment choice for the stats command.
func ProgressionSegment() discord.ApplicationCommandOptionChoiceString {
	return discord.ApplicationCommandOptionChoiceString{
		Name: localization.GetEnglishString("slash_commands/stats/options/segment/progression_name"),
		// NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/progression_name"),
		Value: "progression",
	}
}
