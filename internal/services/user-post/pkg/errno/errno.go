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

var storage = map[string]int{}

var (
	BadRequestInfo = &Errno{message: "请求信息错误", statusCode: 400400}
)
