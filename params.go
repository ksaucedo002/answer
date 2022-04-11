package answer

type Params struct {
	offset  int
	limit   int
	Preload map[string]struct{}
}

func NewParams(offset, limit int, prelaad map[string]struct{}) Params {
	return Params{offset, limit, prelaad}
}
func (p *Params) GetPreloads() []string {
	var prels []string
	for k := range p.Preload {
		prels = append(prels, k)
	}
	return prels
}
func (p *Params) CheckPreload(k string) bool {
	_, ok := p.Preload[k]
	return ok
}
func (p *Params) GetLimit() int  { return p.limit }
func (p *Params) GetOffSet() int { return p.offset }
