package config

type MongoDB struct {
	MongoDBURI string `mapstructure:"mongodb-url" json:"mongodburl" yaml:"mongodb-url"`
	Database   string `mapstructure:"database" json:"database" yaml:"database"`
	Timeout    int64  `mapstructure:"time-out" json:"timeout" yaml:"time-out"`
	ConnectNum int    `mapstructure:"connect-num" json:"connectnum" yaml:"connect-num"`
}
