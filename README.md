# Fitman

## firebase token manager

fitman is the token manager for firebase-auth.
It provides appropriate token management and refresh to prevent unnecessary user creation and unify token management methods during intra-team development.

Please, Put it in .gitignore when you use it by typing the command in the root of the project

# how to install
```
$ go install github.com/gari8/fitman@latest
```

# how to use
## commands
### fitman i, init [profile]
-> initialize fitman with setting profile

### fitman a, add [profile]
-> add profile to fitman's db

### fitman d, delete [profile]
-> delete profile from fitman's db

### fitman ls, list
-> show registered profiles

### fitman g, get [profile]
-> get profile's idToken

### fitman help
### fitman --version