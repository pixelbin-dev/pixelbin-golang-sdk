# Pixelbin Backend SDK for Golang

Pixelbin Backend SDK for Golang helps you integrate the core Pixelbin features with your application.

## Getting Started

Getting started with Pixelbin Backend SDK for Golang

### Installation

```
go get -u "github.com/pixelbin-io/pixelbin-go/v3"
```

---

### Usage

#### Quick Example

```go
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

    file, _ := os.Open("/home/rohit/deidara/1.jpeg")

    // Parameters for FileUpload function
    params := platform.FileUploadXQuery{
        File: file,
    }
    result, err := pixelbin.Assets.FileUpload(params)

    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(result)
}
```

## Uploader

### Upload

Uploads a file to PixelBin with greater control over the upload process.

#### Arguments

| Argument | Type                   | Required | Description                                                                                                                                                 |
| -------- | ---------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `file`   | `io.Reader`            | yes      | The file to be uploaded. It can be any type that implements the `io.Reader` interface, such as an open file or a buffer.                                    |
| `p`      | `UploaderUploadXQuery` | yes      | parameters for the upload, including file name, path, format, access level, and more.                                                                       |
| `opts`   | `uploaderOption...`    | no       | Variadic option functions that allow customization of the upload process, such as setting chunk size, maximum retries, concurrency, and exponential factor. |

#### `UploaderUploadXQuery` Struct

| Field              | Type                                                        | Required | Description                                                                                                                                                                                                                              |
| ------------------ | ----------------------------------------------------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Name`             | `string`                                                    | no       | Name of the file.                                                                                                                                                                                                                        |
| `Path`             | `string`                                                    | no       | Path of the containing folder.                                                                                                                                                                                                           |
| `Format`           | `string`                                                    | no       | Format of the file.                                                                                                                                                                                                                      |
| `Access`           | [AccessEnum](./documentation/platform/ASSETS.md#accessenum) | no       | Access level of the asset, can be either `public-read` or `private`.                                                                                                                                                                     |
| `Tags`             | `[]string`                                                  | no       | Tags associated with the file.                                                                                                                                                                                                           |
| `Metadata`         | `map[string]interface{}`                                    | no       | Metadata associated with the file.                                                                                                                                                                                                       |
| `Overwrite`        | `bool`                                                      | no       | Overwrite flag. If set to `true`, will overwrite any file that exists with the same path, name, and type. Defaults to `false`.                                                                                                           |
| `FilenameOverride` | `bool`                                                      | no       | If set to `true`, will add unique characters to the name if an asset with the given name already exists. If `Overwrite` is also set to `true`, preference will be given to `Overwrite`. If both are set to `false`, an error will occur. |
| `Expiry`           | `float64`                                                   | no       | Expiry time in seconds for the underlying signed URL. Defaults to 3000 seconds.                                                                                                                                                          |

#### Uploader Options

These options can be passed as variadic arguments to fine-tune the upload process:

-   **`WithChunkSize(size uint)`**: Set the size of each chunk to upload. Default is 10 megabytes.
-   **`WithMaxRetries(retries uint)`**: Set the maximum number of retries if an upload fails. Default is 2 retries.
-   **`WithConcurrency(concurrency uint)`**: Set the number of concurrent chunk upload tasks. Default is 3 concurrent chunk uploads.
-   **`WithExponentialFactor(factor uint)`**: Set the exponential factor for retry delay. Default is 2.

#### Returns

-   **On Success**: `map[string]interface{}` containing details about the uploaded file, such as `url`, `name`, `format`, `tags`, and `metadata`.

Sure, here’s the updated "On Success" table using Go datatypes:

#### On Success

| Property     | Type                     | Description                                                       | Example                                                           |
| ------------ | ------------------------ | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `orgId`      | `int`                    | Organization ID.                                                  | `5320086`                                                         |
| `type`       | `string`                 | The type of asset. Always `"file"`.                               | `"file"`                                                          |
| `name`       | `string`                 | Name of the file.                                                 | `"testfile.jpeg"`                                                 |
| `path`       | `string`                 | Path of the containing folder.                                    | `"/path/to/image.jpeg"`                                           |
| `fileId`     | `string`                 | Unique ID of the file.                                            | `"testfile.jpeg"`                                                 |
| `access`     | `string`                 | Access level of the asset, either `"public-read"` or `"private"`. | `"public-read"`                                                   |
| `tags`       | `[]string`               | Tags associated with the file.                                    | `[]string{"tag1", "tag2"}`                                        |
| `metadata`   | `map[string]interface{}` | Metadata associated with the file.                                | `map[string]interface{}{"source": "", "publicUploadId": ""}`      |
| `format`     | `string`                 | File format.                                                      | `"jpeg"`                                                          |
| `assetType`  | `string`                 | Type of asset, e.g., `"image"`.                                   | `"image"`                                                         |
| `size`       | `int64`                  | File size in bytes.                                               | `37394`                                                           |
| `width`      | `int`                    | Width of the file (if applicable).                                | `720`                                                             |
| `height`     | `int`                    | Height of the file (if applicable).                               | `450`                                                             |
| `context`    | `map[string]interface{}` | Contains file metadata and other context information.             | `map[string]interface{}{"steps": [], "meta": { ... }}`            |
| `isOriginal` | `bool`                   | Indicates if the file is the original.                            | `true`                                                            |
| `_id`        | `string`                 | Record ID of the uploaded file.                                   | `"a0b0b19a-d526-4xc07-ae51-0xxxxxx"`                              |
| `url`        | `string`                 | URL of the uploaded file.                                         | `"https://cdn.pixelbin.io/v2/user-e26cf3/original/testfile.jpeg"` |

-   **On Error**: An `error` describing what went wrong during the upload process.

#### Uploading a buffer

```go
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // Create PixelBin config object
    config := platform.NewPixelbinConfig("API_TOKEN", "https://api.pixelbin.io")
    config.SetOAuthClient()

    // Create PixelBin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Open the file to be uploaded
    file, err := os.ReadFile("./path/to/your/file.png")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    buffer := bytes.NewReader(file)

    // Define the upload parameters
    params := platform.UploaderUploadXQuery{
        Name:    "myimage",
        Path:    "folder",
        Format:  "jpeg",
        Access:  "public-read",
		Overwrite:        true,
        FilenameOverride: false,
        Expiry:           3600, // 1 hour
    }

    // Upload the file with custom options
    result, err := pixelbin.Uploader.Upload(buffer, params,
        platform.WithChunkSize(5*1024*1024),  // 5MB
        platform.WithMaxRetries(3),
        platform.WithConcurrency(2),
        platform.WithExponentialFactor(2),
    )

    if err != nil {
        fmt.Println("Error uploading file:", err)
        return
    }

    // Print the result
    fmt.Println("File uploaded successfully:", result["url"])
}
```

#### Uploading a stream

```go
package main

import (
	"fmt"
	"os"

	"github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
    // Create PixelBin config object
    config := platform.NewPixelbinConfig("API_TOKEN", "https://api.pixelbin.io")
    config.SetOAuthClient()

    // Create PixelBin client object
    pixelbin := platform.NewPixelbinClient(config)

    // Open the file to be uploaded
    file, err := os.Open("./path/to/your/file.png")

    if err != nil {
        fmt.Println("Error opening file:", err)
        return;
    }

    // Define the upload parameters
    params := platform.UploaderUploadXQuery{
        Name:    "myimage",
        Path:    "folder",
        Format:  "jpeg",
        Access:  "public-read",
		Overwrite:        true,
        FilenameOverride: false,
        Expiry:           3600, // 1 hour
    }

    // Upload the file with custom options
    result, err := pixelbin.Uploader.Upload(file, params,
        platform.WithChunkSize(5*1024*1024),  // 5MB
        platform.WithMaxRetries(3),
        platform.WithConcurrency(2),
        platform.WithExponentialFactor(2),
    )

    if err != nil {
        fmt.Println("Error uploading file:", err)
        return
    }

    // Print the result
    fmt.Println("File uploaded successfully:", result["url"])
}
```

## Security Utils

### For generating Signed URLs

Generate a signed PixelBin url

| Parameter             | Description                                          | Example                                                                                    |
| --------------------- | ---------------------------------------------------- | ------------------------------------------------------------------------------------------ |
| `url` (string)        | A valid Pixelbin URL to be signed                    | `https://cdn.pixelbin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg` |
| `expirySeconds` (int) | Number of seconds the signed URL should be valid for | `20`                                                                                       |
| `accessKey` (string)  | Access key of the token used for signing             | `a45e52d8-21ac-4a97-bd4f-eb5dd58602e0`                                                     |
| `token` (string)      | Value of the token used for signing                  | `dummy-token`                                                                              |

Example:

```golang
package main

import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/security"
)

func main() {
    signedUrl, err := security.SignURL(
        "https://cdn.pixelbin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg", // urlString
        20, // expirySeconds
        "a45e52d8-21ac-4a97-bd4f-eb5dd58602e0", // accessKey
        "dummy-token", // token
    )
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(signedUrl)
}
// signed_url
// https://cdn.pixelbin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg?pbs=8eb6a00af74e57967a42316e4de238aa88d92961649764fad1832c1bff101f25&pbe=1695635915&pbt=a45e52d8-21ac-4a97-bd4f-eb5dd58602e0
```

Usage with custom domain url

```golang
package main

import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/security"
)

func main() {
    signedUrl, err := security.SignUrl(
        "https://krit.imagebin.io/v2/original/__playground/playground-default.jpeg", // url
        30, // expirySeconds
        "ab110791-db9f-4dca-ac39-d29db5941daa", // accessKey
        "dummy-token", // token
    )
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(signedUrl)
}
// signedUrl
// https://krit.imagebin.io/v2/original/__playground/playground-default.jpeg?pbs=1aef31c1e0ecd8a875b1d3184f324327f4ab4bce419d81d1eb1a818ee5f2e3eb&pbe=1695705975&pbt=ab110791-db9f-4dca-ac39-d29db5941daa
```

## URL Utils

Pixelbin provides url utilities to construct and deconstruct Pixelbin urls.

### UrlToObj

Deconstruct a pixelbin URL

| Parameter              | Description                                                | Example                                                                                               |
| ---------------------- | ---------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `pixelbinUrl` (string) | A valid pixelbin URL                                       | `https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/t.resize(h:100,w:200)~t.flip()/path/to/image.jpeg` |
| `opts` (variadic)      | Functional options for configuring the function (optional) | See `UrlToObjOption` below                                                                            |

**`UrlToObjOption`**:

`UrlToObjOption` is a functional option for configuring the `UrlToObj` function. You can use it to customize the behavior of the function by setting different options. See the table below for a list of available options.

**Options**:

| Option             | Description                               | Default Value |
| ------------------ | ----------------------------------------- | ------------- |
| `WithCustomDomain` | Set `IsCustomDomain` to `true` or `false` | `false`       |

**Returns**:

| Property                  | Description                                          | Example                               |
| ------------------------- | ---------------------------------------------------- | ------------------------------------- |
| `baseURL` (string)        | Base path of the URL                                 | `https://cdn.pixelbin.io`             |
| `filePath` (string)       | Path to the file on Pixelbin storage                 | `/path/to/image.jpeg`                 |
| `version` (string)        | Version of the URL                                   | `v2`                                  |
| `cloudName` (string)      | Cloud name from the URL                              | `your-cloud-name`                     |
| `transformations` (array) | A list of transformation objects                     | `[{ "plugin": "t", "name": "flip" }]` |
| `zone` (string)           | Zone slug from the URL                               | `z-slug`                              |
| `pattern` (string)        | Transformation pattern extracted from the URL        | `t.resize(h:100,w:200)~t.flip()`      |
| `worker` (boolean)        | Indicates if the URL is a URL Translation Worker URL | `False`                               |
| `workerPath` (string)     | Input path to a URL Translation Worker               | `resize:w200,h400/folder/image.jpeg`  |
| `options` (Object)        | Query parameters added, such as "dpr" and "f_auto"   | `{ dpr: 2.5, f_auto: True}`           |

Example:

```golang
package main

import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)

func main() {
    pixelbinUrl := "https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/t.resize(h:100,w:200)~t.flip()/path/to/image.jpeg?dpr=2.0&f_auto=true"
    obj, err := url.UrlToObj(pixelbinUrl)
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(obj)
}
// obj
// {
//     "cloudName": "your-cloud-name",
//     "zone": "z-slug",
//     "version": "v2",
//     "options": {
//         "dpr": "2.0",
//         "f_auto": "true",
//     },
//     "transformations": [
//         {
//             "plugin": "t",
//             "name": "resize",
//             "values": [
//                 {
//                     "key": "h",
//                     "value": "100"
//                 },
//                 {
//                     "key": "w",
//                     "value": "200"
//                 }
//             ]
//         },
//         {
//             "plugin": "t",
//             "name": "flip",
//         }
//     ],
//     "filePath": "path/to/image.jpeg",
//     "baseUrl": "https://cdn.pixelbin.io"
// }
```

Usage with custom domain

```golang
package main

import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)

func main() {
    customDomainUrl :=
        "https://xyz.designify.media/v2/z-slug/t.resize(h:100,w:200)~t.flip()/path/to/image.jpeg"
    obj, err := url.UrlToObj(customDomainUrl, url.WithCustomDomain(true))
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(obj)
}
// obj
// {
//     "zone": "z-slug",
//     "version": "v2",
//     "transformations": [
//         {
//             "plugin": "t",
//             "name": "resize",
//             "values": [
//                 {
//                     "key": "h",
//                     "value": "100"
//                 },
//                 {
//                     "key": "w",
//                     "value": "200"
//                 }
//             ]
//         },
//         {
//             "plugin": "t",
//             "name": "flip",
//         }
//     ],
//     "filePath": "path/to/image.jpeg",
//     "baseUrl": "https://xyz.designify.media",
//     "wrkr": False,
//     "workerPath": "",
//     "options": {}
// }
```

Usage with URL Translation Worker

```golang
package main

import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)

func main() {
    workerUrl :=
        "https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/wrkr/resize:h100,w:200/folder/image.jpeg";
    obj, err := url.UrlToObj(workerUrl)
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(obj)
}
// obj
// {
//     "cloudName": "your-cloud-name",
//     "zone": "z-slug",
//     "version": "v2",
//     "transformations": [],
//     "filePath": "",
//     "worker": True,
//     "workerPath": "resize:h100,w:200/folder/image.jpeg",
//     "baseUrl": "https://cdn.pixelbin.io"
//     "options": {}
// }
```

### ObjToUrl

Converts the extracted url obj to a Pixelbin url.

| Property                   | Description                                          | Example                               |
| -------------------------- | ---------------------------------------------------- | ------------------------------------- |
| `cloudName` (string)       | The cloudname extracted from the URL                 | `your-cloud-name`                     |
| `zone` (string)            | 6 character zone slug                                | `z-slug`                              |
| `version` (string)         | CDN API version                                      | `v2`                                  |
| `transformations` (array)  | Extracted transformations from the URL               | `[{ "plugin": "t", "name": "flip" }]` |
| `filePath` (string)        | Path to the file on Pixelbin storage                 | `/path/to/image.jpeg`                 |
| `baseUrl` (string)         | Base URL                                             | `https://cdn.pixelbin.io/`            |
| `isCustomDomain` (boolean) | Indicates if the URL is for a custom domain          | `False`                               |
| `worker` (boolean)         | Indicates if the URL is a URL Translation Worker URL | `False`                               |
| `workerPath` (string)      | Input path to a URL Translation Worker               | `resize:w200,h400/folder/image.jpeg`  |
| `options` (Object)         | Query parameters added, such as "dpr" and "f_auto"   | `{ "dpr": 2.0, "f_auto": True }`      |

```golang
package main

import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)
func main() {
    obj := map[string]interface{}{
        cloudName: "your-cloud-name",
        zone: "z-slug",
        version: "v2",
        options:  []map[string]interface{}{
            dpr: 2.5,
            f_auto: true,
        },
        transformations: []map[string]interface{}{
            {
                plugin: "t",
                name: "flop",
            },
            {
                plugin: "t",
                name: "flip",
            },
        },
        filePath: "path/to/image.jpeg",
        baseUrl: "https://cdn.pixelbin.io",
    }
    urlstring, err := url.ObjToUrl(obj) // obj is as shown above
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(urlstring)
}
// urlstring
// https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/t.flop()~t.flip()/path/to/image.jpeg?dpr=2.5&f_auto=true
```

Usage with custom domain

```golang
package main

import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)
func main() {
    obj := map[string]interface{}{
        cloudName: "your-cloud-name",
        zone: "z-slug",
        version: "v2",
        options:  []map[string]interface{}{
            dpr: 2.5,
            f_auto: true,
        },
        transformations: []map[string]interface{}{
            {
                plugin: "t",
                name: "flop",
            },
            {
                plugin: "t",
                name: "flip",
            },
        },
        filePath: "path/to/image.jpeg",
        baseUrl: "https://xyz.designify.media",
        isCustomDomain: True,
    }
    urlstring, err := url.ObjToUrl(obj) // obj is as shown above
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(urlstring)
}
// urlstring
// https://xyz.designify.media/v2/z-slug/t.flop()~t.flip()/path/to/image.jpeg?dpr=2.5&f_auto=true
```

Usage with URL Translation Worker

```golang
package main

import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)
func main() {
    obj := map[string]interface{}{
        cloudName: "your-cloud-name",
        zone: "z-slug",
        version: "v2",
        options: []map[string]interface{}{
            dpr: 2.5,
            f_auto: true,
        },
        transformations: []map[string]interface{}{},
        worker: true,
        workerPath: "resize:h100,w:200/folder/image.jpeg",
        filePath: "path/to/image.jpeg",
        baseUrl: "https://cdn.pixelbin.io",
    }
    urlstring, err := url.ObjToUrl(obj) // obj is as shown above
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(urlstring)
}
// urlstring
// https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/wrkr/resize:h100,w:200/folder/image.jpeg
```

## Documentation

-   [API docs](documentation/platform/README.md)
