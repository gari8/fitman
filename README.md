# Fitman

## firebase token manager

fitman is the token manager for firebase-auth.
It provides appropriate token management and refresh to prevent unnecessary user creation and unify token management methods during intra-team development.

Please, Put it in .gitignore when you use it by typing the command in the root of the project.

# how to install
```
$ go install github.com/gari8/fitman@latest
```

# how to use
## commands
### fitman i, init [profile(optional)]
initialize fitman with setting profile
#### options
- o, only-token
  - Result contains only idToken
```bash
fitman init -o
```

### fitman a, add [profile(optional)]
add profile to fitman's db
#### options
- o, only-token
    - Result contains only idToken
```bash
fitman add user1 -o 
```

### fitman d, delete [profile(optional)]
delete profile from fitman's db
```bash
fitman delete user1
```

### fitman ls, list
show registered profiles
```bash
fitman ls
```

### fitman g, get [profile(optional)]
get profile's idToken
#### options
- o, only-token
    - Result contains only idToken
```bash
fitman get user1 -o
```

### fitman help
```bash
fitman help
```
## version
```bash
fitman --version
fitman -v
```

# how to develop
```bash
go install go.uber.org/mock/mockgen
```