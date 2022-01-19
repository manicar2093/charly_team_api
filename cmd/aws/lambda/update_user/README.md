# Update User

## Updated

```JSON
{
  "code": 200,
  "status": "OK",
  "body": {
    "biotype_id": null,
    "bone_density_id": null,
    "gender_id": null,
    "user_uuid": "",
    "avatar_url": "",
    "birthday": "0001-01-01T00:00:00Z",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    ...
  }
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
