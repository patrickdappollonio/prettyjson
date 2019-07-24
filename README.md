# `prettyjson` [![Build Status](https://travis-ci.org/patrickdappollonio/prettyjson.svg?branch=master)](https://travis-ci.org/patrickdappollonio/prettyjson)

A JSON CLI Tool to prettify JSON. It will take from `stdin` any source of JSON, for example, an API, and it will
print to `stdout` a prettified, human-readable JSON.

### Example

Given an API that returns compressed JSON, and making a `cURL` request like this:

```bash
$ curl https://example.com
{"firstName":"Bidhan","lastName":"Chatterjee","age":40,"email":"bidhan@example.com"}
```

We can improve the printing by executing `prettyjson` at the end, piping the response from `curl`, and omitting the `cURL` download information output:

```bash
$ curl -s https://example.com | prettyjson
{
  "firstName": "Bidhan",
  "lastName": "Chatterjee",
  "age": 40,
  "email": "bidhan@example.com"
}
```
