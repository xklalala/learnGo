package meta

//文件源信息结构体
type FileMeta struct{
	FileSha1 string `json:"fileSha1"`
	FileName string `josn:"fileName"`
	FileSize int64	`json:"fileSize"`
	Location string	`json:"location"`
	UploadAt string	`json:"uploadAt"`//文件上传时间
}

var fileMetas map[string]FileMeta

func init(){
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmeta FileMeta){
	fileMetas[fmeta.FileSha1] = fmeta
}

//通过sha1值获取文件的源信息对象
func GetFileMeta(FileSha1 string) FileMeta{
	return fileMetas[FileSha1]
}