package day05

func TopStacksFromContainer(c *Container) []string {

	for _, instr := range c.Instructions {
		amount, fcolumn, tcolumn := instr[0], instr[1], instr[2]

		fc := c.Stacks[fcolumn]
		elemns := fc[len(fc)-amount:]

		var relemns = make([]string, 0)

		for i := len(elemns) - 1; i >= 0; i-- {
			relemns = append(relemns, elemns[i])
		}

		c.Stacks[tcolumn] = append(c.Stacks[tcolumn], relemns...)
		c.Stacks[fcolumn] = fc[:len(fc)-amount]
	}

	var out = make([]string, 0)

	for i := 1; i <= len(c.Stacks); i++ {
		out = append(out, c.Stacks[i][len(c.Stacks[i])-1:]...)
	}

	return out
}
