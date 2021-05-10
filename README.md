# knative-colors

Colors is a bare bone app that serves a simple blank html page with background color based on an ENV variable. The the branch names tells you the code for each iteration. The versions ( branches ) build on top of this. 

This is the companion app to showcase the basic features of [knative](http://knative.dev). This App in conjunction with the [cologrid](https://github.com/clue2solve/colorgrid) makes it simple to understand the flow. 

<img src="/images/colors.png">

```bash
# Create the color grid service 
kn service create  colorgrid --image ghcr.io/clue2solve/colorgrid/colorgrid:latest

# create the rev-1 of the colors app
kn service create bg-demo --image ghcr.io/clue2solve/knative-colors/colors:start  --env HELLO_BG_COLOR="RED"

# create the rev-2 of the app and split the traffic 
kn service update  bg-demo --image ghcr.io/clue2solve/knative-colors/colors:start  --env HELLO_BG_COLOR="blue"
kn service update bg-demo --traffic bg-demo-00002=50,bg-demo-00001=50

# create the rev-3 of the app and re-split the app 
kn service update  bg-demo --image ghcr.io/clue2solve/knative-colors/colors:start  --env HELLO_BG_COLOR="green"
kn service update bg-demo --traffic bg-demo-00002=40,bg-demo-00001=30,bg-demo-00003=30
```
At each step above,  you should be able to hit the URL for the colorgrid with a parameter `?url={URL_OF_BG_DEMO}`

Example: 
```
http://colorgrid.default.knative.g.demome.cloud/?url=http://bg-demo.default.knative.g.demome.cloud
```

## Rainbow Deploys demo

Getting to eventually a `spec.traffic` like:

```yaml
  traffic:
  - percent: 40
    revisionName: colors-00002
    tag: green
  - percent: 50
    revisionName: colors-00003
    tag: blue
  - percent: 10
    revisionName: colors-00001
    tag: red
```