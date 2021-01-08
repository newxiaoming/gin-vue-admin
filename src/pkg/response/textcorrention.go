package response

type TextCorrentionResponseItem struct {
	VecFragment []map[string]interface{} `json:"vecFragment"`
	Score       float64                  `json:"score"`
	Line        string                   `json:"line"`
	Text        string                   `json:"text"`
	RequestID   string                   `json:"requestId"`
	Total       string                   `json:"total"`
}

type VecFragment struct {
	OriFrag     string `json:"ori_frag"`
	BeginPos    string `json:"begin_pos"`
	CorrectFrag string `json:"correct_frag"`
	EndPos      string `json:"end_pos"`
}

type VecFragmentJSON struct {
	*VecFragment
}
