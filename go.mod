module github.com/displague/crossplane-provider-linode

go 1.13

replace github.com/linode/linodego => /home/marques/.local/share/go/src/github.com/linode/linodego

replace github.com/crossplane/crossplane-runtime => github.com/negz/crossplane-runtime v0.0.0-20200417025116-3face651efbf

require (
	github.com/crossplane/crossplane v0.10.0-rc.0.20200410142608-84b1c08d1890
	github.com/crossplane/crossplane-runtime v0.7.1-0.20200424213213-10ecf0f09a8a
	github.com/crossplane/crossplane-tools v0.0.0-20200303232609-b3831cbb446d
	github.com/crossplane/provider-gcp v0.8.0
	github.com/crossplanebook/provider-template v0.0.0-20200331182914-b7c332a53f1e
	github.com/linode/linodego v0.14.0
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/pkg/errors v0.8.1
	golang.org/x/net v0.0.0-20191004110552-13f9640d40b9
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.18.0
	k8s.io/apimachinery v0.18.0
	k8s.io/client-go v0.18.0
	sigs.k8s.io/controller-runtime v0.5.1-0.20200422200944-a457e2791293
	sigs.k8s.io/controller-tools v0.2.4
)
