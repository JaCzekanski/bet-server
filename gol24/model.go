package gol24

////////

type score struct {
	Team  int     `json:"team_id"`
	Score *string `json:"score"`
}

type data struct {
	ID          int     `json:"id"`
	LeagueID    int     `json:"league_id"`
	EventStatus string  `json:"event_status"`
	Date        string  `json:"date"`
	Scores      []score `json:"scores"`
}

/////////

type team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type league struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type refs struct {
	Teams   map[string]team   `json:"teams"`
	Leagues map[string]league `json:"leagues"`
}

// Top level object
type response struct {
	Data []data `json:"data"`
	Refs refs   `json:"refs"`
}
