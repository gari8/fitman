# Fitman

## firebase token manager

fitman is the token manager for firebase-auth.
It provides appropriate token management and refresh to prevent unnecessary user creation and unify token management methods during intra-team development.

Please, Put it in .gitignore when you use it by typing the command in the root of the project

# how to use
```
[sub commands]
// create .fitman directory & get idToken
fitman init


// show idToken (after init) 
fitman get

// show help
fitman help

// show version
fitman version

[option]
fitman -v get
{
  "access_token": "dummy",
  "expires_in": "3600",
  "token_type": "Bearer",
  "refresh_token": "dummy",
  "id_token": "dummy",
  "user_id": "dummy",
  "project_id": "dummy"
}

```

