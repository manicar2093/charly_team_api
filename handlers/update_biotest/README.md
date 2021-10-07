# Update Biotest

If `id` or `uuid` are not in request a `identifier` validation error should be returned

## Updated

```JSON
{
    "code": 200,
    "status": "OK"
}
```

## Without id

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

## Validation Errors

```JSON
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
