##### [Back to Pixelbin API docs](./README.md)

## Assets Methods

Asset Uploader Service

-   [AddCredentials](#addcredentials)
-   [UpdateCredentials](#updatecredentials)
-   [DeleteCredentials](#deletecredentials)
-   [GetFileById](#getfilebyid)
-   [GetFileByFileId](#getfilebyfileid)
-   [UpdateFile](#updatefile)
-   [DeleteFile](#deletefile)
-   [DeleteFiles](#deletefiles)
-   [CreateFolder](#createfolder)
-   [GetFolderDetails](#getfolderdetails)
-   [UpdateFolder](#updatefolder)
-   [DeleteFolder](#deletefolder)
-   [GetFolderAncestors](#getfolderancestors)
-   [ListFiles](#listfiles)
-   [ListFilesPaginator](#listfilespaginator)
-   [GetDefaultAssetForPlayground](#getdefaultassetforplayground)
-   [GetModules](#getmodules)
-   [GetModule](#getmodule)
-   [AddPreset](#addpreset)
-   [GetPresets](#getpresets)
-   [UpdatePreset](#updatepreset)
-   [DeletePreset](#deletepreset)
-   [GetPreset](#getpreset)
-   [FileUpload](#fileupload)
-   [UrlUpload](#urlupload)
-   [CreateSignedUrl](#createsignedurl)
-   [CreateSignedUrlV2](#createsignedurlv2)

## Methods with example and description

### AddCredentials

**Summary**: Add credentials for a transformation module.

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.AddCredentialsXQuery{
        Credentials: map[string]interface{}{"region":"ap-south-1","accessKeyId":"123456789ABC","secretAccessKey":"DUMMY1234567890"},
        PluginId: "awsRek",
    }
    result, err := pixelbin.Assets.AddCredentials(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument    | Type                   | Required | Description                                                 |
| ----------- | ---------------------- | -------- | ----------------------------------------------------------- |
| Credentials | map[string]interface{} | yes      | Credentials of the plugin                                   |
| PluginId    | string                 | yes      | Unique identifier for the plugin this credential belongs to |

Add a transformation modules's credentials for an organization.

_Returned Response:_

[AddCredentialsResponse](#addcredentialsresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "123ee789-7ae8-4336-b9bd-e4f33c049002",
    "createdAt": "2022-10-04T09:52:09.545Z",
    "updatedAt": "2022-10-04T09:52:09.545Z",
    "orgId": 23,
    "pluginId": "awsRek"
}
```

</details>

### UpdateCredentials

**Summary**: Update credentials of a transformation module.

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.UpdateCredentialsXQuery{
        PluginId: "awsRek",,
        Credentials: map[string]interface{}{"region":"ap-south-1","accessKeyId":"123456789ABC","secretAccessKey":"DUMMY1234567890"},
    }
    result, err := pixelbin.Assets.UpdateCredentials(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument    | Type                   | Required | Description                                          |
| ----------- | ---------------------- | -------- | ---------------------------------------------------- |
| PluginId    | string                 | yes      | ID of the plugin whose credentials are being updated |
| Credentials | map[string]interface{} | yes      | Credentials of the plugin                            |

Update credentials of a transformation module, for an organization.

_Returned Response:_

[AddCredentialsResponse](#addcredentialsresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "123ee789-7ae8-4336-b9bd-e4f33c049002",
    "createdAt": "2022-10-04T09:52:09.545Z",
    "updatedAt": "2022-10-04T09:52:09.545Z",
    "orgId": 23,
    "pluginId": "awsRek"
}
```

</details>

### DeleteCredentials

**Summary**: Delete credentials of a transformation module.

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.DeleteCredentialsXQuery{
        PluginId: "awsRek",
    }
    result, err := pixelbin.Assets.DeleteCredentials(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                                          |
| -------- | ------ | -------- | ---------------------------------------------------- |
| PluginId | string | yes      | ID of the plugin whose credentials are being deleted |

Delete credentials of a transformation module, for an organization.

_Returned Response:_

[AddCredentialsResponse](#addcredentialsresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "123ee789-7ae8-4336-b9bd-e4f33c049002",
    "createdAt": "2022-10-04T09:52:09.545Z",
    "updatedAt": "2022-10-04T09:52:09.545Z",
    "orgId": 23,
    "pluginId": "awsRek"
}
```

</details>

### GetFileById

**Summary**: Get file details with \_id

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetFileByIdXQuery{
        ID: "c9138153-94ea-4dbe-bea9-65d43dba85ae",
    }
    result, err := pixelbin.Assets.GetFileById(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description  |
| -------- | ------ | -------- | ------------ |
| ID       | string | yes      | \_id of File |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### GetFileByFileId

**Summary**: Get file details with fileId

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetFileByFileIdXQuery{
        FileId: "path/to/file/name",
    }
    result, err := pixelbin.Assets.GetFileByFileId(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                              |
| -------- | ------ | -------- | ---------------------------------------- |
| FileId   | string | yes      | Combination of `path` and `name` of file |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### UpdateFile

**Summary**: Update file details

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.UpdateFileXQuery{
        FileId: "path/to/file/name",,
        Name: "asset",
        Path: "dir",
        Access: "private",
        IsActive: false,
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{"key":"value"},
    }
    result, err := pixelbin.Assets.UpdateFile(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type                   | Required | Description                                                     |
| -------- | ---------------------- | -------- | --------------------------------------------------------------- |
| FileId   | string                 | yes      | Combination of `path` and `name`                                |
| Name     | string                 | no       | Name of the file                                                |
| Path     | string                 | no       | Path of the file                                                |
| Access   | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private` |
| IsActive | bool                   | no       | Whether the file is active                                      |
| Tags     | []string               | no       | Tags associated with the file                                   |
| Metadata | map[string]interface{} | no       | Metadata associated with the file                               |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### DeleteFile

**Summary**: Delete file

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.DeleteFileXQuery{
        FileId: "path/to/file/name",
    }
    result, err := pixelbin.Assets.DeleteFile(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                      |
| -------- | ------ | -------- | -------------------------------- |
| FileId   | string | yes      | Combination of `path` and `name` |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### DeleteFiles

**Summary**: Delete multiple files

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.DeleteFilesXQuery{
        Ids: []string{"_id_1","_id_2","_id_3"},
    }
    result, err := pixelbin.Assets.DeleteFiles(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type     | Required | Description                   |
| -------- | -------- | -------- | ----------------------------- |
| Ids      | []string | yes      | Array of file \_ids to delete |

_Returned Response:_

[[]FilesResponse](#[filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
[
    {
        "_id": "dummy-uuid",
        "name": "asset",
        "path": "dir",
        "fileId": "dir/asset",
        "format": "jpeg",
        "size": 1000,
        "access": "private",
        "isActive": true,
        "tags": ["tag1", "tag2"],
        "metadata": {
            "key": "value"
        },
        "url": "https://domain.com/filename.jpeg"
    }
]
```

</details>

### CreateFolder

**Summary**: Create folder

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.CreateFolderXQuery{
        Name: "subDir",
        Path: "dir",
    }
    result, err := pixelbin.Assets.CreateFolder(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description        |
| -------- | ------ | -------- | ------------------ |
| Name     | string | yes      | Name of the folder |
| Path     | string | no       | Path of the folder |

Create a new folder at the specified path. Also creates the ancestors if they do not exist.

_Returned Response:_

[FoldersResponse](#foldersresponse)

Success - List of all created folders

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "subDir",
    "path": "dir",
    "isActive": true
}
```

</details>

### GetFolderDetails

**Summary**: Get folder details

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetFolderDetailsXQuery{
        Path: "dir1/dir2",
        Name: "dir",
    }
    result, err := pixelbin.Assets.GetFolderDetails(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description |
| -------- | ------ | -------- | ----------- |
| Path     | string | no       | Folder path |
| Name     | string | no       | Folder name |

Get folder details

_Returned Response:_

[exploreItem](#exploreitem)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
[
    {
        "_id": "dummy-uuid",
        "createdAt": "2022-10-05T10:43:04.117Z",
        "updatedAt": "2022-10-05T10:43:04.117Z",
        "name": "asset2",
        "type": "file",
        "path": "dir",
        "fileId": "dir/asset2",
        "format": "jpeg",
        "size": 1000,
        "access": "private",
        "metadata": {},
        "height": 100,
        "width": 100
    }
]
```

</details>

### UpdateFolder

**Summary**: Update folder details

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.UpdateFolderXQuery{
        FolderId: "path/to/folder/name",,
        IsActive: false,
    }
    result, err := pixelbin.Assets.UpdateFolder(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                      |
| -------- | ------ | -------- | -------------------------------- |
| FolderId | string | yes      | combination of `path` and `name` |
| IsActive | bool   | no       | whether the folder is active     |

Update folder details. Eg: Soft delete it
by making `isActive` as `false`.
We currently do not support updating folder name or path.

_Returned Response:_

[FoldersResponse](#foldersresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "subDir",
    "path": "dir",
    "isActive": true
}
```

</details>

### DeleteFolder

**Summary**: Delete folder

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.DeleteFolderXQuery{
        ID: "c9138153-94ea-4dbe-bea9-65d43dba85ae",
    }
    result, err := pixelbin.Assets.DeleteFolder(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                  |
| -------- | ------ | -------- | ---------------------------- |
| ID       | string | yes      | \_id of folder to be deleted |

Delete folder and all its children permanently.

_Returned Response:_

[FoldersResponse](#foldersresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "subDir",
    "path": "dir",
    "isActive": true
}
```

</details>

### GetFolderAncestors

**Summary**: Get all ancestors of a folder

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetFolderAncestorsXQuery{
        ID: "c9138153-94ea-4dbe-bea9-65d43dba85ae",
    }
    result, err := pixelbin.Assets.GetFolderAncestors(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description        |
| -------- | ------ | -------- | ------------------ |
| ID       | string | yes      | \_id of the folder |

Get all ancestors of a folder, using the folder ID.

_Returned Response:_

[GetAncestorsResponse](#getancestorsresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "folder": {
        "_id": "dummy-uuid",
        "name": "subDir",
        "path": "dir1/dir2",
        "isActive": true
    },
    "ancestors": [
        {
            "_id": "dummy-uuid-2",
            "name": "dir1",
            "path": "",
            "isActive": true
        },
        {
            "_id": "dummy-uuid-2",
            "name": "dir2",
            "path": "dir1",
            "isActive": true
        }
    ]
}
```

</details>

### ListFiles

**Summary**: List and search files and folders.

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.ListFilesXQuery{
        Name: "cat",
        Path: "cat-photos",
        Format: "jpeg",
        Tags: ["cats","animals"],
        OnlyFiles: "false",
        OnlyFolders: "false",
        PageNo: 1,
        PageSize: 10,
        Sort: "name",
    }
    result, err := pixelbin.Assets.ListFiles(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument    | Type     | Required | Description                                                                  |
| ----------- | -------- | -------- | ---------------------------------------------------------------------------- |
| Name        | string   | no       | Find items with matching name                                                |
| Path        | string   | no       | Find items with matching path                                                |
| Format      | string   | no       | Find items with matching format                                              |
| Tags        | []string | no       | Find items containing these tags                                             |
| OnlyFiles   | bool     | no       | If true will fetch only files                                                |
| OnlyFolders | bool     | no       | If true will fetch only folders                                              |
| PageNo      | float64  | no       | Page No.                                                                     |
| PageSize    | float64  | no       | Page Size                                                                    |
| Sort        | string   | no       | Key to sort results by. A "-" suffix will sort results in descending orders. |

List all files and folders in root folder. Search for files if name is provided. If path is provided, search in the specified path.

_Returned Response:_

[ListFilesResponse](#listfilesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "items": [
        {
            "_id": "dummy-uuid",
            "name": "dir",
            "type": "folder"
        },
        {
            "_id": "dummy-uuid",
            "name": "asset2",
            "type": "file",
            "path": "dir",
            "fileId": "dir/asset2",
            "format": "jpeg",
            "size": 1000,
            "access": "private"
        },
        {
            "_id": "dummy-uuid",
            "name": "asset1",
            "type": "file",
            "path": "dir",
            "fileId": "dir/asset1",
            "format": "jpeg",
            "size": 1000,
            "access": "private"
        }
    ],
    "page": {
        "type": "number",
        "size": 4,
        "current": 1,
        "hasNext": false
    }
}
```

</details>

### GetDefaultAssetForPlayground

**Summary**: Get default asset for playground

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetDefaultAssetForPlaygroundXQuery{
    }
    result, err := pixelbin.Assets.GetDefaultAssetForPlayground(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

Get default asset for playground

_Returned Response:_

[UploadResponse](#uploadresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "isActive": true,
    "orgId": "1",
    "type": "file",
    "name": "abc.jpeg",
    "path": "/xyz",
    "fileId": "xyz/abc.jpeg",
    "format": "jpeg",
    "size": 100,
    "tags": null,
    "metadata": null,
    "access": "public-read",
    "width": null,
    "height": null,
    "meta": {},
    "context": null,
    "assetType": null,
    "isOriginal": true,
    "_id": "35675e3a-5dd8-4b19-a611-1cb64e676c5e",
    "url": "https://cdn.pixelbin.io/v2/dummy-cloudname/original/xyz/abc.jpeg"
}
```

</details>

### GetModules

**Summary**: Get all transformation modules

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetModulesXQuery{
    }
    result, err := pixelbin.Assets.GetModules(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

Get all transformation modules.

_Returned Response:_

[TransformationModulesResponse](#transformationmodulesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "delimiters": {
        "operationSeparator": "~",
        "parameterSeparator": ":"
    },
    "plugins": {
        "erase": {
            "identifier": "erase",
            "name": "EraseBG",
            "description": "EraseBG Background Removal Module",
            "credentials": {
                "required": false
            },
            "operations": [
                {
                    "params": {
                        "name": "Industry Type",
                        "type": "enum",
                        "enum": ["general", "ecommerce"],
                        "default": "general",
                        "identifier": "i",
                        "title": "Industry type"
                    },
                    "displayName": "Remove background of an image",
                    "method": "bg",
                    "description": "Remove the background of any image"
                }
            ],
            "enabled": true
        }
    },
    "presets": [
        {
            "_id": "dummy-id",
            "createdAt": "2022-02-14T10:06:17.803Z",
            "updatedAt": "2022-02-14T10:06:17.803Z",
            "isActive": true,
            "orgId": "265",
            "presetName": "compressor",
            "transformation": "t.compress(q:95)",
            "archived": false
        }
    ]
}
```

</details>

### GetModule

**Summary**: Get Transformation Module by module identifier

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetModuleXQuery{
        Identifier: "t",
    }
    result, err := pixelbin.Assets.GetModule(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument   | Type   | Required | Description                         |
| ---------- | ------ | -------- | ----------------------------------- |
| Identifier | string | yes      | identifier of Transformation Module |

Get Transformation Module by module identifier

_Returned Response:_

[TransformationModuleResponse](#transformationmoduleresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "identifier": "erase",
    "name": "EraseBG",
    "description": "EraseBG Background Removal Module",
    "credentials": {
        "required": false
    },
    "operations": [
        {
            "params": {
                "name": "Industry Type",
                "type": "enum",
                "enum": ["general", "ecommerce"],
                "default": "general",
                "identifier": "i",
                "title": "Industry type"
            },
            "displayName": "Remove background of an image",
            "method": "bg",
            "description": "Remove the background of any image"
        }
    ],
    "enabled": true
}
```

</details>

### AddPreset

**Summary**: Add a preset.

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.AddPresetXQuery{
        PresetName: "pre-set_1",
        Transformation: "t.resize(w:$w,h:$h)~t.extract()",
        Params: map[string]interface{}{"w":{"type":"integer","default":200},"h":{"type":"integer","default":400}},
    }
    result, err := pixelbin.Assets.AddPreset(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument       | Type                   | Required | Description                                    |
| -------------- | ---------------------- | -------- | ---------------------------------------------- |
| PresetName     | string                 | yes      | Name of the preset                             |
| Transformation | string                 | yes      | A chain of transformations, separated by `~`   |
| Params         | map[string]interface{} | no       | Parameters object for transformation variables |

Add a preset for an organization.

_Returned Response:_

[AddPresetResponse](#addpresetresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "orgId": 23,
    "presetName": "pre-set_1",
    "transformation": "t.resize(w:$w,h:$h)~t.extract()",
    "params": {
        "w": {
            "type": "integer",
            "default": 200
        },
        "h": {
            "type": "integer",
            "default": 400
        }
    },
    "_id": "821c6816-3cbb-40fd-8629-0098007fc949",
    "createdAt": "2024-03-21T10:35:47.822Z",
    "updatedAt": "2024-03-21T10:35:47.822Z",
    "isActive": true,
    "archived": false
}
```

</details>

### GetPresets

**Summary**: Get presets for an organization

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetPresetsXQuery{
        PageNo: 1,
        PageSize: 5,
        Name: "t_0",
        Transformation: "t.resize(a:0)",
        Archived: "false",
        Sort: ["updatedAt"],
    }
    result, err := pixelbin.Assets.GetPresets(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument       | Type     | Required | Description                                     |
| -------------- | -------- | -------- | ----------------------------------------------- |
| PageNo         | float64  | no       | Page number                                     |
| PageSize       | float64  | no       | Page size                                       |
| Name           | string   | no       | Preset name                                     |
| Transformation | string   | no       | Transformation applied                          |
| Archived       | bool     | no       | Indicates whether the preset is archived or not |
| Sort           | []string | no       | Sort the results by a specific key              |

Retrieve presets for a specific organization.

_Returned Response:_

[GetPresetsResponse](#getpresetsresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "items": [
        {
            "_id": "f1ae2fc0-a931-4cef-bd1a-3644dad5ae9b",
            "createdAt": "2024-03-21T10:45:06.623Z",
            "updatedAt": "2024-03-21T10:45:06.623Z",
            "isActive": true,
            "orgId": 23,
            "presetName": "t_0",
            "transformation": "t.resize(a:0)",
            "archived": false,
            "params": {}
        },
        {
            "_id": "b40a03f1-7fa5-42b1-8cc6-ffe84c9e6629",
            "createdAt": "2024-03-21T10:45:06.637Z",
            "updatedAt": "2024-03-21T10:45:06.637Z",
            "isActive": true,
            "orgId": 23,
            "presetName": "t_1",
            "transformation": "t.resize(a:1)",
            "archived": false,
            "params": {}
        }
    ],
    "page": {
        "type": "number",
        "size": 2,
        "current": 1,
        "hasNext": true,
        "itemTotal": 10
    }
}
```

</details>

### UpdatePreset

**Summary**: Update a preset.

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.UpdatePresetXQuery{
        PresetName: "p1",,
        Archived: true,
    }
    result, err := pixelbin.Assets.UpdatePreset(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument   | Type   | Required | Description                               |
| ---------- | ------ | -------- | ----------------------------------------- |
| PresetName | string | yes      | Name of the preset to be updated          |
| Archived   | bool   | yes      | Indicates if the preset has been archived |

Update a preset of an organization.

_Returned Response:_

[AddPresetResponse](#addpresetresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "orgId": 23,
    "presetName": "pre-set_1",
    "transformation": "t.resize(w:$w,h:$h)~t.extract()",
    "params": {
        "w": {
            "type": "integer",
            "default": 200
        },
        "h": {
            "type": "integer",
            "default": 400
        }
    },
    "_id": "821c6816-3cbb-40fd-8629-0098007fc949",
    "createdAt": "2024-03-21T10:35:47.822Z",
    "updatedAt": "2024-03-21T10:35:47.822Z",
    "isActive": true,
    "archived": true
}
```

</details>

### DeletePreset

**Summary**: Delete a preset.

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.DeletePresetXQuery{
        PresetName: "pre-set_1",
    }
    result, err := pixelbin.Assets.DeletePreset(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument   | Type   | Required | Description                      |
| ---------- | ------ | -------- | -------------------------------- |
| PresetName | string | yes      | Name of the preset to be deleted |

Delete a preset of an organization.

_Returned Response:_

[AddPresetResponse](#addpresetresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "orgId": 23,
    "presetName": "pre-set_1",
    "transformation": "t.resize(w:$w,h:$h)~t.extract()",
    "params": {
        "w": {
            "type": "integer",
            "default": 200
        },
        "h": {
            "type": "integer",
            "default": 400
        }
    },
    "_id": "821c6816-3cbb-40fd-8629-0098007fc949",
    "createdAt": "2024-03-21T10:35:47.822Z",
    "updatedAt": "2024-03-21T10:35:47.822Z",
    "isActive": true,
    "archived": false
}
```

</details>

### GetPreset

**Summary**: Get a preset.

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.GetPresetXQuery{
        PresetName: "p1",
    }
    result, err := pixelbin.Assets.GetPreset(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument   | Type   | Required | Description                      |
| ---------- | ------ | -------- | -------------------------------- |
| PresetName | string | yes      | Name of the preset to be fetched |

Get a preset of an organization.

_Returned Response:_

[AddPresetResponse](#addpresetresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "orgId": 23,
    "presetName": "p1",
    "transformation": "t.resize(w:$w,h:$h)~t.extract()",
    "params": {
        "w": {
            "type": "integer",
            "default": 200
        },
        "h": {
            "type": "integer",
            "default": 400
        }
    },
    "_id": "821c6816-3cbb-40fd-8629-0098007fc949",
    "createdAt": "2024-03-21T10:35:47.822Z",
    "updatedAt": "2024-03-21T10:35:47.822Z",
    "isActive": true,
    "archived": false
}
```

</details>

### FileUpload

**Summary**: Upload File

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.FileUploadXQuery{
        File: os.Open("your-file-path"),
        Path: "path/to/containing/folder",
        Name: "filename",
        Access: "public-read",
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{},
        Overwrite: false,
        FilenameOverride: true,
    }
    result, err := pixelbin.Assets.FileUpload(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument         | Type                   | Required | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| File             | \*os.File              | yes      | Asset file                                                                                                                                                                                                                       |
| Path             | string                 | no       | Path where you want to store the asset                                                                                                                                                                                           |
| Name             | string                 | no       | Name of the asset, if not provided name of the file will be used. Note - The provided name will be slugified to make it URL safe                                                                                                 |
| Access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| Tags             | []string               | no       | Asset tags                                                                                                                                                                                                                       |
| Metadata         | map[string]interface{} | no       | Asset related metadata                                                                                                                                                                                                           |
| Overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| FilenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

Upload File to Pixelbin

_Returned Response:_

[UploadResponse](#uploadresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### UrlUpload

**Summary**: Upload Asset with url

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.UrlUploadXQuery{
        URL: "www.dummy.com/image.png",
        Path: "path/to/containing/folder",
        Name: "filename",
        Access: "public-read",
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{},
        Overwrite: false,
        FilenameOverride: true,
    }
    result, err := pixelbin.Assets.UrlUpload(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument         | Type                   | Required | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| URL              | string                 | yes      | Asset URL                                                                                                                                                                                                                        |
| Path             | string                 | no       | Path where you want to store the asset                                                                                                                                                                                           |
| Name             | string                 | no       | Name of the asset, if not provided name of the file will be used. Note - The provided name will be slugified to make it URL safe                                                                                                 |
| Access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| Tags             | []string               | no       | Asset tags                                                                                                                                                                                                                       |
| Metadata         | map[string]interface{} | no       | Asset related metadata                                                                                                                                                                                                           |
| Overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| FilenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

Upload Asset with url

_Returned Response:_

[UploadResponse](#uploadresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### CreateSignedUrl

**Summary**: S3 Signed URL upload

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.CreateSignedUrlXQuery{
        Name: "filename",
        Path: "path/to/containing/folder",
        Format: "jpeg",
        Access: "public-read",
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{},
        Overwrite: false,
        FilenameOverride: true,
    }
    result, err := pixelbin.Assets.CreateSignedUrl(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument         | Type                   | Required | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Name             | string                 | no       | name of the file                                                                                                                                                                                                                 |
| Path             | string                 | no       | Path of the file                                                                                                                                                                                                                 |
| Format           | string                 | no       | Format of the file                                                                                                                                                                                                               |
| Access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| Tags             | []string               | no       | Tags associated with the file.                                                                                                                                                                                                   |
| Metadata         | map[string]interface{} | no       | Metadata associated with the file.                                                                                                                                                                                               |
| Overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| FilenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

For the given asset details, a S3 signed URL will be generated,
which can be then used to upload your asset.

_Returned Response:_

[SignedUploadResponse](#signeduploadresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "s3PresignedUrl": {
        "url": "https://domain.com/xyz",
        "fields": {
            "field1": "value",
            "field2": "value"
        }
    }
}
```

</details>

### CreateSignedUrlV2

**Summary**: Signed multipart upload

```golang
import (
    "fmt"
    "os"
    "github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // create pixelbin config object
    config := platform.NewPixelbinConfig(
        "API_TOKEN",
        "https://api.pixelbin.io",
    )
    // set oauthclient
    config.SetOAuthClient()

    // create pixelbin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Parameters for FileUpload function
    params := platform.CreateSignedUrlV2XQuery{
        Name: "filename",
        Path: "path/to/containing/folder",
        Format: "jpeg",
        Access: "public-read",
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{},
        Overwrite: false,
        FilenameOverride: true,
        Expiry: 3000,
    }
    result, err := pixelbin.Assets.CreateSignedUrlV2(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument         | Type                   | Required | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Name             | string                 | no       | name of the file                                                                                                                                                                                                                 |
| Path             | string                 | no       | Path of containing folder.                                                                                                                                                                                                       |
| Format           | string                 | no       | Format of the file                                                                                                                                                                                                               |
| Access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| Tags             | []string               | no       | Tags associated with the file.                                                                                                                                                                                                   |
| Metadata         | map[string]interface{} | no       | Metadata associated with the file.                                                                                                                                                                                               |
| Overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| FilenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |
| Expiry           | float64                | no       | Expiry time in seconds for the signed URL. Defaults to 3000 seconds.                                                                                                                                                             |

For the given asset details, a presigned URL will be generated, which can be then used to upload your asset in chunks via multipart upload.

_Returned Response:_

[SignedUploadV2Response](#signeduploadv2response)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "presignedUrl": {
        "url": "https://api.pixelbin.io/service/public/assets/v1.0/signed-multipart?pbs=8b49e6cdd446be379aa4396e1a&pbe=1700600070390&pbt=92661&pbo=143209&pbu=5fe187e8-8649-4546-9a28-ff551839e0f5",
        "fields": {
            "x-pixb-meta-assetdata": "{\"orgId\":1,\"type\":\"file\",\"name\":\"filename.jpeg\",\"path\":\"\",\"fileId\":\"filename.jpeg\",\"format\":\"jpeg\",\"s3Bucket\":\"erase-erase-erasebg-assets\",\"s3Key\":\"uploads/floral-sun-9617c8/original/a34f1d3-28bf-489c-9aff-cc549ac9e003.jpeg\",\"access\":\"public-read\",\"tags\":[],\"metadata\":{\"source\":\"signedUrl\",\"publicUploadId\":\"5fe187e8-8649-4546-9a28-ff551839e0f5\"},\"overwrite\":false,\"filenameOverride\":false}"
        }
    }
}
```

</details>

### Schemas

#### folderItem

| Properties | Type    | Nullable | Description                          |
| ---------- | ------- | -------- | ------------------------------------ |
| \_id       | string  | no       | Id of the folder item                |
| orgId      | float64 | no       | Organization Id                      |
| name       | string  | no       | Name of the folder item              |
| path       | string  | no       | Path of the folder item              |
| type       | string  | no       | Type of the item. `file` or `folder` |

#### exploreItem

| Properties | Type       | Nullable | Description                                                     |
| ---------- | ---------- | -------- | --------------------------------------------------------------- |
| \_id       | string     | no       | id of the exploreItem                                           |
| orgId      | float64    | no       | Organization Id                                                 |
| name       | string     | no       | name of the item                                                |
| type       | string     | no       | Type of item whether `file` or `folder`                         |
| path       | string     | no       | Path of the folder item                                         |
| fileId     | string     | no       | FileId associated with the item. `path`+`name`                  |
| format     | string     | no       | Format of the file                                              |
| size       | float64    | no       | Size of the file in bytes                                       |
| access     | AccessEnum | no       | Access level of asset, can be either `public-read` or `private` |
| s3Bucket   | string     | no       | Bucket Name                                                     |
| s3Key      | string     | no       | s3 path of file                                                 |

#### page

| Properties | Type    | Nullable | Description                   |
| ---------- | ------- | -------- | ----------------------------- |
| type       | string  | yes      | Type of page                  |
| size       | float64 | yes      | Number of items on the page   |
| current    | float64 | yes      | Current page number.          |
| hasNext    | bool    | yes      | Whether the next page exists. |
| itemTotal  | float64 | yes      | Total number of items.        |

#### exploreResponse

| Properties | Type          | Nullable | Description                  |
| ---------- | ------------- | -------- | ---------------------------- |
| items      | []exploreItem | yes      | exploreItems in current page |
| page       | page          | yes      | page details                 |

#### ListFilesResponse

| Properties | Type          | Nullable | Description                  |
| ---------- | ------------- | -------- | ---------------------------- |
| items      | []exploreItem | no       | exploreItems in current page |
| page       | page          | no       | page details                 |

#### FileUploadRequest

| Properties       | Type                   | Nullable | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| file             | \*os.File              | yes      | Asset file                                                                                                                                                                                                                       |
| path             | string                 | no       | Path where you want to store the asset                                                                                                                                                                                           |
| name             | string                 | no       | Name of the asset, if not provided name of the file will be used. Note - The provided name will be slugified to make it URL safe                                                                                                 |
| access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| tags             | []string               | no       | Asset tags                                                                                                                                                                                                                       |
| metadata         | map[string]interface{} | no       | Asset related metadata                                                                                                                                                                                                           |
| overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| filenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

#### UrlUploadRequest

| Properties       | Type                   | Nullable | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| url              | string                 | yes      | Asset URL                                                                                                                                                                                                                        |
| path             | string                 | no       | Path where you want to store the asset                                                                                                                                                                                           |
| name             | string                 | no       | Name of the asset, if not provided name of the file will be used. Note - The provided name will be slugified to make it URL safe                                                                                                 |
| access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| tags             | []string               | no       | Asset tags                                                                                                                                                                                                                       |
| metadata         | map[string]interface{} | no       | Asset related metadata                                                                                                                                                                                                           |
| overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| filenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

#### UploadResponse

| Properties | Type                   | Nullable | Description                                                 |
| ---------- | ---------------------- | -------- | ----------------------------------------------------------- |
| \_id       | string                 | yes      | \_id of the item                                            |
| fileId     | string                 | yes      | FileId associated with the item. path+name                  |
| name       | string                 | yes      | name of the item                                            |
| path       | string                 | yes      | path to the parent folder                                   |
| format     | string                 | yes      | format of the file                                          |
| size       | float64                | yes      | size of file in bytes                                       |
| access     | AccessEnum             | yes      | Access level of asset, can be either public-read or private |
| tags       | []string               | no       | tags associated with the item                               |
| metadata   | map[string]interface{} | no       | metadata associated with the item                           |
| url        | string                 | no       | url of the item                                             |
| thumbnail  | string                 | no       | url of item thumbnail                                       |

#### SignedUploadRequest

| Properties       | Type                   | Nullable | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name             | string                 | no       | name of the file                                                                                                                                                                                                                 |
| path             | string                 | no       | Path of the file                                                                                                                                                                                                                 |
| format           | string                 | no       | Format of the file                                                                                                                                                                                                               |
| access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| tags             | []string               | no       | Tags associated with the file.                                                                                                                                                                                                   |
| metadata         | map[string]interface{} | no       | Metadata associated with the file.                                                                                                                                                                                               |
| overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| filenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

#### SignedUploadResponse

| Properties     | Type         | Nullable | Description                                  |
| -------------- | ------------ | -------- | -------------------------------------------- |
| s3PresignedUrl | PresignedUrl | yes      | `signedDetails` for upload with frontend sdk |
|  |

#### PresignedUrl

| Properties | Type                   | Nullable | Description                                 |
| ---------- | ---------------------- | -------- | ------------------------------------------- |
| url        | string                 | no       | `presigned url` for upload                  |
|  |
| fields     | map[string]interface{} | no       | signed fields to be sent along with request |

#### FilesResponse

| Properties | Type                   | Nullable | Description                                                    |
| ---------- | ---------------------- | -------- | -------------------------------------------------------------- |
| \_id       | string                 | yes      | \_id of the file                                               |
| name       | string                 | yes      | name of the file                                               |
| path       | string                 | yes      | path to the parent folder of the file                          |
| fileId     | string                 | yes      | FileId associated with the item. `path`+`name`                 |
| format     | string                 | yes      | format of the file                                             |
| size       | float64                | yes      | size of the file in bytes                                      |
| access     | AccessEnum             | yes      | Access level of file, can be either `public-read` or `private` |
| isActive   | bool                   | yes      | Whether the file is active                                     |
| tags       | []string               | no       | Tags associated with the file                                  |
| metadata   | map[string]interface{} | no       | Metadata associated with the file                              |
| url        | string                 | no       | url of the file                                                |
| thumbnail  | string                 | no       | url of the thumbnail of the file                               |

#### UpdateFileRequest

| Properties | Type                   | Nullable | Description                                                     |
| ---------- | ---------------------- | -------- | --------------------------------------------------------------- |
| name       | string                 | no       | Name of the file                                                |
| path       | string                 | no       | Path of the file                                                |
| access     | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private` |
| isActive   | bool                   | no       | Whether the file is active                                      |
| tags       | []string               | no       | Tags associated with the file                                   |
| metadata   | map[string]interface{} | no       | Metadata associated with the file                               |

#### FoldersResponse

| Properties | Type   | Nullable | Description                             |
| ---------- | ------ | -------- | --------------------------------------- |
| \_id       | string | yes      | \_id of the folder                      |
| name       | string | yes      | name of the folder                      |
| path       | string | yes      | path to the parent folder of the folder |
| isActive   | bool   | yes      | whether the folder is active            |

#### CreateFolderRequest

| Properties | Type   | Nullable | Description        |
| ---------- | ------ | -------- | ------------------ |
| name       | string | yes      | Name of the folder |
| path       | string | no       | Path of the folder |

#### UpdateFolderRequest

| Properties | Type | Nullable | Description                  |
| ---------- | ---- | -------- | ---------------------------- |
| isActive   | bool | no       | whether the folder is active |

#### DeleteMultipleFilesRequest

| Properties | Type     | Nullable | Description                   |
| ---------- | -------- | -------- | ----------------------------- |
| ids        | []string | yes      | Array of file \_ids to delete |

#### Delimiter

| Properties         | Type   | Nullable | Description                                                              |
| ------------------ | ------ | -------- | ------------------------------------------------------------------------ |
| operationSeparator | string | no       | separator to separate operations in the url pattern                      |
| parameterSeparator | string | no       | separator to separate parameters used with operations in the url pattern |

#### AddCredentialsRequest

| Properties  | Type                   | Nullable | Description                                                 |
| ----------- | ---------------------- | -------- | ----------------------------------------------------------- |
| credentials | map[string]interface{} | yes      | Credentials of the plugin                                   |
| pluginId    | string                 | yes      | Unique identifier for the plugin this credential belongs to |

#### UpdateCredentialsRequest

| Properties  | Type                   | Nullable | Description               |
| ----------- | ---------------------- | -------- | ------------------------- |
| credentials | map[string]interface{} | yes      | Credentials of the plugin |

#### AddCredentialsResponse

| Properties  | Type                   | Nullable | Description |
| ----------- | ---------------------- | -------- | ----------- |
| credentials | map[string]interface{} | no       |             |

#### GetAncestorsResponse

| Properties | Type              | Nullable | Description |
| ---------- | ----------------- | -------- | ----------- |
| folder     | folderItem        | no       |             |
| ancestors  | []FoldersResponse | no       |             |

#### AddPresetRequest

| Properties     | Type                   | Nullable | Description                                    |
| -------------- | ---------------------- | -------- | ---------------------------------------------- |
| presetName     | string                 | yes      | Name of the preset                             |
| transformation | string                 | yes      | A chain of transformations, separated by `~`   |
| params         | map[string]interface{} | no       | Parameters object for transformation variables |

#### AddPresetResponse

| Properties     | Type                   | Nullable | Description                                    |
| -------------- | ---------------------- | -------- | ---------------------------------------------- |
| presetName     | string                 | no       | Name of the preset                             |
| transformation | string                 | no       | A chain of transformations, separated by `~`   |
| params         | map[string]interface{} | no       | Parameters object for transformation variables |
| archived       | bool                   | no       | Indicates if the preset has been archived      |
| orgId          | float64                | no       | Organization Id                                |
| isActive       | bool                   | no       | Indicates if the preset is active              |
| createdAt      | string                 | no       | Preset creation ISO timestamp                  |
| updatedAt      | string                 | no       | Preset update ISO timestamp                    |

#### UpdatePresetRequest

| Properties | Type | Nullable | Description                               |
| ---------- | ---- | -------- | ----------------------------------------- |
| archived   | bool | yes      | Indicates if the preset has been archived |

#### GetPresetsResponse

| Properties | Type                | Nullable | Description  |
| ---------- | ------------------- | -------- | ------------ |
| items      | []AddPresetResponse | yes      |              |
| page       | page                | yes      | page details |

#### TransformationModuleResponse

| Properties  | Type                   | Nullable | Description                                     |
| ----------- | ---------------------- | -------- | ----------------------------------------------- |
| identifier  | string                 | no       | identifier for the plugin type                  |
| name        | string                 | no       | name of the plugin                              |
| description | string                 | no       | description of the plugin                       |
| credentials | map[string]interface{} | no       | credentials, if any, associated with the plugin |
| operations  | []interface{}          | no       | supported operations in the plugin              |
| enabled     | bool                   | no       | whether the plugin is enabled                   |

#### TransformationModulesResponse

| Properties | Type                                    | Nullable | Description                                         |
| ---------- | --------------------------------------- | -------- | --------------------------------------------------- |
| delimiters | Delimiter                               | no       | Delimiter for parsing plugin schema                 |
| plugins    | map[string]TransformationModuleResponse | no       | Transformations currently supported by the pixelbin |
| presets    | []interface{}                           | no       | List of saved presets                               |

#### SignedUploadRequestV2

| Properties       | Type                   | Nullable | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name             | string                 | no       | name of the file                                                                                                                                                                                                                 |
| path             | string                 | no       | Path of containing folder.                                                                                                                                                                                                       |
| format           | string                 | no       | Format of the file                                                                                                                                                                                                               |
| access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| tags             | []string               | no       | Tags associated with the file.                                                                                                                                                                                                   |
| metadata         | map[string]interface{} | no       | Metadata associated with the file.                                                                                                                                                                                               |
| overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| filenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |
| expiry           | float64                | no       | Expiry time in seconds for the signed URL. Defaults to 3000 seconds.                                                                                                                                                             |

#### SignedUploadV2Response

| Properties   | Type           | Nullable | Description                                 |
| ------------ | -------------- | -------- | ------------------------------------------- |
| presignedUrl | PresignedUrlV2 | yes      | Presigned URL for uploading asset in chunks |

#### PresignedUrlV2

| Properties | Type              | Nullable | Description                                 |
| ---------- | ----------------- | -------- | ------------------------------------------- |
| url        | string            | no       | Presigned URL for uploading asset in chunks |
| fields     | map[string]string | no       | signed fields to be sent along with request |

### Enums

#### [AccessEnum](#AccessEnum)

Type : string

| Name        | Value       | Description |
| ----------- | ----------- | ----------- |
| public-read | public-read | public-read |
| private     | private     | private     |

---
