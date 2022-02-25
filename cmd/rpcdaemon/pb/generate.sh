rm -rvf go && mkdir go
rm -rvf java && mkdir java
rm -rvf python && mkdir python
protoc -I=. --go_out=./go/ Block.proto
protoc -I=. --java_out=./java/ Block.proto
protoc -I=. --python_out=./python/ Block.proto
