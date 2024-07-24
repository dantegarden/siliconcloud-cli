package lib

type FileReq struct {
	Sign string `protobuf:"bytes,1,opt,name=sign,proto3" json:"sign,omitempty" form:"sign" query:"sign"`
}

type UserInfo struct {
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" query:"name"`
	Image         string `protobuf:"bytes,3,opt,name=image,proto3" json:"image" form:"image" query:"image"`
	Email         string `protobuf:"bytes,4,opt,name=email,proto3" json:"email" form:"email" query:"email"`
	IsAdmin       bool   `protobuf:"varint,5,opt,name=isAdmin,proto3" json:"isAdmin" form:"isAdmin" query:"isAdmin"`
	Balance       string `protobuf:"bytes,6,opt,name=balance,proto3" json:"balance" form:"balance" query:"balance"`
	Status        string `protobuf:"bytes,7,opt,name=status,proto3" json:"status" form:"status" query:"status"`
	Introduction  string `protobuf:"bytes,8,opt,name=introduction,proto3" json:"introduction" form:"introduction" query:"introduction"`
	Role          string `protobuf:"bytes,9,opt,name=role,proto3" json:"role" form:"role" query:"role"`
	ChargeBalance string `protobuf:"bytes,10,opt,name=chargeBalance,proto3" json:"chargeBalance" form:"chargeBalance" query:"chargeBalance"`
	TotalBalance  string `protobuf:"bytes,11,opt,name=totalBalance,proto3" json:"totalBalance" form:"totalBalance" query:"totalBalance"`
}

type FilesResp struct {
	File    *FileInfo    `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty" form:"file" query:"file"`
	Storage *StorageInfo `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty" form:"storage" query:"storage"`
}

type FileInfo struct {
	Sign            string `protobuf:"bytes,1,opt,name=sign,proto3" json:"sign,omitempty" form:"sign" query:"sign"`
	ObjectKey       string `protobuf:"bytes,2,opt,name=objectKey,proto3" json:"object_key,omitempty" form:"object_key" query:"object_key"`
	AccessKeyId     string `protobuf:"bytes,3,opt,name=accessKeyId,proto3" json:"access_key_id,omitempty" form:"access_key_id" query:"access_key_id"`
	AccessKeySecret string `protobuf:"bytes,4,opt,name=accessKeySecret,proto3" json:"access_key_secret,omitempty" form:"access_key_secret" query:"access_key_secret"`
	Expiration      string `protobuf:"bytes,5,opt,name=expiration,proto3" json:"expiration,omitempty" form:"expiration" query:"expiration"`
	SecurityToken   string `protobuf:"bytes,6,opt,name=securityToken,proto3" json:"security_token,omitempty" form:"security_token" query:"security_token"`
	Id              string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
}

type StorageInfo struct {
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty" form:"endpoint" query:"endpoint"`
	Bucket   string `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty" form:"bucket" query:"bucket"`
	Region   string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty" form:"region" query:"region"`
}

type FileCommitReq struct {
	Sign      string `protobuf:"bytes,1,opt,name=sign,proto3" json:"sign,omitempty" form:"sign" query:"sign"`
	ObjectKey string `protobuf:"bytes,2,opt,name=objectKey,proto3" json:"object_key,omitempty" form:"object_key" query:"object_key"`
}

type ModelCommitReq struct {
	Name  string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`
	Type  string       `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" form:"type" query:"type"`
	Files []*ModelFile `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty" form:"files" query:"files"`
}

type ModelFile struct {
	FileId string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"file_id,omitempty" form:"file_id" query:"file_id"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty" form:"path" query:"path"`
}

type ModelCommitResp struct {
}

type ModelInfo struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" form:"type" query:"type"`
	UpdatedAt string `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updated_at,omitempty" form:"updated_at" query:"updated_at"`
}

type ModelListReq struct {
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" form:"type" query:"type"`
}

type ModelListResp struct {
	Models []*ModelInfo `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty" form:"models" query:"models"`
}

type ModelListFilesReq struct {
	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" path:"type"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" path:"name" vd:"(len($) > 0 && len($) < 100)"`
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty" form:"region" query:"region"`
}

type ModelListFilesResp struct {
	Files []*ModelFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" form:"files" query:"files"`
}

type ModelDeleteResp struct {
}
