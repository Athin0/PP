// Package slot is generated by gogll. Do not edit.
package slot

import (
	"bytes"
	"fmt"

	"PP/app/internal/grammar/parser/symbols"
)

type Label int

const (
	AddSub0R0 Label = iota
	AddSub0R1
	AddSub1R0
	AddSub1R1
	AddSubCont0R0
	AddSubCont0R1
	AddSubCont1R0
	AddSubCont1R1
	AddSubCont1R2
	AddSubCont1R3
	MinMax0R0
	MinMax0R1
	MinMax1R0
	MinMax1R1
	MinMaxCont0R0
	MinMaxCont0R1
	MinMaxCont1R0
	MinMaxCont1R1
	MinMaxCont1R2
	MinMaxCont1R3
	MulDiv0R0
	MulDiv0R1
	MulDiv1R0
	MulDiv1R1
	MulDivCont0R0
	MulDivCont0R1
	MulDivCont1R0
	MulDivCont1R1
	MulDivCont1R2
	MulDivCont1R3
	Start0R0
	Start0R1
	Start1R0
	Start1R1
	Start1R2
	Start1R3
	Unary0R0
	Unary0R1
	Unary0R2
	Unary0R3
	Unary1R0
	Unary1R1
	Unary1R2
	Unary1R3
	Unary2R0
	Unary2R1
	Unary2R2
	Unary2R3
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
	Label   Label
}

type Index struct {
	NT  symbols.NT
	Alt int
	Pos int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt, alt, pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{
	AddSub0R0: {
		symbols.NT_AddSub, 0, 0,
		symbols.Symbols{
			symbols.T_5,
		},
		AddSub0R0,
	},
	AddSub0R1: {
		symbols.NT_AddSub, 0, 1,
		symbols.Symbols{
			symbols.T_5,
		},
		AddSub0R1,
	},
	AddSub1R0: {
		symbols.NT_AddSub, 1, 0,
		symbols.Symbols{
			symbols.T_4,
		},
		AddSub1R0,
	},
	AddSub1R1: {
		symbols.NT_AddSub, 1, 1,
		symbols.Symbols{
			symbols.T_4,
		},
		AddSub1R1,
	},
	AddSubCont0R0: {
		symbols.NT_AddSubCont, 0, 0,
		symbols.Symbols{
			symbols.NT_MulDivCont,
		},
		AddSubCont0R0,
	},
	AddSubCont0R1: {
		symbols.NT_AddSubCont, 0, 1,
		symbols.Symbols{
			symbols.NT_MulDivCont,
		},
		AddSubCont0R1,
	},
	AddSubCont1R0: {
		symbols.NT_AddSubCont, 1, 0,
		symbols.Symbols{
			symbols.NT_AddSubCont,
			symbols.NT_MulDiv,
			symbols.NT_MulDivCont,
		},
		AddSubCont1R0,
	},
	AddSubCont1R1: {
		symbols.NT_AddSubCont, 1, 1,
		symbols.Symbols{
			symbols.NT_AddSubCont,
			symbols.NT_MulDiv,
			symbols.NT_MulDivCont,
		},
		AddSubCont1R1,
	},
	AddSubCont1R2: {
		symbols.NT_AddSubCont, 1, 2,
		symbols.Symbols{
			symbols.NT_AddSubCont,
			symbols.NT_MulDiv,
			symbols.NT_MulDivCont,
		},
		AddSubCont1R2,
	},
	AddSubCont1R3: {
		symbols.NT_AddSubCont, 1, 3,
		symbols.Symbols{
			symbols.NT_AddSubCont,
			symbols.NT_MulDiv,
			symbols.NT_MulDivCont,
		},
		AddSubCont1R3,
	},
	MinMax0R0: {
		symbols.NT_MinMax, 0, 0,
		symbols.Symbols{
			symbols.T_0,
		},
		MinMax0R0,
	},
	MinMax0R1: {
		symbols.NT_MinMax, 0, 1,
		symbols.Symbols{
			symbols.T_0,
		},
		MinMax0R1,
	},
	MinMax1R0: {
		symbols.NT_MinMax, 1, 0,
		symbols.Symbols{
			symbols.T_9,
		},
		MinMax1R0,
	},
	MinMax1R1: {
		symbols.NT_MinMax, 1, 1,
		symbols.Symbols{
			symbols.T_9,
		},
		MinMax1R1,
	},
	MinMaxCont0R0: {
		symbols.NT_MinMaxCont, 0, 0,
		symbols.Symbols{
			symbols.NT_Unary,
		},
		MinMaxCont0R0,
	},
	MinMaxCont0R1: {
		symbols.NT_MinMaxCont, 0, 1,
		symbols.Symbols{
			symbols.NT_Unary,
		},
		MinMaxCont0R1,
	},
	MinMaxCont1R0: {
		symbols.NT_MinMaxCont, 1, 0,
		symbols.Symbols{
			symbols.T_1,
			symbols.NT_Start,
			symbols.T_2,
		},
		MinMaxCont1R0,
	},
	MinMaxCont1R1: {
		symbols.NT_MinMaxCont, 1, 1,
		symbols.Symbols{
			symbols.T_1,
			symbols.NT_Start,
			symbols.T_2,
		},
		MinMaxCont1R1,
	},
	MinMaxCont1R2: {
		symbols.NT_MinMaxCont, 1, 2,
		symbols.Symbols{
			symbols.T_1,
			symbols.NT_Start,
			symbols.T_2,
		},
		MinMaxCont1R2,
	},
	MinMaxCont1R3: {
		symbols.NT_MinMaxCont, 1, 3,
		symbols.Symbols{
			symbols.T_1,
			symbols.NT_Start,
			symbols.T_2,
		},
		MinMaxCont1R3,
	},
	MulDiv0R0: {
		symbols.NT_MulDiv, 0, 0,
		symbols.Symbols{
			symbols.T_3,
		},
		MulDiv0R0,
	},
	MulDiv0R1: {
		symbols.NT_MulDiv, 0, 1,
		symbols.Symbols{
			symbols.T_3,
		},
		MulDiv0R1,
	},
	MulDiv1R0: {
		symbols.NT_MulDiv, 1, 0,
		symbols.Symbols{
			symbols.T_6,
		},
		MulDiv1R0,
	},
	MulDiv1R1: {
		symbols.NT_MulDiv, 1, 1,
		symbols.Symbols{
			symbols.T_6,
		},
		MulDiv1R1,
	},
	MulDivCont0R0: {
		symbols.NT_MulDivCont, 0, 0,
		symbols.Symbols{
			symbols.NT_MinMaxCont,
		},
		MulDivCont0R0,
	},
	MulDivCont0R1: {
		symbols.NT_MulDivCont, 0, 1,
		symbols.Symbols{
			symbols.NT_MinMaxCont,
		},
		MulDivCont0R1,
	},
	MulDivCont1R0: {
		symbols.NT_MulDivCont, 1, 0,
		symbols.Symbols{
			symbols.NT_MulDivCont,
			symbols.NT_MinMax,
			symbols.NT_MinMaxCont,
		},
		MulDivCont1R0,
	},
	MulDivCont1R1: {
		symbols.NT_MulDivCont, 1, 1,
		symbols.Symbols{
			symbols.NT_MulDivCont,
			symbols.NT_MinMax,
			symbols.NT_MinMaxCont,
		},
		MulDivCont1R1,
	},
	MulDivCont1R2: {
		symbols.NT_MulDivCont, 1, 2,
		symbols.Symbols{
			symbols.NT_MulDivCont,
			symbols.NT_MinMax,
			symbols.NT_MinMaxCont,
		},
		MulDivCont1R2,
	},
	MulDivCont1R3: {
		symbols.NT_MulDivCont, 1, 3,
		symbols.Symbols{
			symbols.NT_MulDivCont,
			symbols.NT_MinMax,
			symbols.NT_MinMaxCont,
		},
		MulDivCont1R3,
	},
	Start0R0: {
		symbols.NT_Start, 0, 0,
		symbols.Symbols{
			symbols.NT_AddSubCont,
		},
		Start0R0,
	},
	Start0R1: {
		symbols.NT_Start, 0, 1,
		symbols.Symbols{
			symbols.NT_AddSubCont,
		},
		Start0R1,
	},
	Start1R0: {
		symbols.NT_Start, 1, 0,
		symbols.Symbols{
			symbols.NT_Start,
			symbols.NT_AddSub,
			symbols.NT_AddSubCont,
		},
		Start1R0,
	},
	Start1R1: {
		symbols.NT_Start, 1, 1,
		symbols.Symbols{
			symbols.NT_Start,
			symbols.NT_AddSub,
			symbols.NT_AddSubCont,
		},
		Start1R1,
	},
	Start1R2: {
		symbols.NT_Start, 1, 2,
		symbols.Symbols{
			symbols.NT_Start,
			symbols.NT_AddSub,
			symbols.NT_AddSubCont,
		},
		Start1R2,
	},
	Start1R3: {
		symbols.NT_Start, 1, 3,
		symbols.Symbols{
			symbols.NT_Start,
			symbols.NT_AddSub,
			symbols.NT_AddSubCont,
		},
		Start1R3,
	},
	Unary0R0: {
		symbols.NT_Unary, 0, 0,
		symbols.Symbols{
			symbols.T_13,
			symbols.T_12,
			symbols.T_14,
		},
		Unary0R0,
	},
	Unary0R1: {
		symbols.NT_Unary, 0, 1,
		symbols.Symbols{
			symbols.T_13,
			symbols.T_12,
			symbols.T_14,
		},
		Unary0R1,
	},
	Unary0R2: {
		symbols.NT_Unary, 0, 2,
		symbols.Symbols{
			symbols.T_13,
			symbols.T_12,
			symbols.T_14,
		},
		Unary0R2,
	},
	Unary0R3: {
		symbols.NT_Unary, 0, 3,
		symbols.Symbols{
			symbols.T_13,
			symbols.T_12,
			symbols.T_14,
		},
		Unary0R3,
	},
	Unary1R0: {
		symbols.NT_Unary, 1, 0,
		symbols.Symbols{
			symbols.T_10,
			symbols.T_12,
			symbols.T_11,
		},
		Unary1R0,
	},
	Unary1R1: {
		symbols.NT_Unary, 1, 1,
		symbols.Symbols{
			symbols.T_10,
			symbols.T_12,
			symbols.T_11,
		},
		Unary1R1,
	},
	Unary1R2: {
		symbols.NT_Unary, 1, 2,
		symbols.Symbols{
			symbols.T_10,
			symbols.T_12,
			symbols.T_11,
		},
		Unary1R2,
	},
	Unary1R3: {
		symbols.NT_Unary, 1, 3,
		symbols.Symbols{
			symbols.T_10,
			symbols.T_12,
			symbols.T_11,
		},
		Unary1R3,
	},
	Unary2R0: {
		symbols.NT_Unary, 2, 0,
		symbols.Symbols{
			symbols.T_7,
			symbols.T_12,
			symbols.T_8,
		},
		Unary2R0,
	},
	Unary2R1: {
		symbols.NT_Unary, 2, 1,
		symbols.Symbols{
			symbols.T_7,
			symbols.T_12,
			symbols.T_8,
		},
		Unary2R1,
	},
	Unary2R2: {
		symbols.NT_Unary, 2, 2,
		symbols.Symbols{
			symbols.T_7,
			symbols.T_12,
			symbols.T_8,
		},
		Unary2R2,
	},
	Unary2R3: {
		symbols.NT_Unary, 2, 3,
		symbols.Symbols{
			symbols.T_7,
			symbols.T_12,
			symbols.T_8,
		},
		Unary2R3,
	},
}

var slotIndex = map[Index]Label{
	Index{symbols.NT_AddSub, 0, 0}:     AddSub0R0,
	Index{symbols.NT_AddSub, 0, 1}:     AddSub0R1,
	Index{symbols.NT_AddSub, 1, 0}:     AddSub1R0,
	Index{symbols.NT_AddSub, 1, 1}:     AddSub1R1,
	Index{symbols.NT_AddSubCont, 0, 0}: AddSubCont0R0,
	Index{symbols.NT_AddSubCont, 0, 1}: AddSubCont0R1,
	Index{symbols.NT_AddSubCont, 1, 0}: AddSubCont1R0,
	Index{symbols.NT_AddSubCont, 1, 1}: AddSubCont1R1,
	Index{symbols.NT_AddSubCont, 1, 2}: AddSubCont1R2,
	Index{symbols.NT_AddSubCont, 1, 3}: AddSubCont1R3,
	Index{symbols.NT_MinMax, 0, 0}:     MinMax0R0,
	Index{symbols.NT_MinMax, 0, 1}:     MinMax0R1,
	Index{symbols.NT_MinMax, 1, 0}:     MinMax1R0,
	Index{symbols.NT_MinMax, 1, 1}:     MinMax1R1,
	Index{symbols.NT_MinMaxCont, 0, 0}: MinMaxCont0R0,
	Index{symbols.NT_MinMaxCont, 0, 1}: MinMaxCont0R1,
	Index{symbols.NT_MinMaxCont, 1, 0}: MinMaxCont1R0,
	Index{symbols.NT_MinMaxCont, 1, 1}: MinMaxCont1R1,
	Index{symbols.NT_MinMaxCont, 1, 2}: MinMaxCont1R2,
	Index{symbols.NT_MinMaxCont, 1, 3}: MinMaxCont1R3,
	Index{symbols.NT_MulDiv, 0, 0}:     MulDiv0R0,
	Index{symbols.NT_MulDiv, 0, 1}:     MulDiv0R1,
	Index{symbols.NT_MulDiv, 1, 0}:     MulDiv1R0,
	Index{symbols.NT_MulDiv, 1, 1}:     MulDiv1R1,
	Index{symbols.NT_MulDivCont, 0, 0}: MulDivCont0R0,
	Index{symbols.NT_MulDivCont, 0, 1}: MulDivCont0R1,
	Index{symbols.NT_MulDivCont, 1, 0}: MulDivCont1R0,
	Index{symbols.NT_MulDivCont, 1, 1}: MulDivCont1R1,
	Index{symbols.NT_MulDivCont, 1, 2}: MulDivCont1R2,
	Index{symbols.NT_MulDivCont, 1, 3}: MulDivCont1R3,
	Index{symbols.NT_Start, 0, 0}:      Start0R0,
	Index{symbols.NT_Start, 0, 1}:      Start0R1,
	Index{symbols.NT_Start, 1, 0}:      Start1R0,
	Index{symbols.NT_Start, 1, 1}:      Start1R1,
	Index{symbols.NT_Start, 1, 2}:      Start1R2,
	Index{symbols.NT_Start, 1, 3}:      Start1R3,
	Index{symbols.NT_Unary, 0, 0}:      Unary0R0,
	Index{symbols.NT_Unary, 0, 1}:      Unary0R1,
	Index{symbols.NT_Unary, 0, 2}:      Unary0R2,
	Index{symbols.NT_Unary, 0, 3}:      Unary0R3,
	Index{symbols.NT_Unary, 1, 0}:      Unary1R0,
	Index{symbols.NT_Unary, 1, 1}:      Unary1R1,
	Index{symbols.NT_Unary, 1, 2}:      Unary1R2,
	Index{symbols.NT_Unary, 1, 3}:      Unary1R3,
	Index{symbols.NT_Unary, 2, 0}:      Unary2R0,
	Index{symbols.NT_Unary, 2, 1}:      Unary2R1,
	Index{symbols.NT_Unary, 2, 2}:      Unary2R2,
	Index{symbols.NT_Unary, 2, 3}:      Unary2R3,
}

var alternates = map[symbols.NT][]Label{
	symbols.NT_Start:      []Label{Start0R0, Start1R0},
	symbols.NT_AddSubCont: []Label{AddSubCont0R0, AddSubCont1R0},
	symbols.NT_MulDivCont: []Label{MulDivCont0R0, MulDivCont1R0},
	symbols.NT_MinMaxCont: []Label{MinMaxCont0R0, MinMaxCont1R0},
	symbols.NT_Unary:      []Label{Unary0R0, Unary1R0, Unary2R0},
	symbols.NT_AddSub:     []Label{AddSub0R0, AddSub1R0},
	symbols.NT_MulDiv:     []Label{MulDiv0R0, MulDiv1R0},
	symbols.NT_MinMax:     []Label{MinMax0R0, MinMax1R0},
}
