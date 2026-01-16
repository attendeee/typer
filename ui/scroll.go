package ui

func ScrollUp(p *Pager) {

	if p.UpperOffsetIdx-p.OffsetStep >= 0 {
		p.UpperOffsetIdx -= p.OffsetStep
	} else {
		p.UpperOffsetIdx = 0
	}

	if p.UpperOffsetIdx+p.OffsetStep < len(p.Offsets) {
		p.BottomOffsetIdx = p.UpperOffsetIdx + p.OffsetStep
	} else {
		p.BottomOffsetIdx = p.UpperOffsetIdx + len(p.Offsets) - p.UpperOffsetIdx
	}

	p.UpperOffset = p.Offsets[p.UpperOffsetIdx]

	p.BottomOffset = p.Offsets[p.BottomOffsetIdx]

}

func ScrollDown(p *Pager) {
	// Adjust BottomOffsetIdx first
	if p.BottomOffsetIdx+p.OffsetStep < len(p.Offsets) {
		p.BottomOffsetIdx += p.OffsetStep
	} else {
		p.BottomOffsetIdx = len(p.Offsets) - 1 // Ensure it does not exceed bounds
	}

	p.UpperOffsetIdx = p.BottomOffsetIdx - p.OffsetStep

	// Ensure UpperOffsetIdx does not fall below zero
	if p.UpperOffsetIdx < 0 {
		p.UpperOffsetIdx = 0
	}

	p.UpperOffset = p.Offsets[p.UpperOffsetIdx]
	p.BottomOffset = p.Offsets[p.BottomOffsetIdx]

}

func UpdateUpperOffsetIdx(m *Model, p *Pager) {
	for i, v := range p.Offsets {
		if m.CursorPos >= v && v <= m.CursorPos {
			p.UpperOffsetIdx = i
		}
	}
}

func UpdateBottomOffsetIdx(p *Pager) {
	if p.UpperOffsetIdx+p.OffsetStep < len(p.Offsets)-1 {
		p.BottomOffsetIdx = p.UpperOffsetIdx + p.OffsetStep
	} else {
		p.BottomOffsetIdx = len(p.Offsets) - 1
	}
}

func UpdateOffsets(m *Model, p *Pager) {
	p.Offsets = []int{}
	p.Offsets = append(p.Offsets, 0)
	for i, v := range m.Text {
		if v == '\n' {
			m.Pager.Offsets = append(p.Offsets, i+1)
		}
	}
}
