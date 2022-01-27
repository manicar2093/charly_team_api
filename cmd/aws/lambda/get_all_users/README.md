# Get All Users

## Request

This use pagination request, but it is a must request contains page_number

```JSON
{
    "page_number": 1
}
```

## OK Response

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
                "tag": "page_number",
                "validation": "required"
            },
        ]
    }
}
```
