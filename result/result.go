package result

type R struct {
}

func (R) Success(data interface{}) map[string]interface{} {
	if data != nil {
		return map[string]interface{}{
			"code": 200,
			"msg":  "success",
			"data": data,
		}
	} else {
		return map[string]interface{}{
			"code": 200,
			"msg":  "success",
		}
	}
}

func (R) Error(code int, msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": code, "msg": msg,
	}
}
