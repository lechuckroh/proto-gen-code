# proto-gen-code

Source code generator from protobuf file


## Usage

```shell
# Generate *.ts from *.proto
$ ./protogen-ts \
    --proto _test/constants.proto \
    --const _test/constants.ts \
    --msg _test/constantMessages.ts \
    --rename Error=ErrorType
```
