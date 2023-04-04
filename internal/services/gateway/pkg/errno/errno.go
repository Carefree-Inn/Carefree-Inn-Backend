package errno

type Errno struct {
	message    string `json:"message"`
	statusCode int    `json:"status_code"`
}

func Is(errA error, errB error) bool {
	if _, ok := storage[errA.Error()]; !ok {
		return false
	}
	if _, ok := storage[errB.Error()]; !ok {
		return false
	}
	return storage[errA.Error()] == storage[errB.Error()]
}

func GetCode(message string) int {
	if code, ok := storage[message]; ok {
		return code
	}
	return 20000
}

func (e *Errno) Error() string {
	return e.message
}

var storage = map[string]int{
	JsonDataError.message:       JsonDataError.statusCode,
	LoginServerError.message:    LoginServerError.statusCode,
	LoginWrongInfoError.message: LoginWrongInfoError.statusCode,
	UserNotExistError.message:   UserNotExistError.statusCode,
	TokenGenerateError.message:  TokenGenerateError.statusCode,
	DatabaseError.message:       DatabaseError.statusCode,
}

var (
	// data
	JsonDataError  = &Errno{message: "JSON数据绑定失败", statusCode: 400422}
	PathDataError  = &Errno{message: "path数据绑定失败", statusCode: 400421}
	ParamDataError = &Errno{message: "param数据绑定失败", statusCode: 400422}
	
	// user
	LoginWrongInfoError = &Errno{message: "账号或密码错误", statusCode: 400401}
	LoginServerError    = &Errno{message: "登录出现异常", statusCode: 500500}
	UserNotExistError   = &Errno{message: "用户不存在", statusCode: 200204}
	TokenGenerateError  = &Errno{message: "token生成失败", statusCode: 500501}
	DatabaseError       = &Errno{message: "数据库出现异常", statusCode: 500502}
	UserNotVerifyError  = &Errno{message: "用户身份认证失败", statusCode: 400401}
	TokenNotValidate    = &Errno{message: "无效的用户令牌", statusCode: 400401}
	
	// post
	DeletePostError              = &Errno{message: "删除帖子失败", statusCode: 500501}
	CreatePostError              = &Errno{message: "创建帖子失败", statusCode: 500501}
	GetCategoryInfoError         = &Errno{message: "获取分区失败", statusCode: 500501}
	GetCategoryCategoryPostError = &Errno{message: "获取分区帖子失败", statusCode: 500501}
)
