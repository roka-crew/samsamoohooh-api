package config

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

type duration time.Duration

func (d *duration) UnmarshalYAML(value *yaml.Node) error {
	var raw string

	if err := value.Decode(&raw); err != nil {
		return fmt.Errorf("failed to decode YAML node to string: %w", err)
	}

	parsedDuration, err := time.ParseDuration(raw)
	if err != nil {
		return fmt.Errorf("invalid duration format '%s': %w", raw, err)
	}

	*d = duration(parsedDuration)
	return nil
}
