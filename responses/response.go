package responses

type BaseResponse struct {
	Success bool
	Message string
}

type BaseResponseList struct {
	Response
	RecordsTotal    int
	RecordsFiltered int
}

type Response struct {
	BaseResponse
	Data interface{}
}

type ResponseList struct {
	BaseResponseList
	Rows interface{}
}
