package main

import (
	"fmt"
	"github.com/caarlos0/env"
)

type DB struct{
	User string `env:"DB_USER"`
}

func main() {
	//var conf DB
	//
	//if err := envConf.Init(&conf); err != nil {
	//	fmt.Printf("err=%s\n", err)
	//}
	//
	//fmt.Println(conf)


	//for _, e := range os.Environ() {
	//	pair := strings.Split(e, "=")
	//	fmt.Println(pair[0])
	//}
	cfg := DB{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}