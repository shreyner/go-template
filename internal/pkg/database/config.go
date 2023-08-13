package database

type DataBaseConfig struct {
	Uri string `yaml:"uri" env:"DATABASE_URI,required"`
}
