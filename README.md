# Crossplane Stack Development: Walk through with Linode

### ⚠️ Deprecation notice ⚠️
Please use Linode's actively maintained provider here: https://github.com/linode/provider-linode

What is a [Crossplane Stack](https://crossplane.io/docs/v0.3/concepts.html#stacks)?  It's a way to add functionality to Crossplane.
What is Crossplane? If you are asking that, you should probably read [Welcome to Crossplane!](https://crossplane.io/docs/v0.3/) first.

Stacks can be used to teach Crossplane how to power another infrastructure
provider, or Stacks can configure the pieces of a Kubernetes application that
depends on managed cloud resources playing with your typical Kubernetes resources.

Authoring a Crossplane Stack currently involves writing a Kubernetes 
Controller.  The [Crossplane Slack](https://slack.crossplane.io/) [#sig-stacks](https://crossplane.slack.com/messages/CKXQK4P27) community will be looking for ways to simplify this.

## From Kubebuilder 2 Controllers to Crossplane Runtime Stacks

Kubebuilder offers an easy way to get a Kubernetes controller project
bootstrapped.  Like any framework, Kubebuilder has a set of features, rules,
and opinions.

At the core of [Crossplane's own managed resources](https://github.com/crossplane/crossplane/tree/master/apis) and [infrastructure Stacks](https://github.com/crossplaneio?utf8=%E2%9C%93&q=stack-&type=&language=),
you'll find the [Crossplane Runtime](https://github.com/crossplane/crossplane-runtime).  This underpinning of Crossplane offers
an improved experience for state reconciliation.  It provides a number of hooks that
come in handy when creating a Stack or any managed resource controller working
in the narrow field of cloud resource provisioning.

## Pick a Stack any Stack

As a current Crossplane Stacks developer and a former Linode developer,
I'm venturing into Crossplane Services land by taking along
the things that I know well enough to hack on without getting overtly stuck.

Crossplane's most prominent demos to date have involved some combination of
wrangling a herd of managed Kubernetes, managed database, managed caching,
and managed object storage solutions. Recent developments allow network
and user management to be roped in.

These services are common stock among GCP, AWS, and Azure.

Providers like Packet and Linode don't necessarily offer all of these
services, which could make them seem like an awkward choice.

Crossplane's portable Resource Classes and Resource Claims can't be
easily taken advantage of.

On the other hand, the simpler array of services is easier to reason about.
The lack of appropriate Resource Classes give us cause to create new ones.

## Enough of the why.. Let's get to the how

Now that our reasons have been stated, let's start our Stack project by
creating a new workspace and taking full advantage of Kubebuilder's ease of use:

```sh
mkdir stack-linode
cd stack-linode
git init

kubebuilder init --domain stack.crossplane.io \
  --owner "The Crossplane Authors" \
  --repo github.com/displague/stack-linode
```

The piece that is not quite obvious here is the `domain`.  This will serve as the
base for the FQDN used in our API groups.

Now let's create an API group for Linode and introduce some `kind` types.

```sh
kubebuilder create api --group="linode" --version v1alpha1 \
  --kind Provider --resource=true --controller=false
```

The [Provider kinds](https://crossplane.io/docs/v0.3/services-developer-guide.html#provider-kinds) will
act as our configuration of Linode API settings.  You can have multiple `Provider` records, each
using a different API Token or a different account.  Some Provider records may allow you to change
the Project or Region.  This is top level configuration for the resource provider.

As a bag of configuration, `Provider` kinds don't need a controller.  They are not active. Crossplane Managed controllers are the actor.  They will find references to the `Provider` and they will read its settings in order to create managed resources.

In Crossplane, the typical API group and kind for Provider kinds is `provider.foo.crossplane.io`, where `foo` is the name of the Stack. In this project we have gone against convention and will have `provider.foo.stack.crossplane.io`.  

Similarly, Crossplane [Managed Resource kinds](https://crossplane.io/docs/v0.3/services-developer-guide.html#managed-resource-kinds) are typically `bar.compute.foo.crossplane.io`, where `bar` is the name of the external resource (`instance`, `mysql`, `gkecluster`).  Crossplane's convention for resource classes uses a common group prefix for related resource classes (`compute`, `database`, `cache`, `storage`).

I'm pointing these conventions out because we are breaking them all by using the API groups that Kubebuilder generates, and that's just fine.  These are just naming conventions and nothing in Crossplane cares about them.  We'll just need to make sure that we don't trip over the differing API groups ourselves.

This is all explained in the "[What Makes a Crossplane Managed Service?](https://crossplane.io/docs/v0.3/services-developer-guide.html#what-makes-a-crossplane-managed-service)" section of the [Services Developer Guide](https://crossplane.io/docs/v0.3/services-developer-guide.html).

```sh
kubebuilder create api --group="linode" --version v1alpha1 \
  --kind Instance --resource=true --controller=false
```

We are leaving out the Kubebuilder generated controller because Crossplane Runtime
takes a different approach with [managed resource controllers](https://crossplane.io/docs/v0.3/services-developer-guide.html#managed-resource-controllers).

Linode Instances will need a controller. We will follow the steps outlined in the
 managed resource controllers section of the Service Developers Guide.

