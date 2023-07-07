protoc.exe --go_out=../ *.proto
protoc --go-grpc_out=require_unimplemented_servers=false:../ pitaya.proto
cd ../protos/
del pitaya.pb.go
ren pitaya_grpc.pb.go pitaya.pb.go
mockgen -source=pitaya.pb.go -destination=mocks/pitaya.go
cd ../cluster
mockgen -source=cluster.go -destination=mocks/cluster.go -package=mocks
pause