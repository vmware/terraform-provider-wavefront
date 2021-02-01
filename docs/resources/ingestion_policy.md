---
layout: "wavefront"
page_title: "Wavefront: Ingestion Policy"
description: |-
  Provides a Wavefront Ingestion Policy Resource. This allows ingestion policies to be created, updated, and deleted.
---

# Resource : wavefront_ingestion_policy

Provides a Wavefront Ingestion Policy Resource. This allows ingestion policies to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_ingestion_policy" "basic" {
  name  = "test_ingestion"
  description = "An ingestion policy for testing"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the ingestion policy
* `description` - (Required) The description of the ingestion policy

### Example

```hcl

resource "wavefront_ingestion_policy" "basic" {
  name = "test_ingestion"
  description = "An ingestion policy for testing"
}
```

## Import

ingestion policies can be imported using the `id`, e.g.

```
$ terraform import wavefront_ingestion_policy.basic test_ingestion-1611946841064
```
