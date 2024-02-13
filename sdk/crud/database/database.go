package database


type Database struct {
	Adapter 	 Adapter
	Migrations Migrator
	Seeder 		 ORMSeeder
}

func New(adapter Adapter) *Database {
	return &Database{
		Adapter: adapter,
	}
}

func (db *Database)Init() {
	db.Migrations.Run()
	db.Seeder.Run()	
}

func (db *Database)AddSeeder(seeder ORMSeeder) {
	db.Seeder = seeder
}

func (db *Database)Seed() {
	db.Seeder.Run()
}

func (db *Database)AddMigrations(migrator Migrator) {
	db.Migrations = migrator
}

func (db *Database)Migrate() {
	db.Migrations.Run()
}

