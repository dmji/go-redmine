# Journals


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ID`                                                       | *int64*                                                    | :heavy_check_mark:                                         | N/A                                                        |
| `User`                                                     | [components.IDName](../../models/components/idname.md)     | :heavy_check_mark:                                         | N/A                                                        |
| `Notes`                                                    | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |
| `CreatedOn`                                                | [time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                         | N/A                                                        |
| `PrivateNotes`                                             | *bool*                                                     | :heavy_check_mark:                                         | N/A                                                        |
| `Details`                                                  | [][components.Details](../../models/components/details.md) | :heavy_check_mark:                                         | N/A                                                        |