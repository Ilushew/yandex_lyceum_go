package main

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	rate := (float64(p.Goals) + float64(p.Assists)) / 2
	if p.Misses != 0{
		p.Rating = rate
	} else{
		p.Rating = rate / float64(p.Misses)
	}
} 

func NewPlayer(name string, goals, misses, assists int) Player{
	p := Player{
		Name: name,
		Goals: goals,
		Misses: misses,
		Assists: assists,
	}
	p.calculateRating()
	return p
}
