# Get Catalogs

## Request

```JSON
{
    "catalog_names": [
        "biotest",
        "roles"
    ]
}
```

## OK. Get Catalogs

```JSON
{
    "code": 200,
    "status": "OK",
    "body": {
        "biotype": [
            ...
        ],
        "roles": [
            ...
        ]
    }
}
```

## NotFound Catalog does not exists

``` JSON
{
    "code": 404,
    "status": "Not Found",
    "body": {
        "error": "catalog 'no_exists' does not exists"
    }
}
```

## BadRequest Validation errors

``` JSON
{
    "code": 400,
    "status": "Bad Request",
    "body": {
        "errors": [
            {
                "tag": "biotype",
                "validation": "required"
            }
        ],
    }
}
```

## Internal Server error

```JSON
{
    "code": 500,
    "status": "Internal Server Error",
    "body": {
        "error": "An ordinary error :O"
    }
}
```
