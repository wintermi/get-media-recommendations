# Get Media API Recommendations

[![Workflows](https://github.com/wintermi/get-media-recommendations/workflows/go.yaml/badge.svg)](https://github.com/wintermi/get-media-recommendations/actions)
[![Go Report](https://goreportcard.com/badge/github.com/wintermi/get-media-recommendations)](https://goreportcard.com/report/github.com/wintermi/get-media-recommendations)
[![License](https://img.shields.io/github/license/wintermi/get-media-recommendations.svg)](https://github.com/wintermi/get-media-recommendations/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/wintermi/get-media-recommendations?include_prereleases)](https://github.com/wintermi/get-media-recommendations/releases)


## Description

A command line application designed to provide a simple method to request recommendations from a Google Cloud Discovery Engine API model for the parameters contained within the input file.

```
USAGE:
    get-media-recommendations -p PROJECT_NUMBER -s SERVING_CONFIG -i INPUT_FILE

ARGS:
  -b string
    	Branch (default "0")
  -d string
    	Data Store (default "default_data_store")
  -f string
    	Filter String
  -i string
    	Parameter Input File  (Required)
  -l string
    	Location (default "global")
  -n int
    	Number of Results, 1 to 100 (default 5)
  -p string
    	Project Number  (Required)
  -s string
    	Serving Config  (Required)
  -v	Output Verbose Detail
```

## Example Parameter Input File

```
[
  {
    "eventType": "view-item",
    "userPseudoId": "1",
    "documents": [
      {
        "id": "0598"
      }
    ]
  },
  {
    "eventType": "view-item",
    "userPseudoId": "2",
    "documents": [
      {
        "id": "0986"
      }
    ]
  }
]
```


## License

**get-media-recommendations** is released under the [Apache License 2.0](https://github.com/wintermi/get-media-recommendations/blob/main/LICENSE) unless explicitly mentioned in the file header.
