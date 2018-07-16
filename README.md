
![Netograph API](assets/netograph-logo.png)

# Netograph API

This repository contains a command-line tool and language packs for interacting
with the Netograph API.

# ngc: Netograph command-line interface

**ngc** is a command-line tool that exposes the complete Netograph interace for
interactive use. If you have a Go development environment set up, you can
install it as follows:

```sh
go get -u github.com/netograph/netograph-api/go/cmd/ngc
```

After installation, try running `ngc` for a high-level overview of the API, and
`ngc help command` for help on any specific command.

## Configuration

**ngc** is configured through environment variables.

| Variable | Description |
| --- | --- |
| NGC_DATASET | The dataset to operate on, unless over-ridden with the `--dset` flag. Defaults to `netograph:social`, our largest public data repository
| NGC_TOKEN  | The authentication token for API access. Can be over-ridden with the `--token` command-line flag.  |


## Output formats

By defult **ngc** outputs data in an easy-to-read text format. There are two
other formats available - intendted JSON (the `--json` flag) and compact JSON
(the `--cjson` flag). Compact JSON consists of one record per line with no
newlines, and so is suitable for programmatic use.


## Examples

Begin by exporting your access token:

```sh
export NGC_TOKEN=MYTOKEN
```

You can now list the datasets that you have access to:

```sh
ngc datasets
```

Note that public datasets will be marked read-only. You can query these, but not
submit new URLs to be captured to them. For datasets you have write access to,
see the `ngc submitcapture` command for capture submission.

The default dataset is the `netograph:social` dataset, which aggregates a
sizeable fraction of all URLs passing through social media in real time. As a
first step, let's list al the satellites we've ever seen for a domain query.

```sh
ngc satellitesforroot rt.com
```

There are a few things to note here.

- This command searches for rt.com and ALL its subdomains. See the [Domain
  Queries](#domain-queries) section for a description of how to restrict to a
  specific domain.
- The output is pretty verbose, and it would be nice to be able to extract only
  the information we're interested in. We recommend
  [jq](https://stedolan.github.io/jq/), a light-weight command-line JSON
  processor for this.
- The command name is pretty long. Most commands have shorter aliases you can
  view using `ng help`.
- By default we limit the number of responses for queries to 100. It's pretty
  easy to craft a command that will return millions of records. In this case, we
  want to list ALL of the satellites we've ever seen on `rt.com`, and we can
  disable this limit by setting it to `0`.

Putting all of this together, we have a command like this:

```sh
ngc -n 0 --cjson satsforroot  | jq -r .satelite
```

At the time of writing, this command cleanly lists about 3600 third-party
domains for `rt.com`.






# Language packs


## Go


## Python


# API Notes

## Domain queries

## Resumption tokens
