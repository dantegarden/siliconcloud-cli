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
	TypeCheckpoint   UploadFileType = "bizyair/checkpoint"
	TypeVae          UploadFileType = "bizyair/vae"
	TypeLora         UploadFileType = "bizyair/lora"
	TypeControlNet   UploadFileType = "bizyair/controlnet"
	TypeEmbedding    UploadFileType = "bizyair/embedding"
	TypeHyperNetwork UploadFileType = "bizyair/hypernetwork"
	TypeClip         UploadFileType = "bizyair/clip"
	TypeClipVision   UploadFileType = "bizyair/clip_vision"
	TypeUpscale      UploadFileType = "bizyair/upscale"
	TypeOther        UploadFileType = "other"
)

var ModelTypes = []UploadFileType{
	TypeCheckpoint,
	TypeVae,
	TypeLora,
	TypeControlNet,
	TypeEmbedding,
	TypeHyperNetwork,
	TypeClip,
	TypeClipVision,
	TypeUpscale,
	TypeOther,
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
	EnvAPIKey               = "SF_API_KEY"
	OSSObjectKey            = "https://%s.%s.aliyuncs.com/%s"
	OKCode                  = 20000
)

// IgnoreUploadDirs ignore files when upload
var IgnoreUploadDirs = []string{
	".git",
	".idea",
}
