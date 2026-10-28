package model

type CreateMsg struct {
	Msg string `json:"msg"`
}
type CreateMsgResponse struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
}
type RetriveMsg struct {
	Code   string `json:"code"`
	Secret string `json:"secret"`
}
type RetriveMsgResponse struct {
	Msg string `json:"msg"`
}
