package types

type Movement struct {
	Count       int
	Source      int
	Destination int
}

type Configuration struct {
	Stacks    map[int][]rune
	Movements []*Movement
}

func (c *Configuration) Move(retainOrder bool) {
	for _, movement := range c.Movements {
		source := c.Stacks[movement.Source]
		destination := c.Stacks[movement.Destination]

		if retainOrder {
			sourceSeparator := len(source) - movement.Count
			destination = append(destination, source[sourceSeparator:]...)
			source = source[0:sourceSeparator]
		} else {
			for i := 0; i < movement.Count; i++ {
				sourceLast := len(source) - 1
				destination = append(destination, source[sourceLast])
				source = source[0:sourceLast]
			}
		}

		c.Stacks[movement.Source] = source
		c.Stacks[movement.Destination] = destination
	}
}

func (c *Configuration) Result() string {
	result := make([]rune, len(c.Stacks))
	for position, stack := range c.Stacks {
		lastIndex := len(stack) - 1
		result[position-1] = stack[lastIndex]
	}
	return string(result)
}
