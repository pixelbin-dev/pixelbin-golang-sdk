package tests

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

var settings = map[string]string{"domain": "https://api.testdomain.com", "apiSecret": "test-api-secret"}
var config *platform.PixelbinConfig
var pixelbin *platform.PixelbinClient
var domain string
var apiSecret string

func Init() {
	settings["apiSecret"] = apiSecret
	settings["domain"] = domain
	settings["folderName"] = "testdir"
	settings["folderPath"] = "/"
	config = platform.NewPixelbinConfig(settings["apiSecret"], settings["domain"])
	config.SetOAuthClient()
	pixelbin = platform.NewPixelbinClient(config)
}

func TestCreateFolder(t *testing.T) {
	params := platform.CreateFolderXQuery{Name: settings["folderName"], Path: settings["folderPath"]}
	_, err := pixelbin.Assets.CreateFolder(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetFolderDetails(t *testing.T) {
	params := platform.GetFolderDetailsXQuery{
		Path: "",
		Name: settings["folderName"],
	}

	_, err := pixelbin.Assets.GetFolderDetails(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetFolderAncestors(t *testing.T) {
	params := platform.CreateFolderXQuery{Name: "folder", Path: "nested"}
	resp, err := pixelbin.Assets.CreateFolder(params)
	jsonBody, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("Failed ! error in marshalling %v", err)
	}
	var x platform.FoldersResponse
	err = json.Unmarshal(jsonBody, &x)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	}
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		reqParams := platform.GetFolderAncestorsXQuery{ID: x.ID}
		_, err := pixelbin.Assets.GetFolderAncestors(reqParams)
		if err != nil {
			t.Errorf("Failed ! got err %v", err)
		} else {
			t.Logf("Success")
		}
	}
}

func TestFileUploadCase1(t *testing.T) {
	file, _ := os.Open("1.jpeg")
	params := platform.FileUploadXQuery{File: file}
	_, err := pixelbin.Assets.FileUpload(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestFileUploadCase2(t *testing.T) {
	tags := []string{"tag1", "tag2"}
	file, _ := os.Open("1.jpeg")
	params := platform.FileUploadXQuery{
		File:             file,
		Path:             settings["folderName"],
		Name:             "1",
		Access:           "public-read",
		Tags:             tags,
		Metadata:         map[string]interface{}{},
		Overwrite:        false,
		FilenameOverride: true,
	}
	res, err := pixelbin.Assets.FileUpload(params)

	jsonbody, err := json.Marshal(res)
	if err != nil {
		t.Errorf("Failed ! error in marshalling %v", err)
	}
	var x platform.UploadResponse
	err = json.Unmarshal(jsonbody, &x)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	}

	for i, tag := range tags {
		// println("tag  ", tag, " exp ", tags[i])
		if tag != x.Tags[i] {
			t.Errorf("Failed ! expected %v got err %v", tag, x.Tags[i])
		}
	}
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestFileUploadCaseOverwrite(t *testing.T) {
	file, _ := os.Open("1.jpeg")
	params := platform.FileUploadXQuery{
		File:             file,
		Name:             "test",
		Access:           "public-read",
		Tags:             []string{"tag1", "tag2"},
		Metadata:         map[string]interface{}{},
		Overwrite:        true,
		FilenameOverride: false,
	}
	_, err := pixelbin.Assets.FileUpload(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}

	file2, _ := os.Open("1.jpeg")
	params2 := platform.FileUploadXQuery{
		File:             file2,
		Name:             "test",
		Access:           "public-read",
		Tags:             []string{"tag1", "tag2"},
		Metadata:         map[string]interface{}{},
		Overwrite:        true,
		FilenameOverride: false,
	}
	_, err2 := pixelbin.Assets.FileUpload(params2)
	if err2 != nil {
		t.Errorf("Failed ! got err2 %v", err2)
	} else {
		t.Logf("Success")
	}
}

func TestListFilesCase1(t *testing.T) {
	params := platform.ListFilesXQuery{}
	_, err := pixelbin.Assets.ListFiles(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestListFilesCase2(t *testing.T) {
	params := platform.ListFilesXQuery{}
	_, err := pixelbin.Assets.ListFiles(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUrlUpload(t *testing.T) {
	tags := []string{"cat", "animal"}

	params := platform.UrlUploadXQuery{
		URL:              "https://www.fetchfind.com/blog/wp-content/uploads/2017/08/cat-2734999_1920-5-common-cat-sounds.jpg",
		Path:             settings["folderName"],
		Name:             "2",
		Access:           "public-read",
		Tags:             tags,
		Metadata:         map[string]interface{}{},
		Overwrite:        false,
		FilenameOverride: true,
	}
	res, err := pixelbin.Assets.UrlUpload(params)

	jsonbody, err := json.Marshal(res)
	if err != nil {
		t.Errorf("Failed ! error in marshalling %v", err)
	}
	var x platform.UploadResponse
	err = json.Unmarshal(jsonbody, &x)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	}

	for i, tag := range tags {
		// println("tag  ", tag, " exp ", tags[i])
		if tag != x.Tags[i] {
			t.Errorf("Failed ! expected %v got err %v", tag, x.Tags[i])
		}
	}
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestCreateSignedUrlCase1(t *testing.T) {
	params := platform.CreateSignedUrlXQuery{}
	_, err := pixelbin.Assets.CreateSignedUrl(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestCreateSignedUrlCase2(t *testing.T) {
	params := platform.CreateSignedUrlXQuery{
		Name:             "1",
		Path:             settings["folderName"],
		Format:           "jpeg",
		Access:           "public-read",
		Tags:             []string{"tag1", "tag2"},
		Metadata:         map[string]interface{}{},
		Overwrite:        false,
		FilenameOverride: true,
	}
	_, err := pixelbin.Assets.CreateSignedUrl(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestCreateSignedUrlV2Case1(t *testing.T) {
	params := platform.CreateSignedUrlV2XQuery{}
	_, err := pixelbin.Assets.CreateSignedUrlV2(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestCreateSignedUrlV2Case2(t *testing.T) {
	params := platform.CreateSignedUrlV2XQuery{
		Name:             "1",
		Path:             settings["folderName"],
		Format:           "jpeg",
		Access:           "public-read",
		Tags:             []string{"tag1", "tag2"},
		Metadata:         map[string]interface{}{},
		Overwrite:        false,
		FilenameOverride: true,
	}
	_, err := pixelbin.Assets.CreateSignedUrlV2(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUpdateFileCase1(t *testing.T) {
	params := platform.UpdateFileXQuery{
		FileId: "1",
		Name:   "1_",
	}
	_, err := pixelbin.Assets.UpdateFile(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUpdateFileCase2(t *testing.T) {
	params := platform.UpdateFileXQuery{
		FileId:   fmt.Sprintf("%s/1", settings["folderName"]),
		Name:     fmt.Sprintf("%s_", settings["folderName"]),
		Path:     settings["folderName"],
		Access:   "private",
		IsActive: false,
		Tags:     []string{"tag1", "tag2"},
		Metadata: map[string]interface{}{"key": "value"},
	}
	_, err := pixelbin.Assets.UpdateFile(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetFileByFileId(t *testing.T) {
	params := platform.GetFileByFileIdXQuery{
		FileId: fmt.Sprintf("%s/2", settings["folderName"]),
	}
	_, err := pixelbin.Assets.GetFileByFileId(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestDeleteFile(t *testing.T) {
	params := platform.DeleteFileXQuery{
		FileId: "1_",
	}
	_, err := pixelbin.Assets.DeleteFile(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUpdateFolder(t *testing.T) {
	params := platform.UpdateFolderXQuery{
		FolderId: settings["folderName"],
		IsActive: false,
	}
	_, err := pixelbin.Assets.UpdateFolder(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestDeleteFolder(t *testing.T) {
	params := platform.ListFilesXQuery{}
	res, err := pixelbin.Assets.ListFiles(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	}
	jsonbody, err := json.Marshal(res)
	if err != nil {
		t.Errorf("Failed ! error in marshalling %v", err)
	}

	var id string = ""
	var x platform.ListFilesResponse
	err = json.Unmarshal(jsonbody, &x)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	}
	for i := range x.Items {
		item := x.Items[i]
		if item.Type == "folder" && item.Name == settings["folderName"] {
			id = item.ID
			break
		}
	}
	if id != "" {
		param := platform.DeleteFolderXQuery{
			ID: id,
		}
		_, err = pixelbin.Assets.DeleteFolder(param)
		if err != nil {
			t.Errorf("Failed ! got err %v", err)
		} else {
			t.Logf("Success")
		}
	}
}

func TestGetModules(t *testing.T) {
	params := platform.GetModulesXQuery{}
	_, err := pixelbin.Assets.GetModules(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetModule(t *testing.T) {
	params := platform.GetModuleXQuery{
		Identifier: "t",
	}
	_, err := pixelbin.Assets.GetModule(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestAddCredentials(t *testing.T) {
	params := platform.AddCredentialsXQuery{
		Credentials: map[string]interface{}{"apiKey": "dummy_key_replace_with_real"},
		PluginId:    "remove",
	}
	_, err := pixelbin.Assets.AddCredentials(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUpdateCredentials(t *testing.T) {
	params := platform.UpdateCredentialsXQuery{
		PluginId:    "remove",
		Credentials: map[string]interface{}{"apiKey": "dummy_key_replace_with_real"},
	}
	_, err := pixelbin.Assets.UpdateCredentials(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestDeleteCredentials(t *testing.T) {
	params := platform.DeleteCredentialsXQuery{
		PluginId: "remove",
	}
	_, err := pixelbin.Assets.DeleteCredentials(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestAddPreset(t *testing.T) {
	params := platform.AddPresetXQuery{
		PresetName:     "p1",
		Transformation: "t.flip()~t.flop()",
		Params: map[string]interface{}{
			"w": map[string]interface{}{"type": "integer", "default": 200},
			"h": map[string]interface{}{"type": "integer", "default": 400},
		},
	}
	_, err := pixelbin.Assets.AddPreset(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetPresets(t *testing.T) {
	params := platform.GetPresetsXQuery{}
	_, err := pixelbin.Assets.GetPresets(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUpdatePreset(t *testing.T) {
	params := platform.UpdatePresetXQuery{
		PresetName: "p1",
		Archived:   true,
	}
	_, err := pixelbin.Assets.UpdatePreset(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetPreset(t *testing.T) {
	params := platform.GetPresetXQuery{
		PresetName: "p1",
	}
	_, err := pixelbin.Assets.GetPreset(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}

}

func TestDeletePreset(t *testing.T) {
	params := platform.DeletePresetXQuery{
		PresetName: "p1",
	}
	_, err := pixelbin.Assets.DeletePreset(params)

	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetDefaultAssetForPlayground(t *testing.T) {
	params := platform.GetDefaultAssetForPlaygroundXQuery{}
	_, err := pixelbin.Assets.GetDefaultAssetForPlayground(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetTransformationContext(t *testing.T) {
	params := platform.GetTransformationContextXQuery{
		// provide link of your image
		URL: "/v2/restless-cloud-a0494f/erase.bg()/sample1.webp",
	}
	resp, err := pixelbin.Transformation.GetTransformationContext(params)

	if err != nil {
		t.Errorf("Failed! Got error: %v", err)
	} else {
		t.Logf("Success!")
	}

	if _, ok := resp["context"]; !ok {
		t.Errorf("Failed! Expected 'context' key in response, but it was not found.")
	}
}

func TestMain(m *testing.M) {

	flag.StringVar(&domain, "domain", settings["domain"], "Domain Address")
	flag.StringVar(&apiSecret, "apiSecret", settings["apiSecret"], "Api secret token")

	flag.Parse()
	Init()
	code := m.Run()
	os.Exit(code)
}
