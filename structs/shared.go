package structs

type Stat struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Extra string `json:"extra"`
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
