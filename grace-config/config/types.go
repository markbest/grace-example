package g

type config struct {
	DB database `toml:"database"`
	Redis redis `toml:"redis"`
}

type database struct {
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type redis struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Db       int    `toml:"db"`
	Password string `toml:"password"`
}
