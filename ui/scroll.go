package ui

func ScrollUp(m *Model) {

	if m.CursorPos-1 < m.UpperOffset {
		if m.UpperOffsetIdx-m.OffsetStep >= 0 {
			m.UpperOffsetIdx -= m.OffsetStep
		} else {
			m.UpperOffsetIdx = 0
		}

		if m.UpperOffsetIdx+m.OffsetStep < len(m.Offsets) {
			m.BottomOffsetIdx = m.UpperOffsetIdx + m.OffsetStep
		} else {
			m.BottomOffsetIdx = m.UpperOffsetIdx + len(m.Offsets) - m.UpperOffsetIdx
		}

		m.UpperOffset = m.Offsets[m.UpperOffsetIdx]

		m.BottomOffset = m.Offsets[m.BottomOffsetIdx]
		m.CursorPos = m.BottomOffset - 1

	}
}

func ScrollDown(m *Model) {
	if m.CursorPos+1 > m.BottomOffset {
		// Adjust BottomOffsetIdx first
		if m.BottomOffsetIdx+m.OffsetStep < len(m.Offsets) {
			m.BottomOffsetIdx += m.OffsetStep
		} else {
			m.BottomOffsetIdx = len(m.Offsets) - 1 // Ensure it does not exceed bounds
		}

		m.UpperOffsetIdx = m.BottomOffsetIdx - m.OffsetStep

		// Ensure UpperOffsetIdx does not fall below zero
		if m.UpperOffsetIdx < 0 {
			m.UpperOffsetIdx = 0
		}

		m.UpperOffset = m.Offsets[m.UpperOffsetIdx]
		m.BottomOffset = m.Offsets[m.BottomOffsetIdx]
	}
}

func UpdateUpperOffsetIdx(m *Model) {
	for i, v := range m.Offsets {
		if m.CursorPos >= v && v <= m.CursorPos {
			m.UpperOffsetIdx = i
		}
	}
}

func UpdateBottomOffsetIdx(m *Model) {
	if m.UpperOffsetIdx+m.OffsetStep < len(m.Offsets)-1 {
		m.BottomOffsetIdx = m.UpperOffsetIdx + m.OffsetStep
	} else {
		m.BottomOffsetIdx = len(m.Offsets) - 1
	}
}

func UpdateOffsets(m *Model) {
	m.Offsets = append(m.Offsets, 0)
	for i, v := range m.Text {
		if v == '\n' {
			m.Offsets = append(m.Offsets, i+1)
		}
	}
}
