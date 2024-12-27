package cli

import (
	"encoding/json"
	"fmt"
)

type ListCommand struct{}

type ShowCommand struct {
	ID int `short:"i" long:"id" description:"ToDo ID" required:"true"`
}

func (l *ListCommand) Execute(args []string) error {
	body, err := GetData(apiURL + "/todos")
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	var todos []map[string]interface{}
	if err := json.Unmarshal(body, &todos); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	for _, todo := range todos {
		fmt.Printf("%v: %v (completed: %v)\n", todo["id"], todo["title"], todo["completed"])
	}
	return nil
}

func (s *ShowCommand) Execute(args []string) error {
	url := fmt.Sprintf("%s/%d", apiURL+"/todos", s.ID)
	body, err := GetData(url)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	var todo map[string]interface{}
	if err := json.Unmarshal(body, &todo); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	fmt.Printf("ToDo ID: %v\nTitle: %v\nCompleted: %v\n", todo["id"], todo["title"], todo["completed"])
	return nil
}
