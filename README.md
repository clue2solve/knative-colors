# knative-colors

Colors is a bare bone app that serves a simple blank html page with background color based on an ENV variable. The the branch names tells you the code for each iteration. The versions ( branches ) build on top of this. 

This is the companion app to showcase the basic features of [knative](http://knative.dev). This App in conjunction with the [cologrid](https://github.com/clue2solve/colorgrid) makes it simple to understand the flow. 

<img src="/images/colors.png">

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