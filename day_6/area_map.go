package main

type AreaMap struct {
	Contents [][]rune
}

func (m *AreaMap) Dimensions() [2]int {
	return [2]int{len(m.Contents), len(m.Contents[0])}
}
