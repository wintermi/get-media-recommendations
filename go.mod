module github.com/wintermi/get-media-recommendations

go 1.19

replace (
	cloud.google.com/go/discoveryengine/apiv1alpha => ./cloud.google.com/go/discoveryengine/apiv1alpha
	google.golang.org/genproto/googleapis/cloud/discoveryengine/v1alpha => ./google.golang.org/genproto/googleapis/cloud/discoveryengine/v1alpha
)

require (
	cloud.google.com/go/discoveryengine/apiv1alpha v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.29.0
	google.golang.org/genproto/googleapis/cloud/discoveryengine/v1alpha v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.1
)

require (
	cloud.google.com/go v0.109.0 // indirect
	cloud.google.com/go/compute v1.18.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/longrunning v0.4.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.1 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/oauth2 v0.4.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/api v0.109.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57 // indirect
	google.golang.org/grpc v1.52.3 // indirect
)
