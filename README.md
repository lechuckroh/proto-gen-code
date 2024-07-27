# proto-gen-code

Source code generator from protobuf file


## Usage

```shell
# Generate *.ts from *.proto
$ ./protogen-ts \
    --proto test/proto/constants.proto \
    --const test/output/constants.ts \
    --msg test/output/constantMessages.ts \
    --rename Error=ErrorType
```
