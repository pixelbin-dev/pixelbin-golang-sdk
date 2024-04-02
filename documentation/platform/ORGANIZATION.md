##### [Back to Pixelbin API docs](./README.md)

## Organization Methods

Organization Service

-   [GetAppOrgDetails](#getapporgdetails)

## Methods with example and description

### GetAppOrgDetails

**Summary**: Get App Details

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
    params := platform.GetAppOrgDetailsXQuery{
    }
    result, err := pixelbin.Organization.GetAppOrgDetails(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

Get App and org details

_Returned Response:_

[AppOrgDetails](#apporgdetails)

Success. Returns a JSON object as shown below. Refer `AppOrgDetails` for more details.

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "app": {
        "_id": 123,
        "orgId": 12,
        "name": "Desktop Client App",
        "permissions": ["read", "read_write"],
        "active": false,
        "createdAt": "2021-07-15T07:47:00Z",
        "updatedAt": "2021-07-15T07:47:00Z"
    },
    "org": {
        "_id": 12,
        "name": "org_1",
        "cloudName": "testcloudname",
        "accountType": "individual",
        "industry": "Ecommerce",
        "strength": "1",
        "active": "false"
    }
}
```

</details>

### Schemas

#### OrganizationDetailSchema

| Properties | Type    | Nullable | Description                                        |
| ---------- | ------- | -------- | -------------------------------------------------- |
| \_id       | float64 | no       | Id of the organization                             |
| name       | string  | no       | Organization Name                                  |
| cloudName  | string  | no       | Unique cloud name associated with the organization |
| ownerId    | string  | no       | User Id of the organization owner                  |
| active     | bool    | no       | Whether the organization is active                 |
| createdAt  | string  | no       | Timestamp when the organization was created        |
| updatedAt  | string  | no       | Timestamp when the organization was last updated   |

#### AppSchema

| Properties  | Type     | Nullable | Description                             |
| ----------- | -------- | -------- | --------------------------------------- |
| \_id        | float64  | no       | App id                                  |
| orgId       | float64  | no       | Organization id the app belongs to      |
| name        | string   | no       | Name of the app                         |
| token       | string   | no       | api token for the app                   |
| permissions | []string | no       | Permissions applied on the app          |
| active      | bool     | no       | Whether the app is active               |
| createdAt   | string   | no       | Timestamp when the app was created      |
| updatedAt   | string   | no       | Timestamp when the app was last updated |

#### AppOrgDetails

| Properties | Type                     | Nullable | Description |
| ---------- | ------------------------ | -------- | ----------- |
| app        | AppSchema                | no       |             |
| org        | OrganizationDetailSchema | no       |             |

#### ErrorSchema

| Properties | Type   | Nullable | Description |
| ---------- | ------ | -------- | ----------- |
| message    | string | no       |             |
