package model

type UserCanAccess struct {
	CanAccess bool `json:"can_access"`
}

type User struct {
	FeatureName string `json:"featureName" validate:"required"`
	Email       string `json:"email" validate:"required"`
	CanAccess   *bool  `json:"can_access" validate:"required"`
}

type ResponseInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	ResponseInfo ResponseInfo  `json:"response"`
	Data         UserCanAccess `json:"data"`
}
