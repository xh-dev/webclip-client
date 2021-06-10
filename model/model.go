package model

type CreateMsg struct {
	Msg string `json:"msg"`
}
type CreateMsgResponse struct {
	Id string `json:"id"`
}
type RetriveMsg struct {
	Code int `json:"code"`
}
type RetriveMsgResponse struct {
	Msg string `json:"msg"`
}
