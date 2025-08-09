package v1

type PolishContentRequest struct {
	Prompt string `json:"prompt" valid:"stringlength(0|16384)"`
	Content string `json:"content" valid:"stringlength(0|16384)"`
}

type PolishContentResponse struct {
	Content string `json:"content" valid:"stringlength(0|16384)"`
}

type GenerateTitleRequest struct {
	Prompt string `json:"prompt" valid:"stringlength(0|16384)"`
	Content string `json:"content" valid:"stringlength(0|16384)"`
}

type GenerateTitleResponse struct {
	Title string `json:"title" valid:"stringlength(0|16384)"`
}
