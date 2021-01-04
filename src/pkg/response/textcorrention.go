package response

type TextCorrentionResponseItem struct {
	VecFragment []struct {
		OriFrag     string `json:"ori_frag"`
		BeginPos    int    `json:"begin_pos"`
		CorrectFrag string `json:"correct_frag"`
		EndPos      int    `json:"end_pos"`
	} `json:"vecFragment"`
	Score     float64 `json:"score"`
	Line      string  `json:"line"`
	Text      string  `json:"text"`
	RequestID string  `json:"requestId"`
	Total     string  `json:"total"`
}
