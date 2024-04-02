## Migration Guide:

### Pixelbin SDK v2.x.x to v3.x.x

#### Breaking Changes

1. **Updated Function Signature in `security` Package**

    The arguments for `SignURL` function has been updated. The `tokenID` parameter has been replaced with `accessKey`.

    **Previous Method (v2.3.x and above):**

    ```go
    import (
        "fmt"
        "github.com/pixelbin-dev/pixelbin-go/v2/sdk/utils/security"
    )

    func main() {
        signedUrl, err := security.SignURL(
            "https://cdn.pixelbin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg",
            20, // expirySeconds
            42, // tokenID
            "dummy-token", // token
        )
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(signedUrl)
    }
    ```

    **New Method (v3.x.x):**

    ```go
    import (
        "fmt"
        "github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/security"
    )

    func main() {
        signedUrl, err := security.SignURL(
            "https://cdn.pixelbin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg",
            20, // expirySeconds
            "a45e52d8-21ac-4a97-bd4f-eb5dd58602e0", // accessKey
            "dummy-token", // token
        )
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(signedUrl)
    }
    ```

    **How to Find Your Access Key:**

    - Go to `console.pixelbin.io`
    - Navigate to `Settings` > `Tokens`
    - Click on any Token
    - Find the `Access Key` in the token details

2. **Package Name Update**

    The package name has been updated to reflect its new location:

    - **From:** `github.com/pixelbin-dev/pixelbin-go`
    - **To:** `github.com/pixelbin-io/pixelbin-go`

    Update your project's imports to the new package name to ensure compatibility with version 3.x.x.
