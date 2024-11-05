package t_bot

// Observer interface
type Observer interface {
	Update(message string)
}

// Observable interface
type Observable interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}

func (g *game) AddObserver(observer Observer) {
	g.observers = append(g.observers, observer)
}

func (g *game) RemoveObserver(observer Observer) {
	// Implement observer removal logic
}

func (g *game) NotifyObservers(message string) {
	for _, observer := range g.observers {
		observer.Update(message)
	}
}

