package config

type Mysql struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Db       string `mapstructure:"db" yaml:"db"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" yaml:"charset"`

	ConnMaxIdleTime string `mapstructure:"connMaxIdleTime" yaml:"connMaxIdleTime"`
	ConnMaxLifeTime string `mapstructure:"connMaxLifeTime" yaml:"connMaxLifeTime"`
	MaxIdleconns    int    `mapstructure:"maxIdleconns" yaml:"maxIdleconns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" yaml:"maxOpenConns"`
}
type Redis struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Db       int    `mapstructure:"db" yaml:"db"`
	PoolSize int    `mapstructure:"poolSize" yaml:"poolSize"`
}
