# Operations

### [POST][/operations] - Create an operation

#### Request Body  (Type: [asset,liability])

    {
     "name": "withdrawal",
     "type": "asset" 
    }

#### Return 201

    {
        "Id": "b3208db2-7a0c-4244-8011-1655d4754820",
        "Name": "withdrawal",
        "Type": "asset",
        "CreatedAt": "2024-11-02T11:57:06.861965339Z",
        "DeletedAt": null
    }

#### Return 400 - Invalid Type

    {
        "message": "This operation does not exist, choose 'asset' or 'liability"
    }

### [GET][/operations] - Get All operations

#### Return 200

    [
        {
            "Id": "b3208db2-7a0c-4244-8011-1655d4754820",
            "Name": "withdrawal",
            "Type": "asset",
            "CreatedAt": "2024-11-02T11:57:06.861965Z",
            "DeletedAt": null
        },
        {
            "Id": "cbe56fb2-f019-4c77-b064-4ce9998e65a8",
            "Name": "prurchase",
            "Type": "liability",
            "CreatedAt": "2024-11-02T12:02:45.316693Z",
            "DeletedAt": null
        }
    ]

### [DELETE][/operations/:id] - Delete an operation

#### Return 200

    {
        "id": "338ae3ab-ccca-48ce-b420-ea8c70a3f013",
        "message": "The operation has been successfully deleted."
    }
