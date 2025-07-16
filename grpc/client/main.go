package main

import (
	"client/client_api"
	"log"
	"proto_api/pkg/grpc/v1/serial_api"
)

func main() {
	client, err := client_api.NewSerialClient("localhost:5001")
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer client.Close()

	// Создание сериала
	newSerial := &serial_api.CreateSerialRequest{
		Title:            "Breaking Bad",
		FileId:           12345,
		Description:      "",
		Rating:           9.5,
		Duration:         45.0,
		Sort:             1,
		ProductionPeriod: "2008-2013",
		Quality:          "HD",
	}

	created, err := client.CreateSerial(newSerial)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Created Serial: %+v\n", created)

	// Получение сериала
	serialID := created.GetId()
	fetched, err := client.GetSerial(serialID)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	log.Printf("Fetched Serial: %+v\n", fetched)

	// Обновление сериала
	fetched.Rating = 9.7
	updated, err := client.UpdateSerial(fetched)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Updated Serial: %+v\n", updated)

	// Получение всех сериалов
	allSerials, err := client.GetAllSerials()
	if err != nil {
		log.Fatalf("GetAll failed: %v", err)
	}
	log.Println("All Serials:")
	for _, s := range allSerials {
		log.Printf(" - %d: %s (%.1f)", s.GetId(), s.GetTitle(), s.GetRating())
	}

	// Удаление сериала
	err = client.DeleteSerial(serialID)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Println("Serial deleted successfully")
}
