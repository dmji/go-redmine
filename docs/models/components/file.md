# File


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ID`                                                    | *int64*                                                 | :heavy_check_mark:                                      | N/A                                                     |
| `Filename`                                              | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `Filesize`                                              | *int64*                                                 | :heavy_check_mark:                                      | N/A                                                     |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `Description`                                           | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `ContentURL`                                            | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `ThumbnailURL`                                          | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     |
| `Author`                                                | [components.IDName](../../models/components/idname.md)  | :heavy_check_mark:                                      | N/A                                                     |
| `CreatedOn`                                             | [time.Time](https://pkg.go.dev/time#Time)               | :heavy_check_mark:                                      | N/A                                                     |
| `Version`                                               | [*components.IDName](../../models/components/idname.md) | :heavy_minus_sign:                                      | N/A                                                     |
| `Digest`                                                | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `Downloads`                                             | *int64*                                                 | :heavy_check_mark:                                      | N/A                                                     |