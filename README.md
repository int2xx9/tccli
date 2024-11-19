tccli
====================

An extremely limited TeamCity client.

## Usage

### Specify a server and authenticate
You need to pass `--server` and `--token` to use tccli.

Use `http://example.com` for `--server` if `http://example.com/app/rest` is the TeamCity API endpoint.

Provide your access token for `--token`.

```sh
# example
tccli --server http://example.com --token ...
```

### Example: Show builds for a project
```sh
# example
tccli --server http://example.com --token ... builds getAll --build-locator 'project:TeamCityProject'
```

### Example: Show builds as csv
```sh
# example
tccli --server http://example.com --token ... builds getAll --build-locator 'project:TeamCityProject' --output csv --csv-format '{{.Number}},{{.Status}}|status,{{.Href}}|href'
```

Result:
```csv
{{.Number}},status,href
7,FAILURE,/app/rest/builds/id:110
6,FAILURE,/app/rest/builds/id:109
5,FAILURE,/app/rest/builds/id:5
4,SUCCESS,/app/rest/builds/id:4
3,SUCCESS,/app/rest/builds/id:3
2,SUCCESS,/app/rest/builds/id:2
1,SUCCESS,/app/rest/builds/id:1
```

> [!NOTE]
> See teamcity/schema/*.go to know a name of field to specify in `{{...}}`.

### Example: Download artifacts as an archive
```sh
# example
tccli --server http://example.com --token ... builds artifacts getArchive --build-locator 'id:1' > artifacts.zip
```

## Supported APIs

Basically tccli commands correspond to TeamCity REST API.

Below is a list of supported APIs. For instance, `tccli buildTypes get` corresponds to `/app/rest/buildTypes/{btLocator}`.

- buildTypes
  - get: `/app/rest/buildTypes/{btLocator}`
  - getAll: `/app/rest/buildTypes`
- builds
  - artifacts
    - getArchive: `/app/rest/builds/{buildLocator}/artifacts/archived`
  - get: `/app/rest/builds/{buildLocator}`
  - getAll: `/app/rest/builds`
  - testOccurrences: `/app/rest/{buildLocator}/testOccurrences`
