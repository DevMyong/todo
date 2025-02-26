package database

// Postgresql is struct to connect Postgresql
type Postgresql struct {
	URI string
}

// Connect is method to connect Postgresql
func (p *Postgresql) Connect() error {
	return nil
}

// Disconnect is method to disconnect Postgresql
func (p *Postgresql) Disconnect() error {
	return nil
}

func NewPostgresql(uri string) *Postgresql {
	return &Postgresql{
		URI: uri,
	}
}
