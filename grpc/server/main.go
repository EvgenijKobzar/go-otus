package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"otus/pkg/lib"
	"otus/pkg/lib/mapstructure"
	serialpb "proto_api/pkg/grpc/v1/serial_api"
	"time"
)

// gRPC сервер
type serialServer struct {
	serialpb.UnimplementedSerialServiceServer
}

func newSerialServer() *serialServer {
	return &serialServer{}
}

func (s *serialServer) CreateSerial(ctx context.Context, req *serialpb.CreateSerialRequest) (*serialpb.Serial, error) {
	serial := &serialpb.Serial{
		Title:            req.GetTitle(),
		FileId:           req.GetFileId(),
		Description:      req.GetDescription(),
		Rating:           req.GetRating(),
		Duration:         req.GetDuration(),
		Sort:             req.GetSort(),
		ProductionPeriod: req.GetProductionPeriod(),
		Quality:          req.GetQuality(),
	}

	m, _ := mapstructure.StructToMap(serial)

	mapModel, err := lib.Add(m)
	if err != nil {
		return nil, err
	}
	mapstructure.MapToStruct(mapModel, &serial)

	return serial, nil
}
func (s *serialServer) GetSerial(ctx context.Context, req *serialpb.SerialRequest) (*serialpb.Serial, error) {
	entity, err := lib.Get(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	m, _ := mapstructure.StructToMap(entity)
	var serial serialpb.Serial
	mapstructure.MapToStruct(m, &serial)

	return &serial, nil
}
func (s *serialServer) UpdateSerial(ctx context.Context, req *serialpb.Serial) (*serialpb.Serial, error) {

	serial := &serialpb.Serial{}

	inputFields, _ := mapstructure.StructToMap(req)
	mapModel, err := lib.Update(int(req.GetId()), inputFields)
	if err != nil {
		return nil, err
	}
	mapstructure.MapToStruct(mapModel, &serial)

	return serial, nil
}
func (s *serialServer) DeleteSerial(ctx context.Context, req *serialpb.SerialRequest) (*emptypb.Empty, error) {
	err := lib.Delete(int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *serialServer) GetAllSerials(ctx context.Context, _ *emptypb.Empty) (*serialpb.SerialList, error) {

	items, err := lib.GetList()
	if err != nil {
		return nil, err
	}

	var serials []*serialpb.Serial
	for _, item := range items {
		m, _ := mapstructure.StructToMap(item)
		var serial serialpb.Serial
		mapstructure.MapToStruct(m, &serial)

		serials = append(serials, &serial)
	}

	return &serialpb.SerialList{Serials: serials}, nil
}

func main() {
	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(10 * time.Second),
	)

	serialService := newSerialServer()
	serialpb.RegisterSerialServiceServer(grpcServer, serialService)

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server started on port 5001")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
