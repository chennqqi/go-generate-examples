# gogen-avro example: Generate Go code to serialize and deserialize Avro schemas

about avro:
- https://www.jianshu.com/p/a5c0cbfbf608
- https://link.jianshu.com/?t=http://avro.apache.org/docs/current/spec.html
- http://www.cloudera.com/blog/2009/11/avro-a-new-format-for-data-interchange/ 
- https://www.cnblogs.com/wqbin/p/11228188.html

```
#Generate Go types for definitions and responses in a Swagger V2 spec.
Use x-nullable: true to make something a pointer, defaults to false
Use required on the object to make a field not omitempty, defaults to omitempty.
Use title of a schema to set the Go struct of field name
```

```Go
go install -u  github.com/actgardner/gogen-avro/v7/cmd/gogen-avro@latest
go generate
go build
```