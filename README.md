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
    }
    fmt.Println(result)
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
    }
    fmt.Println(signedUrl)
}
// signed_url
// https://cdn.pixelbin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg?pbs=8eb6a00af74e57967a42316e4de238aa88d92961649764fad1832c1bff101f25&pbe=1695635915&pbt=a45e52d8-21ac-4a97-bd4f-eb5dd58602e0
```

Usage with custom domain url

```golang
import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/security"
)

func main() {
    signedUrl := security.SignUrl(
        "https://krit.imagebin.io/v2/original/__playground/playground-default.jpeg", // url
        30, // expirySeconds
        "ab110791-db9f-4dca-ac39-d29db5941daa", // accessKey
        "dummy-token", // token
    )
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
import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)

func main() {
    pixelbinUrl := "https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/t.resize(h:100,w:200)~t.flip()/path/to/image.jpeg?dpr=2.0&f_auto=true"
    obj := url.UrlToObj(pixelbinUrl)
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
import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)

func main() {
    customDomainUrl :=
        "https://xyz.designify.media/v2/z-slug/t.resize(h:100,w:200)~t.flip()/path/to/image.jpeg"
    obj := url.UrlToObj(pixelbinUrl, WithCustomDomain(true))
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
import (
	"fmt"
	"os"
	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)

func main() {
    workerUrl :=
        "https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/wrkr/resize:h100,w:200/folder/image.jpeg";
    obj := url.UrlToObj(pixelbinUrl)
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
    urlstring := url.ObjToUrl(obj) // obj is as shown above
}
// urlstring
// https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/t.flop()~t.flip()/path/to/image.jpeg?dpr=2.5&f_auto=true
```

Usage with custom domain

```golang
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
    urlstring := url.ObjToUrl(obj) // obj is as shown above
}
// urlstring
// https://xyz.designify.media/v2/z-slug/t.flop()~t.flip()/path/to/image.jpeg?dpr=2.5&f_auto=true
```

Usage with URL Translation Worker

```golang
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
    urlstring := url.ObjToUrl(obj) // obj is as shown above
}
// urlstring
// https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/wrkr/resize:h100,w:200/folder/image.jpeg
```

## Documentation

-   [API docs](documentation/platform/README.md)
