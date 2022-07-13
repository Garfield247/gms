package resp

import "encoding/json"

type ResponseInterface interface {
}

// 响应请求的公共模型
type ResponseModel struct {
	Code    int    `json:"code"`    // 响应码
	Msg     string `json:"msg"`     // 响应信息
	Results string `json:"results"` // 数据
}

// 上传响应数据
type UpdateDate struct {
	Size  int64  `json:"size"`  // 大小
	Mime  string `json:"mime"`  // 图片类型
	Imgid string `json:"imgid"` // 图片id
}

// 上传响应模型
type UpdateResponse struct {
	ResponseModel
	Results UpdateDate `json:"results"`
}

// 响应 json 打包
func ResponseJson(res ResponseInterface) []byte {

	data, err := json.Marshal(res)
	if err != err {

		// 打包失败
		data, _ = json.Marshal(ResponseModel{
			StatusJson,
			StatusText(StatusJson),
			""})
	}

	return data
}
