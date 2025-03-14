![Mux Go Banner](https://banner.mux.dev/?image=go)

![](https://github.com/muxinc/mux-go/workflows/Integration%20Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/muxinc/mux-go?status.svg)](https://godoc.org/github.com/muxinc/mux-go)

# Mux Go

Official Mux API wrapper for golang projects, supporting both Mux Data and Mux Video.

[Mux Video](https://mux.com/video) is an API-first platform, powered by data and designed by video experts to make beautiful video possible for every development team.

[Mux Data](https://mux.com/data) is a platform for monitoring your video streaming performance with just a few lines of code. Get in-depth quality of service analytics on web, mobile, and OTT devices.

Not familiar with Mux? Check out https://mux.com/ for more information.

## Installation

```
go get github.com/muxinc/mux-go/v6
```

## Getting Started

### Overview

Mux Go is a code generated lightweight wrapper around the Mux REST API and reflects them accurately. This has a few consequences you should watch out for:

1) For almost all API responses, the object you're looking for will be in the `data` field on the API response object, as in the example below. This is because we designed our APIs with similar concepts to the [JSON:API](https://jsonapi.org/) standard. This means we'll be able to return more metadata from our API calls (such as related entities) without the need to make breaking changes to our APIs. We've decided not to hide that in this library.

2) We don't use a lot of object orientation. For example API calls that happen on a single asset don't exist in the asset class, but are API calls in the AssetsApi which require an asset ID.

### Authentication

To use the Mux API, you'll need an access token and a secret. [Details on obtaining these can be found here in the Mux documentation.](https://docs.mux.com/docs#section-1-get-an-api-access-token)

Its up to you to manage your token and secret. In our examples, we read them from `MUX_TOKEN_ID` and `MUX_TOKEN_SECRET` in your environment.

### Example Usage

Below is a quick example of using mux-go to list the Video assets stored in your Mux account.

Be sure to also checkout the [exmples directory](examples/).

```go
package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go/v6"
)

func main() {
	// API Client Init
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// List Assets
	fmt.Println("Listing Assets...\n")
	r, err := client.AssetsApi.ListAssets()
	if err != nil {
		fmt.Printf("err: %s \n\n", err)
		os.Exit(255)
	}
	for _, asset := range r.Data {
		fmt.Printf("Asset ID: %s\n", asset.Id)
		fmt.Printf("Status: %s\n", asset.Status)
		fmt.Printf("Duration: %f\n\n", asset.Duration)
	}
}
```

## Errors & Error Handling

All API calls return an err as their final return value. Below is documented the errors you might want to check for. You can check `error.Body` on all errors to see the full HTTP response.

### BadRequestError

`BadRequestError` is returned when a you make a bad request to Mux, this likely means you've passed in an invalid parameter to the API call. 

### UnauthorizedError

`UnauthorizedError` is returned when Mux cannot authenticate your request. [You should check you have configured your credentials correctly.](#authentication)

### ForbiddenError

`ForbiddenError` is returned you don't have permission to access that resource. [You should check you have configured your credentials correctly.](#authentication)

### NotFoundError

`NotFoundError` is returned when a resource is not found. This is useful when trying to get an entity by its ID.

### TooManyRequestsError

`TooManyRequestsError` is returned when you exceed the maximum number of requests allowed for a given time period. Please get in touch with [support@mux.com](mailto:support@mux.com) if you need to talk about this limit.

### ServiceError

`ServiceError` is returned when Mux returns a HTTP 5XX Status Code. If you encounter this reproducibly, please get in touch with [support@mux.com](mailto:support@mux.com).

### GenericOpenAPIError

`GenericOpenAPIError` is a fallback Error which may be returned in some edge cases. This will be deprecated in a later release but remains present for API compatibility.

## Documentation

[Be sure to check out the documentation in the `docs` directory.](docs/)

## Issues

If you run into problems, [please raise a GitHub issue,](https://github.com/muxinc/mux-go/issues) filling in the issue template. We'll take a look as soon as possible.

## Contributing

Please do not submit PRs against this package. It is generated from our OpenAPI definitions - [Please open an issue instead!](https://github.com/muxinc/mux-go/issues)

## License

[MIT License.](LICENSE) Copyright 2019 Mux, Inc.
