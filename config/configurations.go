package config

type Configurations struct {
	ElasticHost string `required:"true"`
	ElasticPort string `required:"true"`
}
