package behaviourflag

type PrintVersionFlag struct {
	set bool
}

func (p *PrintVersionFlag) Set(_ string) error {
	p.set = true
	return nil
}

func (p *PrintVersionFlag) IsSet() bool {
	return p.set
}

func (p *PrintVersionFlag) IsBoolFlag() bool {
	return true
}

func (p *PrintVersionFlag) String() string {
	return ""
}

func (p *PrintVersionFlag) Usage() string {
	return "print the program version"
}
