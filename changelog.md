# Notes for Haron & Cameron
Drawing inspiration from Haron's presentation of his version of the GoLang + GraphQL API, I decided to create my own repo 
applying my own patterns & architecture. Since Haron was the first to complete his version, I've decided to iterate and
build on top of his repo. Midway through build I felt it'd be useful to have a document highlighting the differences.
The documentation is brokedown into following sections:
- **Refactor:** - This is code I decided to change because of technical terminology, different architectural decisions, etc.
- **Additions:** - This is new functional code I decided to add, which includes a description of its purpose.
- **Subtractions:** — This is code I decided to omit that otherwise existed in Haron's version.

I feel we are a few iterations before launching a production version of the GoLang API, so by no means will this repo
be the final version. I invite discussion and collaboration in order to continue building & architecting a better version 
of the API.

## Refactor
### Renamed Controllers to Repositories

**Controller:**
> a controller is the go-between for models and views. It accepts the action (as presented by a view) that the user has 
> chosen, determines which models are relevant to this action, and passes those models on to one or more views for
> presentation.

**Repository:**
> The repository is a gateway between your domain/business layer and a data mapping layer, which is the layer that 
> accesses the database and does the operations. Basically the repository is an abstraction to you database access.

We were using controllers as repositories, so I decided to rename them accordingly. Furthermore, I don't believe we really
even need controllers as our resolvers effectively fill that role. This means instead of `user.controller.go` we are now
using `user.repository.go` which lives in the `src/repositories` directory. [Reference Link](https://softwareengineering.stackexchange.com/a/337897). 

### Renamed queries & mutations in the graph
I decided to rename some queries to be a bit more specific and moved the schema to a single file for the time being.
Furthermore, I decided to add a few missing fields to a couple of types.

### Moved `middleware.go` to its own directory
We will have more than just one middleware file, so I decided to create a middleware directory and rename the middleware
accordingly. So now instead of having a `middleware.go` file that only handles auth, I've decided to create a 
`src/middleware/authentication.middleware.go`.

### Relocated helper folders
As a backend dev, I've learned having a single umbrella of helper files never ends well. So I decided to adopt the service
pattern, common among GoLang applications that creates directories that are relevant to the service functionality. For example,
instead of having the `jwt/` and `hasher/` directories under helpers, they've now been moved to the `auth/` service directory.

## Addition
### Added an SQL Compiler
I agree with Haron's approach of not using an ORM as he stated the following: 
- It will take time for developers to learn the ORM
- We fully control everything about our queries
- No mixing & matching Raw Queries/ORM

The biggest issue of not using an orm will be organization, so I decided to use an SQL compiler which will essentially
allow us to put our queries into `.sql` files and then compile it to Go code. This is super useful for organization 
as we will have our `.sql` code in a single directory and have our `.go` code in another directory. Here is an example:

Create an sql file `database/queries/mysql/user_queries.sql`
```mysql
-- name: FindUserByID :one
SELECT * FROM users
WHERE id = $1 AND deleted_at IS NULL LIMIT 1;
```
will create and compile to `src/db/user_queries.sql.go`
```go
const findUserByID = `-- name: FindUserByID :one
SELECT id, uuid, first_name, last_name, email, oauth_provider, oauth_identifier, password, avatar, last_seen_at, last_seen_location, created_at, updated_at, deleted_at FROM users
WHERE id = $1 AND deleted_at IS NULL LIMIT 1
`

func (q *Queries) FindUserByID(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.OauthProvider,
		&i.OauthIdentifier,
		&i.Password,
		&i.Avatar,
		&i.LastSeenAt,
		&i.LastSeenLocation,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
```

Read this [blog post](https://conroy.org/introducing-sqlc) to learn how SQLC will help us. Furthermore, here is the 
[documentation](https://docs.sqlc.dev/en/latest/tutorials/getting-started.html)
### Docker Support
I added Docker to ensure our development environments are symmetric. Our docker compose supports: 
- `web` the actual API application itself
- `mysql` for our database
- `redis` for handling a future redis queue & cache
- `mailhog` for a local smtp server

To get started, just install docker and run `docker-compose build` then `docker-compose up`. 
## Subtract
### Removed resources
The SQL Compiler reads our database tables and automatically generates our models. So I felt we no longer needed the 
resources structs.
