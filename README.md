# Self-hosted git helper

## For what?
Converting your

***git@domain:user/repo.git***

to 

***ssh://git@domain:port/user/repo.git***

## Using
1. First of all you need to set project name in ./configs/.project_name (e.x. ***self-gitea***)
2. Then you can build app by command:
```
$ make bld
```
3. Or you can clean build directory by command:
```
$ make clean
```

For converting its needed just change ***git*** to ***self-gitea*** (your built app):
```
$ self-gitea clone git@github.com:EgorKharlamov/self-gitea.git
```
You can set port with flag:
```
$ self-gitea -port=22 clone git@github.com:EgorKharlamov/self-gitea.git
```
