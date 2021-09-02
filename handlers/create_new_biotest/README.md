## Success

```
{
    "code": 201,
    "status": "Created",
    "body": {
        "biotest_id": 1
    }
}
```

## Internal Server Error
```
{
    "code": 500,
    "status": "Internal Server Error",
    "body": {
        "error": "An error occured :O"
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
                "tag": "weight",
                "validation": "required"
            },
            {
                "tag": "height",
                "validation": "required"
            }
        ]
    }
}
```