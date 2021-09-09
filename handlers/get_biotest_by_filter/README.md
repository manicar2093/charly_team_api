# BiotestFilters

## Registered filters

- _get_biotest_by_uuid_

## OK

```JSON
{
    "code": 200,
    "status": "OK",
    "body": {
        "id": 1,
        "higher_muscle_density": {
            "neck": null,
            "shoulders": null,
            "back": null,
            "chest": null,
            "back_chest": null,
            "right_relaxed_bicep": null,
            "right_contracted_bicep": null,
            "left_relaxed_bicep": null,
            "left_contracted_bicep": null,
            "right_forearm": null,
            "left_forearm": null,
            "wrists": null,
            "high_abdomen": null,
            "lower_abdomen": null
        },
        "lower_muscle_density": {
            "hips": null,
            "right_leg": null,
            "left_leg": null,
            "right_calf": null,
            "left_calf": null
        },
        "skin_folds": {
            "subscapular": null,
            "suprailiac": null,
            "bicipital": null,
            "tricipital": null
        },
        "glucose": null,
        "resting_heart_rate": null,
        "maximum_heart_rate": null,
        "observations": null,
        "recommendations": null,
        "front_picture": null,
        "back_picture": null,
        "right_side_picture": null,
        "left_side_picture": null,
        "next_evaluation": null,
        "created_at": "0001-01-01T00:00:00Z"
    }
}
```

## Not Found

```JSON
{
    "code": 404,
    "status": "Not Found",
    "body": {
        "error": "Record not found"
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

## Validation error

```JSON
{
    "code": 400,
    "status": "Bad Request",
    "body": {
        "error": [
            {
                "tag": "biotest_id",
                "validation": "required"
            }
        ]
    }
}
```
