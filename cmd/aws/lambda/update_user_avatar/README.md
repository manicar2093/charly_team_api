# Update user avatar

## Ok Response

```JSON
{
  "code": 200,
  "status": "OK",
  "body": {
    // User updated ...
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
                "tag": "user_uuid",
                "validation": "required"
            },
            {
                "tag": "avatar_url",
                "validation": "required"
            }
        ]
    }
}
```
