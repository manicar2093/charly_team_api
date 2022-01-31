# Entities

Due `gopkg.in/guregu/null.v4` package autosave on `github.com/go-rel/rel` package does not work properly. This is why it is need to save first parent entity and then children data.

This can be a problem if some slice have to be saved, but this is a future worry.

This errror must be consider on update too.
