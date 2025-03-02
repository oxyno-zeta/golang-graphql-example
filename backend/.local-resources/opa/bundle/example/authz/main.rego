package example.authz

default allowed = false

allowed if input.user.preferred_username == "user"
