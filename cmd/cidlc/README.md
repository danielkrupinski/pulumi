# cidlc

CIDLC compiles an Go-based IDL into Coconut metadata and packages.

## IDL

CIDLC accepts a subset of Go.  Please refer to [the IDL design document](/docs/idl.md) for details.

## Providers

The primary use case for CIDLC is to author resource packages and providers.  A resource package is a low level
Coconut package with metadata associated with a set of resource type definitions.  Its associated provider is a dynamic
plugin that implements the behavior associated with those resources, their CRUD functions, and operational semantics.

The CIDLC toolset cuts down on boilerplate and makes it easy to author new resource packages and providers.

## Building

To build the CIDLC tool, run

    $ go install github.com/pulumi/coconut/cmd/cidlc

from an enlistment with a proper `GOPATH` set.

## Running

To generate code, run:

    $ cidlc pkg-name idl-path [flags]

where the `pkg-name` is the name of the target Coconut package, `idl-path` is a path to a directory containing IDL, and
`flags` are an optional set of flags.  All `*.go` files in the target directory are parsed and processed as IDL.  To
recursively fetch sub-packages, pass the `--recursive` (or `-r`) flag.

The output includes the following:

* A Coconut package in CocoJS, containing resource definitions.
* An RPC package in Go to aid in the creation of a resource provider, containing:
    - A base resource provider that handles marshaling goo at the edges.
    - A marshalable type for each resource type (used for dynamic plugin serialization).

The two flags, `--out-pack` and `--out-rpc` control if and where the output will be generated, respectively.  It is
possible to specify just one or the other, or you can generate both simultaneously.

By default, the IDL Go package is inferred from a combination of `idl-path` and `GOPATH`.  This is used to generate
inter- and intra-package references.  If you are generating RPC code, that too is inferred, based on `--out-rpc` and
`GOPATH`.  In the event you are running outside of a Go workspace (where `GOPATH` is not set), or need to customize
these, the packages can be set by hand using the flags `--pkg-base-idl` and `--pkg-base-rpc`, respectively.

After generating the code and implementing behavior associated with the resource in the provider, the Coconut package
may then be distributed to consumers using CocoPy, CocoRu, CocoJS, and other CocoLangs.
