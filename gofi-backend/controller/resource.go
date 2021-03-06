package controller

type Resource struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Code    int         `json:"code"` // 200 means success, other means fail
	Data    interface{} `json:"data"`
}

const StatusOk int = 200       // success
const StatusUnknown int = -1   // fail,reason is unknown
const StatusNotFound int = 404 // fail,reason is not found

// auth
const StatusTokenMiss int = 10000                 // token miss
const StatusTokenInvalid int = 10001              // token is invalid
const StatusTokenExpired int = 10002              // token is expired
const StatusUsernameOrPasswordIsWrong int = 10003 // username or password is wrong

// config
const StatusAlreadyInitialized = 20000   // fail,reason is already initialized
const StatusNotInitialized = 20001       // fail,reason is not initialized
const StatusCurrentIsPreviewMode = 20002 // fail,reason is current is preview mode

// file
const StatusDirNotExist = 30000       // fail,reason is dir not exist
const StatusFileNotExist = 30001      // fail,reason is file not exist
const StatusUploadFileFailed = 30002  // fail,reason is upload file failed
const StatusCantOverwriteFile = 30003 // fail,reason is can't overwrite file

type ResourceBuilder struct {
	resource *Resource
}

// NewResource success default is true, code default is 200
func NewResource() ResourceBuilder {
	return ResourceBuilder{resource: &Resource{Success: true, Code: StatusOk, Message: nil, Data: nil}}
}

func (builder ResourceBuilder) Code(code int) ResourceBuilder {
	builder.resource.Code = code
	builder.resource.Success = code == StatusOk
	return builder
}

func (builder ResourceBuilder) Message(message interface{}) ResourceBuilder {
	builder.resource.Message = message
	return builder
}

func (builder ResourceBuilder) Success() ResourceBuilder {
	builder.resource.Success = true
	builder.resource.Code = StatusOk
	return builder
}

func (builder ResourceBuilder) Fail() ResourceBuilder {
	builder.resource.Success = false
	builder.resource.Code = StatusUnknown
	return builder
}

func (builder ResourceBuilder) Payload(data interface{}) ResourceBuilder {
	builder.resource.Data = data
	return builder
}

func (builder ResourceBuilder) Build() *Resource {
	return builder.resource
}
