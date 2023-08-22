package models

type DouyinPublishActionRequest struct {
	Token string `json:"token" binding:"required"`
	Data  []byte `json:"data" binding:"required"`
	Title string `json:"title" binding:"required"`
}

type DouyinPublishActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
