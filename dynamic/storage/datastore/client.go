package datastore

import (
	"cloud.google.com/go/datastore"
)

type Datastore struct {
	client *datastore.Client
}

func NewDatastore(client *datastore.Client) *Datastore {
	return &Datastore{
		client: client,
	}
}
