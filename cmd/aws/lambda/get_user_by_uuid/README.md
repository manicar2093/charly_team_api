# Get user by UUID

## Request

```JSON
{
    "user_uuid": "83296877ed364c2ca4126e9cb70c1927"
}
```

## OK Response

```JSON
{
  "code": 200,
  "status": "OK",
  "body": {
        "biotype_id": null,
        "bone_density_id": null,
        "gender_id": null,
        "user_uuid": "aaede8005c654592b8e5371cd180b447",
        "avatar_url": "",
        "birthday": "0001-01-01T00:00:00Z",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
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
            }
        ]
    }
}
```
