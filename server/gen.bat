cd pbfiles && protoc --go_out=plugins=grpc:../services prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services prod.proto
cd ..