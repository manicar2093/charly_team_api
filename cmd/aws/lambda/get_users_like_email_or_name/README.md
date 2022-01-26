# Get Users Like Email or Name

## Request

```JSON
{
    "filter_data": "name"
}
```

## Response OK

```JSON
{
    "code": 200,
    "status": "OK",
    "body": [
        {
            // Biotest
        },
        {
            // Biotest
        }
    ]
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
                "tag": "filter_data",
                "validation": "required"
            }
        ]
    }
}
```
