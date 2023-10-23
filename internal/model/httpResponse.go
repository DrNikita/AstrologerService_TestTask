package model

import "github.com/DrNikita/AstrologerService_TestTask/pkg"

type Response struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

type PageableResponse struct {
	Data pageableData `json:"data"`
	Meta interface{}  `json:"meta"`
}

type pageableData struct {
	Content    interface{} `json:"content"`
	Size       int         `json:"-"`
	Page       int         `json:"-"`
	SortField  string      `json:"-"`
	Direction  string      `json:"-"`
	TotalRows  int64       `json:"totalElements"`
	TotalPages int         `json:"totalPages"`
}

func (pr *PageableResponse) New(d interface{}, p pkg.Pagination, m interface{}) PageableResponse {
	return PageableResponse{
		Data: pageableData{
			Content:    d,
			Size:       p.Size,
			Page:       p.Page,
			SortField:  p.SortField,
			Direction:  p.Direction,
			TotalRows:  p.TotalRows,
			TotalPages: p.TotalPages,
		},
		Meta: m,
	}
}

func (r *Response) New(d interface{}, m interface{}) Response {
	return Response{
		Data: d,
		Meta: m,
	}
}

func (pr *PageableResponse) ErrorResponse(err error) PageableResponse {
	return PageableResponse{
		Meta: err.Error(),
	}
}

func (r *Response) ErrorResponse(err error) Response {
	return Response{
		Meta: err.Error(),
	}
}

func (r *Response) SuccessResponse(message interface{}) Response {
	return Response{
		Meta: message,
	}
}

//func ErrorHandler(c *gin.Context, err error) {
//	goErr := status.Wrap(err, "2")
//	httpResponse := Response{ErrorCode: goErr.Error()}
//	c.AbortWithStatusJSON(500, httpResponse)
//}
