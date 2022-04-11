package answer

type Parmas struct {
	offset  int
	limit   int
	Preload map[string]struct{}
}

func NewParams(offset, limit int, prelaad map[string]struct{}) Parmas {
	return Parmas{offset, limit, prelaad}
}
func (p *Parmas) GetPreloads() []string {
	var prels []string
	for k := range p.Preload {
		prels = append(prels, k)
	}
	return prels
}
func (p *Parmas) GetLimit() int  { return p.limit }
func (p *Parmas) GetOffSet() int { return p.offset }
