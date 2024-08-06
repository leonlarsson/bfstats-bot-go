package structs

type Stat struct {
	Slot  uint   `json:"slot"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Extra string `json:"extra"`
}

type BaseData struct {
	Username   string `json:"username"`
	TimePlayed string `json:"timePlayed"`
	Platform   int    `json:"platform"`
	Stats      []Stat `json:"stats"`
}

/** Example data
{
    "username": "MozzyFX",
    "timePlayed": "150 hours",
    "platform": 0,
    "stats": [
        {
            "slot": 1,
            "name": "Kills:",
            "value": "13,637",
            "extra": "Top 13%"
        },
        {
            "slot": 2,
            "name": "Deaths:",
            "value": "3,254",
            "extra": "Top 29%"
        },
        {
            "slot": 3,
            "name": "Assists:",
            "value": "9,158",
            "extra": "Top 12%"
        },
        {
            "slot": 4,
            "name": "Revives:",
            "value": "705",
            "extra": "Top 29%"
        },
        {
            "slot": 6,
            "name": "W/L Ratio:",
            "value": "61.8%",
            "extra": "Top 13%"
        }
    ]
}
*/
