package bind

type QueryMemo struct {
	QueryPage2
}
type BodyMemoCreation struct {
	Content string `json:"content"`
}
type BodyMemoSearch struct {
	QueryPage2
	Content string `json:"content" form:"content"`
}
