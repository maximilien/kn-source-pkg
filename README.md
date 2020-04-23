# source

Knative Client plugin `source`

## Usage

### source

Knative eventing {{.Name}} source plugin

#### Synopsis

Manage your Knative {{.Name}} eventing sources

#### Options

```
  -h, --help   help for source
```

#### SEE ALSO

* [source create](#source-create)	 - create {{.Name}} source
* [source delete](#source-delete)	 - delete {{.Name}} source
* [source describe](#source-describe)	 - describe {{.Name}} source
* [source update](#source-update)	 - update {{.Name}} source

### source create

create {{.Name}} source

#### Synopsis

create {{.Name}} source

```
source create NAME [flags]
```

#### Examples

```
{{.CreateExample}}
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for create
  -n, --namespace string   Specify the namespace to operate in.
  -s, --sink string        Addressable sink for events
```

#### SEE ALSO

* [source](#source)	 - Knative eventing {{.Name}} source plugin

### source delete

delete {{.Name}} source

#### Synopsis

delete {{.Name}} source

```
source delete NAME [flags]
```

#### Examples

```
{{.DeleteExample}}
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for delete
  -n, --namespace string   Specify the namespace to operate in.
```

#### SEE ALSO

* [source](#source)	 - Knative eventing {{.Name}} source plugin

### source describe

describe {{.Name}} source

#### Synopsis

describe {{.Name}} source

```
source describe NAME [flags]
```

#### Examples

```
{{.DescribeExample}}
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for describe
  -n, --namespace string   Specify the namespace to operate in.
```

#### SEE ALSO

* [source](#source)	 - Knative eventing {{.Name}} source plugin

### source update

update {{.Name}} source

#### Synopsis

update {{.Name}} source

```
source update NAME [flags]
```

#### Examples

```
{{.UpdateExample}}
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for update
  -n, --namespace string   Specify the namespace to operate in.
  -s, --sink string        Addressable sink for events
```

#### SEE ALSO

* [source](#source)	 - Knative eventing {{.Name}} source plugin

## More information
	
* [Knative Client](https://github.com/knative/client)
* [How to contribute a plugin](https://github.com/knative/client-contrib#how-to-contribute-a-plugin)

