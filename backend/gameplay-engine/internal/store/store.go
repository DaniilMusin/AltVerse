package store

type Postgres struct{}

type Redis struct{}

func MustPostgres(dsn string) *Postgres { return &Postgres{} }

func MustRedis(url string) *Redis { return &Redis{} }
