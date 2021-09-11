# Update User

## Internal server error

```JSON
{
    "code": 500,
    "status": "Internal Server Error",
    "body": {
        "error": "An ordinary error :O"
    }
}
```

## Updated

```JSON
{
    "code": 200,
    "status": "OK"
}
```

## Request with no id

```JSON
{
    "code": 400,
    "status": "Bad Request",
    "body": [
        {
            "tag": "id",
            "validation": "required"
        }
    ]
}
```

## Validation error

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
