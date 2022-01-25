# Get All Biotests by User UUID

## Request

```JSON
{
    "user_uuid": "uuid",
    "as_catalog": false
}
```

## Ok Response

```JSON
{
    "code": 200,
    "status": "OK",
    "body": {
        "data": [
            {
                // Biotest
            },
            {
                // Biotest
            }
        ]
    }
}
```

## BadRequest Response

```JSON
{
    "code": 400,
    "status": "Bad Request",
    "body": {
        "error": [
            {
                "tag": "name",
                "validation": "required"
            },
            {
                "tag": "last_name",
                "validation": "required"
            }
        ]
    }
}
```