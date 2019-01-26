package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	restful "github.com/emicklei/go-restful"
	do "github.com/ilovelili/digital-ocean-client"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
	"github.com/segmentio/ksuid"
)

// UpdateUserRequest user update request
type UpdateUserRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// UploadAvatar upload user avatar
func UploadAvatar(req *restful.Request, rsp *restful.Response) {
	if err := req.Request.ParseMultipartForm(32 << 20); err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadAvatarFile)
		return
	}

	file, handler, err := req.Request.FormFile("image")
	defer file.Close()
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadAvatarFile)
		return
	}

	filename := handler.Filename
	localfilename := resolveLocalFileName(filename)
	localfile, err := os.Create(localfilename)
	defer os.Remove(localfilename)
	defer localfile.Close()

	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToUploadAvatar)
		return
	}
	io.Copy(localfile, file)

	if !supportedImageMimeType(handler.Header["Content-Type"]) {
		writeError(rsp, errorcode.CoreProxyUnsupportedMimeType)
		return
	}

	spaceservice := do.NewSpaceService(config.Services.APIKey, config.Services.APISecret)
	spaceservice.SetRegion(config.Services.Region)
	spaceservice.SetEndPoint(config.Services.Endpoint)
	spaceservice.SetBucket(config.Services.BucketName)

	uploadopts := &do.UploadOptions{
		Public:   true,
		FileName: localfilename,
	}
	uploadresp := spaceservice.Upload(uploadopts)
	if uploadresp.Error != nil {
		writeError(rsp, errorcode.CoreProxyFailedToUploadAvatar)
		return
	}

	rsp.WriteAsJson(&proto.UploadAvatarResponse{
		Uri: uploadresp.Location,
	})
}

// UpdateUser update user
func UpdateUser(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *UpdateUserRequest
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidUpdateUserRequestBody)
		return
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateUser(ctx(req), &proto.UpdateUserRequest{
		Token:  idtoken,
		Name:   updatereq.Name,
		Avatar: updatereq.Avatar,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// supportedImageMimeType check if uploaded file is image
func supportedImageMimeType(contenttype []string) bool {
	r := regexp.MustCompile("image/(png|jpeg|gif)")
	for _, ct := range contenttype {
		if r.MatchString(ct) {
			return true
		}
	}

	return false
}

// resolveLocalFileName resolve local file name by inject an uuid to avoid duplicates
func resolveLocalFileName(filename string) string {
	dir := filepath.Dir(filename)
	base := fmt.Sprintf("%s_%s", ksuid.New(), filepath.Base(filename))
	return filepath.Join(dir, base)
}
