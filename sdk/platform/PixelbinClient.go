package platform

import (
	"encoding/json"
	"fmt"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/common"
	"os"
)

// PixelbinClient holds all the PixelbinConfig object properties
type PixelbinClient struct {
	Config         *PixelbinConfig
	Assets         *Assets
	Organization   *Organization
	Transformation *Transformation
}

// NewPixelbinClient returns pixelbin service instances
func NewPixelbinClient(config *PixelbinConfig) *PixelbinClient {
	return &PixelbinClient{
		Config:         config,
		Assets:         NewAssets(config),
		Organization:   NewOrganization(config),
		Transformation: NewTransformation(config),
	}
}

// Assets holds Assets object properties
type Assets struct {
	config *PixelbinConfig
}

// NewAssets returns new Assets instance
func NewAssets(config *PixelbinConfig) *Assets {
	return &Assets{config: config}
}

type AddCredentialsXQuery struct {
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	PluginId    string                 `json:"pluginId,omitempty"`
}

/*
summary: Add credentials for a transformation module.

description: Add a transformation modules's credentials for an organization.

params: AddCredentialsXQuery
*/
func (c *Assets) AddCredentials(
	p AddCredentialsXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Credentials map[string]interface{} `json:"credentials,omitempty"`

		PluginId string `json:"pluginId,omitempty"`
	}
	bodydata := &body{

		Credentials: p.Credentials,

		PluginId: p.PluginId,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/credentials",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type UpdateCredentialsXQuery struct {
	PluginId    string                 `json:"pluginId,omitempty"`
	Credentials map[string]interface{} `json:"credentials,omitempty"`
}

/*
summary: Update credentials of a transformation module.

description: Update credentials of a transformation module, for an organization.

params: UpdateCredentialsXQuery
*/
func (c *Assets) UpdateCredentials(
	p UpdateCredentialsXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Credentials map[string]interface{} `json:"credentials,omitempty"`
	}
	bodydata := &body{

		Credentials: p.Credentials,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "patch",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/credentials/%s", p.PluginId),
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type DeleteCredentialsXQuery struct {
	PluginId string `json:"pluginId,omitempty"`
}

/*
summary: Delete credentials of a transformation module.

description: Delete credentials of a transformation module, for an organization.

params: DeleteCredentialsXQuery
*/
func (c *Assets) DeleteCredentials(
	p DeleteCredentialsXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "delete",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/credentials/%s", p.PluginId),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetFileByIdXQuery struct {
	ID string `json:"_id,omitempty"`
}

/*
summary: Get file details with _id

description:

params: GetFileByIdXQuery
*/
func (c *Assets) GetFileById(
	p GetFileByIdXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/files/id/%s", p.ID),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetFileByFileIdXQuery struct {
	FileId string `json:"fileId,omitempty"`
}

/*
summary: Get file details with fileId

description:

params: GetFileByFileIdXQuery
*/
func (c *Assets) GetFileByFileId(
	p GetFileByFileIdXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/files/%s", p.FileId),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type UpdateFileXQuery struct {
	FileId   string                 `json:"fileId,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Path     string                 `json:"path,omitempty"`
	Access   AccessEnum             `json:"access,omitempty"`
	IsActive bool                   `json:"isActive,omitempty"`
	Tags     []string               `json:"tags,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

/*
summary: Update file details

description:

params: UpdateFileXQuery
*/
func (c *Assets) UpdateFile(
	p UpdateFileXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`

		Access AccessEnum `json:"access,omitempty"`

		IsActive bool `json:"isActive,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}
	bodydata := &body{

		Name: p.Name,

		Path: p.Path,

		Access: p.Access,

		IsActive: p.IsActive,

		Tags: p.Tags,

		Metadata: p.Metadata,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "patch",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/files/%s", p.FileId),
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type DeleteFileXQuery struct {
	FileId string `json:"fileId,omitempty"`
}

/*
summary: Delete file

description:

params: DeleteFileXQuery
*/
func (c *Assets) DeleteFile(
	p DeleteFileXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "delete",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/files/%s", p.FileId),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type DeleteFilesXQuery struct {
	Ids []string `json:"ids,omitempty"`
}

/*
summary: Delete multiple files

description:

params: DeleteFilesXQuery
*/
func (c *Assets) DeleteFiles(
	p DeleteFilesXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Ids []string `json:"ids,omitempty"`
	}
	bodydata := &body{

		Ids: p.Ids,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/files/delete",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type CreateFolderXQuery struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

/*
summary: Create folder

description: Create a new folder at the specified path. Also creates the ancestors if they do not exist.

params: CreateFolderXQuery
*/
func (c *Assets) CreateFolder(
	p CreateFolderXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`
	}
	bodydata := &body{

		Name: p.Name,

		Path: p.Path,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/folders",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetFolderDetailsXQuery struct {
	Path string `json:"path,omitempty"`
	Name string `json:"name,omitempty"`
}

/*
summary: Get folder details

description: Get folder details

params: GetFolderDetailsXQuery
*/
func (c *Assets) GetFolderDetails(
	p GetFolderDetailsXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	if p.Path != "" {
		queryParams["path"] = fmt.Sprintf("%v", p.Path)
	}

	if p.Name != "" {
		queryParams["name"] = fmt.Sprintf("%v", p.Name)
	}

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/assets/v1.0/folders",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type UpdateFolderXQuery struct {
	FolderId string `json:"folderId,omitempty"`
	IsActive bool   `json:"isActive,omitempty"`
}

/*
summary: Update folder details

description: Update folder details. Eg: Soft delete it
by making `isActive` as `false`.
We currently do not support updating folder name or path.

params: UpdateFolderXQuery
*/
func (c *Assets) UpdateFolder(
	p UpdateFolderXQuery,
) (map[string]interface{}, error) {

	type body struct {
		IsActive bool `json:"isActive,omitempty"`
	}
	bodydata := &body{

		IsActive: p.IsActive,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "patch",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/folders/%s", p.FolderId),
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type DeleteFolderXQuery struct {
	ID string `json:"_id,omitempty"`
}

/*
summary: Delete folder

description: Delete folder and all its children permanently.

params: DeleteFolderXQuery
*/
func (c *Assets) DeleteFolder(
	p DeleteFolderXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "delete",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/folders/%s", p.ID),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetFolderAncestorsXQuery struct {
	ID string `json:"_id,omitempty"`
}

/*
summary: Get all ancestors of a folder

description: Get all ancestors of a folder, using the folder ID.

params: GetFolderAncestorsXQuery
*/
func (c *Assets) GetFolderAncestors(
	p GetFolderAncestorsXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/folders/%s/ancestors", p.ID),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type ListFilesXQuery struct {
	Name        string        `json:"name,omitempty"`
	Path        string        `json:"path,omitempty"`
	Format      string        `json:"format,omitempty"`
	Tags        []interface{} `json:"tags,omitempty"`
	OnlyFiles   bool          `json:"onlyFiles,omitempty"`
	OnlyFolders bool          `json:"onlyFolders,omitempty"`
	PageNo      float64       `json:"pageNo,omitempty"`
	PageSize    float64       `json:"pageSize,omitempty"`
	Sort        string        `json:"sort,omitempty"`
}

/*
summary: List and search files and folders.

description: List all files and folders in root folder. Search for files if name is provided. If path is provided, search in the specified path.

params: ListFilesXQuery
*/
func (c *Assets) ListFiles(
	p ListFilesXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	if p.Name != "" {
		queryParams["name"] = fmt.Sprintf("%v", p.Name)
	}

	if p.Path != "" {
		queryParams["path"] = fmt.Sprintf("%v", p.Path)
	}

	if p.Format != "" {
		queryParams["format"] = fmt.Sprintf("%v", p.Format)
	}

	if p.Tags != nil {
		queryParams["tags"] = fmt.Sprintf("%v", p.Tags)
	}

	if p.OnlyFiles != false {
		queryParams["onlyFiles"] = fmt.Sprintf("%v", p.OnlyFiles)
	}

	if p.OnlyFolders != false {
		queryParams["onlyFolders"] = fmt.Sprintf("%v", p.OnlyFolders)
	}

	if p.PageNo != 0 {
		queryParams["pageNo"] = fmt.Sprintf("%v", p.PageNo)
	}

	if p.PageSize != 0 {
		queryParams["pageSize"] = fmt.Sprintf("%v", p.PageSize)
	}

	if p.Sort != "" {
		queryParams["sort"] = fmt.Sprintf("%v", p.Sort)
	}

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/assets/v1.0/listFiles",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetDefaultAssetForPlaygroundXQuery struct {
}

/*
summary: Get default asset for playground

description: Get default asset for playground

params: GetDefaultAssetForPlaygroundXQuery
*/
func (c *Assets) GetDefaultAssetForPlayground(
	p GetDefaultAssetForPlaygroundXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/assets/v1.0/playground/default",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetModulesXQuery struct {
}

/*
summary: Get all transformation modules

description: Get all transformation modules.

params: GetModulesXQuery
*/
func (c *Assets) GetModules(
	p GetModulesXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/assets/v1.0/playground/plugins",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetModuleXQuery struct {
	Identifier string `json:"identifier,omitempty"`
}

/*
summary: Get Transformation Module by module identifier

description: Get Transformation Module by module identifier

params: GetModuleXQuery
*/
func (c *Assets) GetModule(
	p GetModuleXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/playground/plugins/%s", p.Identifier),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type AddPresetXQuery struct {
	PresetName     string                 `json:"presetName,omitempty"`
	Transformation string                 `json:"transformation,omitempty"`
	Params         map[string]interface{} `json:"params,omitempty"`
}

/*
summary: Add a preset.

description: Add a preset for an organization.

params: AddPresetXQuery
*/
func (c *Assets) AddPreset(
	p AddPresetXQuery,
) (map[string]interface{}, error) {

	type body struct {
		PresetName string `json:"presetName,omitempty"`

		Transformation string `json:"transformation,omitempty"`

		Params map[string]interface{} `json:"params,omitempty"`
	}
	bodydata := &body{

		PresetName: p.PresetName,

		Transformation: p.Transformation,

		Params: p.Params,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/presets",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetPresetsXQuery struct {
}

/*
summary: Get all presets.

description: Get all presets of an organization.

params: GetPresetsXQuery
*/
func (c *Assets) GetPresets(
	p GetPresetsXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/assets/v1.0/presets",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type UpdatePresetXQuery struct {
	PresetName string `json:"presetName,omitempty"`
	Archived   bool   `json:"archived,omitempty"`
}

/*
summary: Update a preset.

description: Update a preset of an organization.

params: UpdatePresetXQuery
*/
func (c *Assets) UpdatePreset(
	p UpdatePresetXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Archived bool `json:"archived,omitempty"`
	}
	bodydata := &body{

		Archived: p.Archived,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "patch",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/presets/%s", p.PresetName),
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type DeletePresetXQuery struct {
	PresetName string `json:"presetName,omitempty"`
}

/*
summary: Delete a preset.

description: Delete a preset of an organization.

params: DeletePresetXQuery
*/
func (c *Assets) DeletePreset(
	p DeletePresetXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "delete",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/presets/%s", p.PresetName),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetPresetXQuery struct {
	PresetName string `json:"presetName,omitempty"`
}

/*
summary: Get a preset.

description: Get a preset of an organization.

params: GetPresetXQuery
*/
func (c *Assets) GetPreset(
	p GetPresetXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/presets/%s", p.PresetName),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type FileUploadXQuery struct {
	File             *os.File               `json:"file,omitempty"`
	Path             string                 `json:"path,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Access           AccessEnum             `json:"access,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	Overwrite        bool                   `json:"overwrite,omitempty"`
	FilenameOverride bool                   `json:"filenameOverride,omitempty"`
}

/*
summary: Upload File

description: Upload File to Pixelbin

params: FileUploadXQuery
*/
func (c *Assets) FileUpload(
	p FileUploadXQuery,
) (map[string]interface{}, error) {

	type body struct {
		File *os.File `json:"file,omitempty"`

		Path string `json:"path,omitempty"`

		Name string `json:"name,omitempty"`

		Access AccessEnum `json:"access,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`

		Overwrite bool `json:"overwrite,omitempty"`

		FilenameOverride bool `json:"filenameOverride,omitempty"`
	}
	bodydata := &body{

		File: p.File,

		Path: p.Path,

		Name: p.Name,

		Access: p.Access,

		Tags: p.Tags,

		Metadata: p.Metadata,

		Overwrite: p.Overwrite,

		FilenameOverride: p.FilenameOverride,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/upload/direct",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "multipart/form-data",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type UrlUploadXQuery struct {
	URL              string                 `json:"url,omitempty"`
	Path             string                 `json:"path,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Access           AccessEnum             `json:"access,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	Overwrite        bool                   `json:"overwrite,omitempty"`
	FilenameOverride bool                   `json:"filenameOverride,omitempty"`
}

/*
summary: Upload Asset with url

description: Upload Asset with url

params: UrlUploadXQuery
*/
func (c *Assets) UrlUpload(
	p UrlUploadXQuery,
) (map[string]interface{}, error) {

	type body struct {
		URL string `json:"url,omitempty"`

		Path string `json:"path,omitempty"`

		Name string `json:"name,omitempty"`

		Access AccessEnum `json:"access,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`

		Overwrite bool `json:"overwrite,omitempty"`

		FilenameOverride bool `json:"filenameOverride,omitempty"`
	}
	bodydata := &body{

		URL: p.URL,

		Path: p.Path,

		Name: p.Name,

		Access: p.Access,

		Tags: p.Tags,

		Metadata: p.Metadata,

		Overwrite: p.Overwrite,

		FilenameOverride: p.FilenameOverride,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/upload/url",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type CreateSignedUrlXQuery struct {
	Name             string                 `json:"name,omitempty"`
	Path             string                 `json:"path,omitempty"`
	Format           string                 `json:"format,omitempty"`
	Access           AccessEnum             `json:"access,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	Overwrite        bool                   `json:"overwrite,omitempty"`
	FilenameOverride bool                   `json:"filenameOverride,omitempty"`
}

/*
summary: S3 Signed URL upload

description: For the given asset details, a S3 signed URL will be generated,
which can be then used to upload your asset.

params: CreateSignedUrlXQuery
*/
func (c *Assets) CreateSignedUrl(
	p CreateSignedUrlXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`

		Format string `json:"format,omitempty"`

		Access AccessEnum `json:"access,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`

		Overwrite bool `json:"overwrite,omitempty"`

		FilenameOverride bool `json:"filenameOverride,omitempty"`
	}
	bodydata := &body{

		Name: p.Name,

		Path: p.Path,

		Format: p.Format,

		Access: p.Access,

		Tags: p.Tags,

		Metadata: p.Metadata,

		Overwrite: p.Overwrite,

		FilenameOverride: p.FilenameOverride,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/upload/signed-url",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type CreateSignedUrlV2XQuery struct {
	Name             string                 `json:"name,omitempty"`
	Path             string                 `json:"path,omitempty"`
	Format           string                 `json:"format,omitempty"`
	Access           AccessEnum             `json:"access,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	Overwrite        bool                   `json:"overwrite,omitempty"`
	FilenameOverride bool                   `json:"filenameOverride,omitempty"`
	Expiry           float64                `json:"expiry,omitempty"`
}

/*
summary: Signed multipart upload

description: For the given asset details, a presigned URL will be generated, which can be then used to upload your asset in chunks via multipart upload.

params: CreateSignedUrlV2XQuery
*/
func (c *Assets) CreateSignedUrlV2(
	p CreateSignedUrlV2XQuery,
) (map[string]interface{}, error) {

	type body struct {
		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`

		Format string `json:"format,omitempty"`

		Access AccessEnum `json:"access,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`

		Overwrite bool `json:"overwrite,omitempty"`

		FilenameOverride bool `json:"filenameOverride,omitempty"`

		Expiry float64 `json:"expiry,omitempty"`
	}
	bodydata := &body{

		Name: p.Name,

		Path: p.Path,

		Format: p.Format,

		Access: p.Access,

		Tags: p.Tags,

		Metadata: p.Metadata,

		Overwrite: p.Overwrite,

		FilenameOverride: p.FilenameOverride,

		Expiry: p.Expiry,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v2.0/upload/signed-url",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

// Organization holds Organization object properties
type Organization struct {
	config *PixelbinConfig
}

// NewOrganization returns new Organization instance
func NewOrganization(config *PixelbinConfig) *Organization {
	return &Organization{config: config}
}

type GetAppOrgDetailsXQuery struct {
}

/*
summary: Get App Details

description: Get App and org details

params: GetAppOrgDetailsXQuery
*/
func (c *Organization) GetAppOrgDetails(
	p GetAppOrgDetailsXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/organization/v1.0/apps/info",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

// Transformation holds Transformation object properties
type Transformation struct {
	config *PixelbinConfig
}

// NewTransformation returns new Transformation instance
func NewTransformation(config *PixelbinConfig) *Transformation {
	return &Transformation{config: config}
}

type GetTransformationContextXQuery struct {
	URL string `json:"url,omitempty"`
}

/*
summary: Get transformation context

description: Get transformation context

params: GetTransformationContextXQuery
*/
func (c *Transformation) GetTransformationContext(
	p GetTransformationContextXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	if p.URL != "" {
		queryParams["url"] = fmt.Sprintf("%v", p.URL)
	}

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/transformation/context",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}
