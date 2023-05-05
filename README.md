# GCP Resource Validator

This is a sample for GCP resource validation service evaluating asset from [Cloud Asset Inventory Feed](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes). 
The service use [Config Validator](https://github.com/GoogleCloudPlatform/config-validator) and [Policy Library](https://github.com/GoogleCloudPlatform/policy-library). 

## Quick start

If you don't have Go installed, download [here](https://go.dev/dl/) and follow installation instructions.

### Step by step guide
1. Clone this repository and download dependencies:
```
git clone https://github.com/edyw/gcp-resource-validator.git
go mod download
```

2. Run the server:
```
go run gcp-resource-validator
```
You should see this from the server last line:
```
...
[GIN-debug] Listening and serving HTTP on :8080
```

3. Simulate PubSub Push Subscription message using curl. Use another terminal:
```
curl -X POST localhost:8080/validator -d @./test-asset/storage_location_us_msg.json
```
You should this output from first terminal:
```
Validating asset: //storage.googleapis.com/bucket-test-1
Violation 1 (high-GCPStorageLocationConstraintV1.allow_some_storage_location): //storage.googleapis.com/bucket-test-1 is in a disallowed location.
```


## Full article
Design reference, GCP deployment with CAI Feed, Pub/Sub and Cloud Run, and step by step guide on the code: [Evaluating your GCP Resource realtime](https://medium.com/@lumen1e37/evaluating-your-gcp-resource-realtime-57a0f25d4587)
