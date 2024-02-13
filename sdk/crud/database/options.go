package database

type DatabaseOptions func (*DBConfig)

func WithDSN(dsn string) DatabaseOptions {
	return func(config *DBConfig) {
		config.DSN = dsn
	}
}

func WithDBName(dbName string) DatabaseOptions {
	return func(config *DBConfig) {
		config.DBName = dbName
	}
}

func WithDBUser(dbUser string) DatabaseOptions {
	return func(config *DBConfig) {
		config.DBUser = dbUser
	}
}

func WithDBPass(dbPass string) DatabaseOptions {
	return func(config *DBConfig) {
		config.DBPass = dbPass
	}
}

func WithDBHost(dbHost string) DatabaseOptions {
	return func(config *DBConfig) {
		config.DBHost = dbHost
	}
}

func WithDBPort(dbPort int) DatabaseOptions {
	return func(config *DBConfig) {
		config.DBPort = dbPort
	}
}

func WithConfig(config *DBConfig) DatabaseOptions {
	return func(cnf *DBConfig) {
		cnf = config
	}
}