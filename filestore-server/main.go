package main
import (
	"filestore-server/handler"
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	err := http.ListenAndServe(":8088", nil)
	if err != nil{
		fmt.Printf("failed to start server, err:%s", err.Error())
	}
}