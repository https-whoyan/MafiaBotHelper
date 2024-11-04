package bot

type Bots []*Bot

func (b *Bots) Len() int { return len(*b) }

func (b *Bots) GetByNum(num int) *Bot {
	if b.Len() < num {
		return (*b)[num]
	}
	return nil
}

func (b *Bots) GetByName(name string) *Bot {
	// Todo...
	return nil
}
