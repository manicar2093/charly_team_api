# Create User AWS Lambda

## Request

```JSON
{
    "name": "testing",
    "last_name": "main",
    "email": "testing@main-func.com",
    "birthday": "2021-09-27T19:38:57.731173821-05:00",
    "role_id": 3,
    "gender_id": 1
}
```

## Success

```JSON
{
    "code": 201,
    "status": "Created",
    "body": {
        "id": 1,
        "biotype_id": null,
        "bone_density_id": null,
        "gender_id": null,
        "user_uuid": "",
        "avatar_url": "",
        "birthday": "0001-01-01T00:00:00Z",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
    }
 }
```

## Unexpected Error

```JSON
{
    "code": 500,
    "status": "Internal Server Error",
    "body": {
        "error": "an error"
    }
}
```

## Validation Error

```JSON
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
