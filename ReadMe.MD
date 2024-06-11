 <div style="text-align: center">
  <h1>Forge</h1>
 </div>

 <div style="text-align: center">
<img src="forge.png" height="250" width="250">
 </div>

This project is inspired by the likes of t3 create-react-app, ng generate, etc but with a focus on go.

There are a few different options for Forge.
- The first is a generic blank project complete with the words of the ancestors of 'Hello, world'.
- Second is a blank rest api service that generates main.go with net/http, this might be extended in the future for other routing options such as gorilla/mux, echo, or chi, but I have found that the latest version of net/http is good enough
- Third is the bases of a cli project with cobra installed as well.
- There are some helper utilities as well, to create a models folder and file, a service folder and file that can be named with `forge service <YourService>`
<br><br>

# That's cool and all but why though?
I honestly just wanted to explore cobra more thoroughly, and this can be useful for me in the future as I create random Go projects. I know it's not hard to create new projects in go, but maybe this can evolve over time as become more familiar with cobra. I've found the best way to learn something is to put it in action.
<br><br>

# Available Commands
### blank: 
Generic empty go project.
```shell 
forge blank [name]
# blank aliases: b, --b, new 
# name is optional, if not provided user is prompted for one
```
### cli:
Cobra focused CLI project
``` shell
forge cli [name]
# cli aliases: c --c
# name is optional, if not provided user in prompted for one
```
### git: 
just another way to init a git repo
```shell
forge blank [name] -g
```
### help: 
The inbuilt help command
``` shell
forge help
```
### models: 
Create a models folder and file
```shell
forge models
# alias m
```
### restful API: 
Create a rest API with net/http (assuming go 1.22) with middleware, helpers, models, etc.
```shell
forge rest [name]
# alias r
# name is optional, if not provided user is prompted for one
```

### service: 
Create a new service structure for apis
```shell
forge service [name]
# alias s
# name is required
```
<br><br>

# Future stuff
Ideally I would love to add gRPC as a template option, but that seems a bit trickier to do at this time and will play around with it.

<br><br>

There is also an included Taskfile with the ability to build for different OSes.
- build-mac
- bin-mac: adds to /usr/local/bin
- build-linux
- bin-linux: adds to /usr/local/bin
- build-windows
- bin-windows: adds to C:\Windows\System32\
- zsh-ac: creates zsh auto complete instructions
- pwsh-ac: creates pwsh auto complete instructions

# Lastly
If you have any feedback or ways to improve this this I am happy to connect and discuss.