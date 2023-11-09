# LibraryItem

Represents a specific library item that is part of this project.


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `CompletionDate`                                     | **string*                                            | :heavy_minus_sign:                                   | Date the library item was completed.                 |
| `Description`                                        | **string*                                            | :heavy_minus_sign:                                   | Description of the library item.                     |
| `ExternalURL`                                        | **string*                                            | :heavy_minus_sign:                                   | External URL for the library item.                   |
| `Files`                                              | [][components.File](../../models/components/file.md) | :heavy_minus_sign:                                   | List of files associated with the library item.      |
| `ID`                                                 | **int64*                                             | :heavy_minus_sign:                                   | Unique identifier for the library item.              |
| `PublishedBy`                                        | **string*                                            | :heavy_minus_sign:                                   | Publisher of the library item.                       |
| `PublishedDate`                                      | **string*                                            | :heavy_minus_sign:                                   | Date the library item was published.                 |
| `Title`                                              | **string*                                            | :heavy_minus_sign:                                   | Title of the library item                            |
| `Type`                                               | **string*                                            | :heavy_minus_sign:                                   | Identifies the type of library item, e.g. Image      |