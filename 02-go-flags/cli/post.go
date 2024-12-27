package cli

import (
	"encoding/json"
	"fmt"
)

type ListPostCommand struct{}

func (l *ListPostCommand) Execute(args []string) error {
	body, err := GetData(apiURL + "/posts")
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	var todos []map[string]interface{}
	if err := json.Unmarshal(body, &todos); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	for _, todo := range todos {
		fmt.Printf("%v: %v (completed: %v)\n", todo["id"], todo["title"], todo["body"])
	}
	return nil
}
