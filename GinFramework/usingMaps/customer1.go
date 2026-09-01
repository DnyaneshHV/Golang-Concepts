package main

type Customer struct {  // we have given the struct type as it is in the same package , Go is able to see it.
	ID    int
	Name  string
	Email string
}
