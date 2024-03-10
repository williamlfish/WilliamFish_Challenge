### WilliamFish_Challenge 

#### Info
I went into this challenge trying to image that the application would be scaled to a large degree, and that the html that 
is served would eventually be used in the credit card validation so just kept them together, but with the validation as a go pkg
so it could potentially be used for any go project.
decided to use EKS for the application and can be found at https://ccc.nboop.com
and you can test the validator with
```sh
  curl -XPOST -d "{\"ccNumber\":\"4233333333445552\"}" https://ccc.nboop.com
```
the `bin` dir has some scripts for deploying a new version of the application, using the githash as the image tag
so ci/cd could be added easily, and another for bootstrapping the EKS cluster with cert-manager and an ingress controller.

`infra` has kube manifests

`pkg` has the credit card validation logic 

`src` is the server 

`templates` is the static html 

`go test ./..` will run all of the unit tests. 