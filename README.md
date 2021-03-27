# Self-hosted git helper

## For what?
Converting your

***git@domain:user/repo.git***

to 

***ssh://git@domain:port/user/repo.git***

## Using
For converting its needed just change ***git*** to ***self-gitea*** (your built app):
```
$ self-gitea clone git@github.com:EgorKharlamov/self-gitea.git
```
You can set port with flag:
```
$ self-gitea -port=22 clone git@github.com:EgorKharlamov/self-gitea.git
```
