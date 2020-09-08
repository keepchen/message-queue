package queue

import "log"

func init() {
	pool = &DBSet{db: &map[string]*queue{"db0": &queue{}}}
	log.Println("initialize queue DBSet...")
}
