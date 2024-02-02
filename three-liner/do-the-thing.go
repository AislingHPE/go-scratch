package three

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

func YamlStuff() {
	buffer := make(map[string]interface{})
	yamlString := "schema:\n  b: A"
	err := yaml.Unmarshal([]byte(yamlString), buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", buffer)
	_, err = json.Marshal(buffer)
	if err != nil {
		fmt.Println(err)
	}
}
