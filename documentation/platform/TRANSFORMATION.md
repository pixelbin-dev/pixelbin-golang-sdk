##### [Back to Pixelbin API docs](./README.md)

## Transformation Methods

Image Transformation Service

-   [GetTransformationContext](#gettransformationcontext)

## Methods with example and description

### GetTransformationContext

**Summary**: Get transformation context

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
    params := platform.GetTransformationContextXQuery{
        URL: "/v2/fynd-eg/t.resize()/__playground/playground-default.jpeg",
    }
    result, err := pixelbin.Transformation.GetTransformationContext(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                  |
| -------- | ------ | -------- | ---------------------------- |
| URL      | string | no       | CDN URL with transformation. |

Get transformation context

_Returned Response:_

[GetTransformationContextSuccessResponse](#gettransformationcontextsuccessresponse)

Success

<details>
<summary><i>&nbsp; Examples:</i></summary>

<details>
<summary><i>&nbsp; GetTransformationContextSuccessResponse</i></summary>

```json
{
    "value": {
        "context": {
            "steps": [
                {
                    "name": "blur",
                    "operation": "Basic",
                    "identifier": "t",
                    "data": {},
                    "metadata": null,
                    "format": "jpeg",
                    "size": 58650,
                    "width": 1140,
                    "height": 760,
                    "space": "srgb",
                    "channels": 3,
                    "depth": "uchar",
                    "density": 72,
                    "chromaSubsampling": "4:2:0",
                    "isProgressive": false,
                    "resolutionUnit": "inch",
                    "hasProfile": false,
                    "hasAlpha": false,
                    "orientation": 1
                }
            ],
            "metadata": {
                "width": 1140,
                "height": 760,
                "channels": 3,
                "extension": "jpeg",
                "format": "jpeg",
                "contentType": "image/jpeg",
                "size": 62667,
                "assetType": "image",
                "isImageAsset": true,
                "isAudioAsset": false,
                "isVideoAsset": false,
                "isRawAsset": false
            },
            "headers": {
                "host": "api.pixelbin.io",
                "x-real-ip": "0.0.0.0",
                "origin": "https://console.pixelbin.io",
                "accept": "application/json, text/plain, */*",
                "accept-encoding": "gzip, deflate, br",
                "accept-language": "en-US,en;q=0.9"
            },
            "params": {}
        }
    }
}
```

</details>

</details>

### Schemas

#### GetTransformationContextSuccessResponse

| Properties | Type                   | Nullable | Description |
| ---------- | ---------------------- | -------- | ----------- |
| context    | map[string]interface{} | no       |             |
