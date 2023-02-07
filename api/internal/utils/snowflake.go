package utils

import "github.com/bwmarrin/snowflake"

func GenerateSnowflakeID() (snowflake.ID, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}

	id := node.Generate()
	return id, nil
}
