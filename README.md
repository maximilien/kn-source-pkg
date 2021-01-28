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
* [source list](#source-list)	 - list {{.Name}} source
* [source update](#source-update)	 - update {{.Name}} source

### source create

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
  -s, --sink string        Addressable sink for events. You can specify a broker, channel, Knative service or URI. Examples: '--sink broker:nest' for a broker 'nest', '--sink channel:pipe' for a channel 'pipe', '--sink https://event.receiver.uri' for an URI with an 'http://' or 'https://' schema, '--sink ksvc:receiver' or simply '--sink receiver' for a Knative service 'receiver'. If a prefix is not provided, it is considered as a Knative service.
```

#### SEE ALSO

* [source](#source)	 - Knative eventing {{.Name}} source plugin

### source delete

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

### source list

list {{.Name}} source

```
source list NAME [flags]
```

#### Examples

```
{{.ListExample}}
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for list
  -n, --namespace string   Specify the namespace to operate in.
```

#### SEE ALSO

* [source](#source)	 - Knative eventing {{.Name}} source plugin

### source update

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
  -s, --sink string        Addressable sink for events. You can specify a broker, channel, Knative service or URI. Examples: '--sink broker:nest' for a broker 'nest', '--sink channel:pipe' for a channel 'pipe', '--sink https://event.receiver.uri' for an URI with an 'http://' or 'https://' schema, '--sink ksvc:receiver' or simply '--sink receiver' for a Knative service 'receiver'. If a prefix is not provided, it is considered as a Knative service.
```

#### SEE ALSO

* [source](#source)	 - Knative eventing {{.Name}} source plugin

## More information
	
* [Knative Client](https://github.com/knative/client)
* [How to contribute a plugin](https://github.com/knative/client-contrib#how-to-contribute-a-plugin)

