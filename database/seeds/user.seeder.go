package seeds

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/solomon04/go-docker/src/auth/hasher"
	"github.com/uniplaces/carbon"
	"math/rand"
)

func (s Seed) UserSeed() {
	println("Running user seeder")
	for i := 0; i < 100; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO users(uuid, first_name, last_name, email, password, created_at, updated_at) VALUES (?,?,?,?,?,?,?)`)
		// execute query
		firstName := faker.FirstName()
		lastName := faker.LastName()
		email := fmt.Sprintf("%s.%s_%d@mail.com", firstName, lastName, rand.Int())
		_, err := stmt.Exec(faker.UUIDHyphenated(), firstName, lastName, email, hasher.Make( "secret"), carbon.Now().DateTimeString(), carbon.Now().DateTimeString())
		if err != nil {
			panic(err)
		}
	}
}
