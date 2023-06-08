module github.com/ssen110/test

go 1.20

require (
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.9
)

require github.com/ssen110/test/types v0.0.0-00010101000000-000000000000 // indirect

replace github.com/ssen110/test/types => ./types
