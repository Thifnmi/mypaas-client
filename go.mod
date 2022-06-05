module github.com/Thifnmi/mypaas-client

go 1.12

require (
	github.com/ajg/form v0.0.0-20160822230020-523a5da1a92f
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6
	github.com/digitalocean/godo v1.1.1 // indirect
	github.com/docker/docker v20.10.8+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/machine v0.16.1
	github.com/exoscale/egoscale v0.9.31 // indirect
	github.com/fsouza/go-dockerclient v1.7.4
	github.com/ghodss/yaml v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/pmorie/go-open-service-broker-client v0.0.0-20180330214919-dca737037ce6
	github.com/sabhiram/go-gitignore v0.0.0-20171017070213-362f9845770f
	github.com/sethvargo/go-password v0.1.1
	github.com/Thifnmi/config v1.0.0
	github.com/Thifnmi/gnuflag v1.0.0
	github.com/Thifnmi/go-tsuruclient v1.0.0
	github.com/Thifnmi/tablecli v1.0.0
	github.com/Thifnmi/mypaas v1.0.0
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.20.6
)

replace (
	github.com/ajg/form => github.com/cezarsa/form v0.0.0-20210510165411-863b166467b9
	github.com/samalba/dockerclient => github.com/cezarsa/dockerclient v0.0.0-20190924055524-af5052a88081
)
