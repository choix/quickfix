package nstdptys2subgrp

func New() *NstdPtys2SubGrp {
	var m NstdPtys2SubGrp
	return &m
}

//NoNested2PartySubIDs is a repeating group in NstdPtys2SubGrp
type NoNested2PartySubIDs struct {
	//Nested2PartySubID is a non-required field for NoNested2PartySubIDs.
	Nested2PartySubID *string `fix:"760"`
	//Nested2PartySubIDType is a non-required field for NoNested2PartySubIDs.
	Nested2PartySubIDType *int `fix:"807"`
}

//NstdPtys2SubGrp is a fix50 Component
type NstdPtys2SubGrp struct {
	//NoNested2PartySubIDs is a non-required field for NstdPtys2SubGrp.
	NoNested2PartySubIDs []NoNested2PartySubIDs `fix:"806,omitempty"`
}

func (m *NstdPtys2SubGrp) SetNoNested2PartySubIDs(v []NoNested2PartySubIDs) {
	m.NoNested2PartySubIDs = v
}
