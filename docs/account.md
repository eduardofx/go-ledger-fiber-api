# Accounts

### [POST][/accounts] - Create Account

#### Request Body

    {
        "document": "26273128049"
    }

#### Return 201

    {
        "Id": "e44fc7c3-40c9-464c-b483-0295e0b66fdb",
        "document": "26273128049",
        "balance": 0,
        "CreatedAt": "2024-11-02T11:38:50.170086345Z",
        "DeletedAt": null
    }

#### Return 400 - Invalid document

    {
        "message": "Invalid document: 111"
    }

### [DELETE][/accounts/:id] - Delete Account

#### Return 200

    {
        "id": "e44fc7c3-40c9-464c-b483-0295e0b66fdb",
        "message": "The account has been successfully deleted."
    }

#### Return 400 - Invalid document

    {
        record not found
    }

### [GET][/accounts] - Return all Accounts

#### Return 200

    [
        {
            "Id": "e44fc7c3-40c9-464c-b483-0295e0b66fdb",
            "document": "26273128049",
            "balance": 0,
            "CreatedAt": "2024-11-02T11:38:50.170086Z",
            "DeletedAt": null
        },
        {
            "Id": "97fe1c49-4b19-4f75-863c-e687fb22dd3e",
            "document": "20754152006",
            "balance": 0,
            "CreatedAt": "2024-11-02T11:46:09.111412Z",
            "DeletedAt": null
        }
    ]

### [GET][/accounts/:document] - Return one account by CPF or CPNJ

#### Return 200

    {
        "Id": "e44fc7c3-40c9-464c-b483-0295e0b66fdb",
        "document": "26273128049",
        "balance": 0,
        "CreatedAt": "2024-11-02T11:38:50.170086Z",
        "DeletedAt": null
    }

#### Return 404

    {
        record not found
    }
