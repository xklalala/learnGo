package handler
import (
	"net/http"
	"io/ioutil"
	"io"
	"fmt"
	"os"
	"filestore-server/meta"
	"filestore-server/util"
	"time"
	"encoding/json"
)

//处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		//返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil{
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	}else if r.Method == "POST"{
		//接收文件流及存储到本地
		file, head, err := r.FormFile("file")
		if err != nil{
			fmt.Printf("failed to get data, err: %s\n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			// FileSha1 : 
			FileName : head.Filename,
			// FileSize :
			Location : "./temp/" + head.Filename,
			UploadAt : time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil{
			fmt.Printf("failed to create file, err: %s\n", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil{
			fmt.Printf("failed to create file, err: %s\n", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		fmt.Println(fileMeta.FileSha1)
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

//上传已完成
func UploadSucHandler (w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "upload finished!")
}


//获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

//文件下载
func DownloadHandler(w http.ResponseWriter, r * http.Request){
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)

	f, err := os.Open(fm.Location)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(fm.FileName)
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-disposition", fmt.Sprintf("attachment; filename=%s", fm.FileName))
	w.Write(data)
}