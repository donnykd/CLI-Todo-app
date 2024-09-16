package main

import (
	"encoding/json"
	"os"
)

type Storage[G any] struct {
	FileName string
}

// Initialises a new storage of any type
func NewStorage[G any](name string) *Storage[G] {
	return &Storage[G]{FileName: name}
}

func (s *Storage[G]) save(data G) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}
	// 0644 is the permission type
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[G]) load (data *G) error{
	fileData, err := os.ReadFile(s.FileName)

	if err != nil{
		return err
	}

	return json.Unmarshal(fileData, data)
}
