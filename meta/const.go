package meta

import (
	"github.com/samber/lo"
	"strings"
)

const (
	CmdLogin   = "login"
	CmdWhoami  = "whoami"
	CmdLogout  = "logout"
	CmdUpload  = "upload"
	CmdModel   = "model"
	CmdLs      = "ls"
	CmdLsFiles = "ls-files"
	CmdRm      = "rm"
	CmdRmFile  = "rm-file"
)

const (
	DefaultDomain = "https://bizyair-api.siliconflow.cn"
)

const (
	LoadError   = 1
	ServerError = 2
	HttpError   = 3
)

type UploadFileType string

const (
	TypeCheckpoint   UploadFileType = "checkpoint"
	TypeVae          UploadFileType = "vae"
	TypeLora         UploadFileType = "lora"
	TypeControlNet   UploadFileType = "controlnet"
	TypeEmbedding    UploadFileType = "embedding"
	TypeHyperNetwork UploadFileType = "hypernetwork"
	TypeClip         UploadFileType = "clip"
	TypeUpscale      UploadFileType = "upscale"
	TypeLLM          UploadFileType = "llm"
)

var ModelTypes = []UploadFileType{
	TypeCheckpoint,
	TypeVae,
	TypeLora,
	TypeControlNet,
	TypeEmbedding,
	TypeHyperNetwork,
	TypeClip,
	TypeUpscale,
	TypeLLM,
}

var ModelTypesStr = func(arr []UploadFileType) string {
	strs := lo.Map[UploadFileType, string](arr, func(v UploadFileType, _ int) string {
		return string(v)
	})
	return "'" + strings.Join(strs, "','") + "'"
}(ModelTypes)

const (
	PercentEncode           = "%2F"
	HTTPGet                 = "GET"
	HTTPPost                = "POST"
	HTTPPut                 = "PUT"
	HTTPDelete              = "DELETE"
	HeaderAuthorization     = "Authorization"
	HeaderContentType       = "Content-Type"
	HeaderSiliconCliVersion = "X-Silicon-CLI-Version"
	JsonContentType         = "application/json"
	APIv1                   = "v1"
	SfFolder                = ".siliconflow"
	SfApiKey                = "apikey"
	OSWindows               = "windows"
	EnvUserProfile          = "USERPROFILE"
	EnvHome                 = "HOME"
	OSSObjectKey            = "https://%s.%s.aliyuncs.com/%s"
	OKCode                  = 20000
)

// IgnoreUploadDirs ignore files when upload
var IgnoreUploadDirs = []string{
	".git",
	".idea",
}
