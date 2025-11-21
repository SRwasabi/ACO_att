package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadMatrix(path string) ([]MatrixConfig, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	var matrix []MatrixConfig
	if err := json.Unmarshal(content, &matrix); err == nil {
		return matrix, nil
	}
	
	var single MatrixConfig
	if err := json.Unmarshal(content, &single); err == nil {
		return []MatrixConfig{single}, nil
	}
	
	return nil, fmt.Errorf("formato JSON inválido (deve ser objeto ou lista de objetos)")
}