## Success

```
{
    "code": 201,
    "status": "Created",
    "body": {
        "user_id": 1
    }
}
```

## Unexpected Error

```
{
    "code": 500,
    "status": "Internal Server Error",
    "body": {
        "error": "an error"
    }
}
```

## Validation Error

```
{
    "code": 400,
    "status": "Bad Request",
    "body": {
        "error": [
            {
                "tag": "name",
                "validation": "required"
            }
        ]
    }
}
```