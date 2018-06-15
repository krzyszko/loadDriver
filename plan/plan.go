package plan

type Component interface {
	Run() error
	Init(map[string][]byte) error
	New([]byte) Component
}

type Plan struct {
	components []Component
	registry   map[string][]byte
}

func (p *Plan) AddComponent(c Component) {
	p.components = append(p.components, c)
}

func (p *Plan) Run() error {
	p.registry = make(map[string][]byte)
	defer func() { p.registry = nil }()
	for _, c := range p.components {
		err := c.Init(p.registry)
		if err != nil {
			return err
		}
	}
	for _, c := range p.components {
		err := c.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
