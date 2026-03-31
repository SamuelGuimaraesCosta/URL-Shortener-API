package dto

type CreateURLRequest struct {
	URL string `json:"url"`
}

type CreateURLResponse struct {
	Code  string `json:"code"`
	Short string `json:"short"`
}
