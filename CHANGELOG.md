# 3.1.0

-   Added [upload](./README.md#upload) method.
-   Removed support for signature from all APIs.

# 3.0.0

-   **Breaking Changes:**
    -   Function signature of `SignURL` in `security` package has been updated. The argument `tokenID` has been replaced with `accessKey`. Access key for a token can be found by going to console.pixelbin.io > `Settings` > `Tokens` > Click on any Token > `Access Key`.
    -   Package name has been updated to reflect the url where its located now
        `github.com/pixelbin-dev/pixelbin-go` => `github.com/pixelbin-io/pixelbin-go`

# 2.4.0

-   Added method for generating v2 Signed Multipart Upload Urls `CreateSignedUrlV2`

# 2.3.0

-   Added support for generating signed Custom Domain and PixelBin CDN urls

# 2.2.0

-   Added worker and Custom Domain url support to `UrlToObj` & `ObjToUrl`

# 2.1.0

-   Added a method for obtaining the context of a file via url.

# 2.0.0

-   Fixed `tags` being stringified inadvertently. If you are experiencing validation errors around `tags` in previous versions, you should upgrade your SDKs.
-   Fixed non-string fields that were not working, such as `Overwrite` (boolean) in FileUpload.
-   method for getting org details has changed from `getAppByToken` => `getAppOrgDetails`
