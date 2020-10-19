cd pbfiles && protoc --go_out=plugins=grpc:../services prod.proto

protoc --go_out=plugins=grpc:../services models.proto
protoc --go_out=plugins=grpc:../services orders.proto
protoc --go_out=plugins=grpc:../services users.proto

protoc --grpc-gateway_out=logtostderr=true:../services prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services orders.proto

cd ..