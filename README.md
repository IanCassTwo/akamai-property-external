# akamai-property-external

This is a utility that allows you to utilize an existing Akamai Property Manager configuration as a "Golden Master".

```xml
data "external" "example" {
  program = ["./akamai-property-external"]
  query = {
    edgerc = "/home/icass/.edgerc"
    section = "papi"
    propertyId = "prp_585981"
    groupId = "${data.akamai_group.group.id}"
  }
}
```

and then in your akamai_property resource...

```xml
...
  rules       = data.external.example.result.Rules
...
```

Note, you'll need to identify the property id of your golden master configuration. You can do this via the GUI or you can use the Akamai Property CLI.
