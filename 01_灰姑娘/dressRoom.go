package main

// DressRoom ...
type DressRoom struct {
	Dresses []*Dress
}

// NewDressRoom ...
func NewDressRoom() *DressRoom {
	return &DressRoom{}
}

// Store -- 保管连衣裙
func (d *DressRoom) Store(actors ...*Actor) {
	for _, a := range actors {
		cos := NewDress(a)
		if dress, ok := cos.(*Dress); ok {
			d.Dresses = append(d.Dresses, dress)
		}
	}
}

// GetDress -- 获取连衣裙
func (d *DressRoom) GetDress(a *Actor) {
	for _, dress := range d.Dresses {
		if dress.Wear(a) {
			a.SetCostume(dress)
		}
	}
}
