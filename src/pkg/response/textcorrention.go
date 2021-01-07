package response

type TextCorrentionResponseItem struct {
	VecFragment []VecFragmentJSON `json:"vecFragment"`
	Score       float64           `json:"score"`
	Line        string            `json:"line"`
	Text        string            `json:"text"`
	RequestID   string            `json:"requestId"`
	Total       string            `json:"total"`
}

type VecFragment struct {
	OriFrag     string `json:"ori_frag,omitempty"`
	BeginPos    string `json:"begin_pos,omitempty"`
	CorrectFrag string `json:"correct_frag,omitempty"`
	EndPos      string `json:"end_pos,omitempty"`
}

type VecFragmentJSON struct {
	*VecFragment
}
