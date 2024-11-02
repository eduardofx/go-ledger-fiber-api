# Transactions

### [POST][/transactions] - Create an Transaction

#### Request Body

    {
        "amount": 1.50,
        "idempotency_key": "1ds2das",
        "account_id": "12686d6b-6df2-45b6-a450-c9517018ce52",
        "operation_id": "d6ae2972-ff10-48f3-a3df-d50e00880230"
    }

#### Return 201 - Transaction created

    {
        "Id": "8b991d8e-b6a3-4689-8629-5a954bd80e1e",
        "amount": 1.50,
        "account_id": "97fe1c49-4b19-4f75-863c-e687fb22dd3e",
        "operation_id": "b3208db2-7a0c-4244-8011-1655d4754820",
        "IdempotencyKey": "1ds2das",
        "CreatedAt": "2024-11-02T12:09:56.273306833Z"
    }

#### Return 409 - IdempotenceKey exists

    {
        "Id": "8b991d8e-b6a3-4689-8629-5a954bd80e1e",
        "amount": 1.50,
        "account_id": "97fe1c49-4b19-4f75-863c-e687fb22dd3e",
        "operation_id": "b3208db2-7a0c-4244-8011-1655d4754820",
        "IdempotencyKey": "1ds2das",
        "CreatedAt": "2024-11-02T12:09:56.273306Z"
    }

#### Return 400 - Insufficient balances

    {
        "message": "Insufficient balances"
    }

#### Return 400 - Amount format verification

    {
        "message": "Amount must have at most two decimal places"
    }

#### Return 404 - Account not found

    {
        "message": "Account not found"
    }

#### Return 404 - Operation not found

    {
         "message": "Operation not found"
    }
