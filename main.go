package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Data struct {
	Endpoints []string `json:"endpoints"`
}

type Result struct {
    Endpoint string
    Status   string
    Duration time.Duration
    Err      error
}


func main() {
	filePath := flag.String("file", "data.json", "Path to the JSON file")
	fileType := flag.String("type", "json", "Type of the file (json or yaml)")

	flag.Parse()

	var wg sync.WaitGroup

	raw_data, err := openFile(*filePath, *fileType)

	if err != nil {
		panic(err)
	}

	// should be a parseFile func because we can add yaml files too
	// too lazy to add it now...
	data, err := parseJson(raw_data)

	if err != nil {
		panic(err)
	}


	results := make(chan Result, len(data.Endpoints))

	for _, endpoint := range data.Endpoints {
		wg.Add(1)
		go func (endpoint string) {
			defer wg.Done()

			results <- callEndpoint(endpoint)

		}(endpoint)
	}
	wg.Wait()
	close(results)

	for result := range results {
    if result.Err != nil {
        fmt.Printf("FAIL %s: %v\n",
            result.Endpoint, result.Err)
        continue
    }

    fmt.Printf("OK %s: %s (%v)\n",
        result.Endpoint,
        result.Status,
        result.Duration)
}
}

func parseJson(raw_data []byte) (Data, error) {
	var data Data

	err := json.Unmarshal(raw_data, &data)

	if err != nil {
		return Data{}, err
	}

	return data, nil
}

func callEndpoint(endpoint string) Result {
	start := time.Now()

	result := Result{Endpoint: endpoint}

	resp, err := client.Get(endpoint)
	result.Duration = time.Since(start)

	if err != nil {
		result.Err = err

		return result
	}

	defer resp.Body.Close()

	result.Status = resp.Status

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		result.Err = fmt.Errorf(
        "unexpected HTTP status: %s",
            resp.Status,
    	)
    }


	return result
}

func openFile(filePath string, fileType string) ([]byte, error) {
	switch fileType {
	case "json":
		raw_data, err := readJsonFile(filePath)
		if err != nil {
			return nil, err
		}
		return raw_data, nil
	case "yaml":
		fmt.Println("YAML file type is not supported yet.")
		// raw_data, err := readYamlFile(filePath)
		// if err != nil {
		// 	panic(err)
		// }
		return nil, errors.New("YAML file type is not supported yet")
	default:
		fmt.Println("Unsupported file type. Please use 'json' or 'yaml'.")
		return nil, errors.New("unsupported file type")
	}
}


func readJsonFile(filePath string) ([]byte, error) {
	currentDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	path := filepath.Join(currentDir, filePath)

	raw_data, error := os.ReadFile(path)

	if error != nil {
		return nil, error
	}

	return raw_data, nil

}