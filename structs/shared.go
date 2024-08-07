package structs

// BaseData represents the base data for all image creations.
// Each game embeds BaseData together with its own GAMEStats struct.
type BaseData struct {
	Identifier string `json:"identifier"`
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
	Platform   int    `json:"platform"`
	Meta       Meta   `json:"meta"`
}

type Meta struct {
	Game    string `json:"game"`
	Segment string `json:"segment"`
}

// Stat represents a single stat with a name, value and extra information.
type Stat struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Extra string `json:"extra"`
}

// RankStat represents a rank stat with a name, value, rankInt and extra information.
type RankStat struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	RankInt int    `json:"rankInt"`
	Extra   string `json:"extra"`
}

/** Example data
{
  "identifier": "BF2042-001",
  "username": "MozzyFX",
  "timePlayed": "150 hours",
  "platform": 0,
  "stats": {
    "kills": {
      "name": "Kills:",
      "value": "13,637",
      "extra": "Top 13%"
    },
    "deaths": {
      "name": "Deaths:",
      "value": "3,254",
      "extra": "Top 29%"
    },
    "assists": {
      "name": "Assists:",
      "value": "9,158",
      "extra": "Top 12%"
    },
    "revives": {
      "name": "Revives:",
      "value": "705",
      "extra": "Top 29%"
    },
    "wlRatio": {
      "name": "W/L Ratio:",
      "value": "61.8%",
      "extra": "Top 13%"
    },
    "bestClass": {
      "name": "Best Class:",
      "value": "Angel",
      "extra": "2,813 kills | 15 hours"
    },
    "killsPerMatch": {
      "name": "Kills/Match:",
      "value": "32.55",
      "extra": "Top 7%"
    },
    "kdRatio": {
      "name": "K/D Ratio:",
      "value": "4.19 (3.4)",
      "extra": "Top 3.3%"
    },
    "killsPerMinute": {
      "name": "Kills/Minute:",
      "value": "1.51",
      "extra": "Top 10%"
    },
    "rank": {
      "name": "Rank 114 (S015)",
      "value": "96% to next rank",
      "extra": "XP: 7,586,196"
    }
  }
}
*/
