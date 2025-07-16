package client_api

import (
	"context"
	"log"
	"proto_api/pkg/grpc/v1/serial_api"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SerialClient struct {
	client serial_api.SerialServiceClient
	conn   *grpc.ClientConn
}

func NewSerialClient(serverAddr string) (*SerialClient, error) {
	conn, err := grpc.NewClient(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	client := serial_api.NewSerialServiceClient(conn)

	return &SerialClient{client: client, conn: conn}, nil
}

func (c *SerialClient) Close() {
	c.conn.Close()
}

func (c *SerialClient) CreateSerial(serial *serial_api.CreateSerialRequest) (*serial_api.Serial, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.client.CreateSerial(ctx, serial)
}

func (c *SerialClient) GetSerial(id int64) (*serial_api.Serial, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.client.GetSerial(ctx, &serial_api.SerialRequest{Id: id})
}

func (c *SerialClient) UpdateSerial(serial *serial_api.Serial) (*serial_api.Serial, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.client.UpdateSerial(ctx, serial)
}

func (c *SerialClient) DeleteSerial(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.client.DeleteSerial(ctx, &serial_api.SerialRequest{Id: id})
	return err
}

func (c *SerialClient) GetAllSerials() ([]*serial_api.Serial, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	res, err := c.client.GetAllSerials(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return res.Serials, nil
}
