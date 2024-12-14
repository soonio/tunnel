package mon

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/ssh"
)

type Mongo struct {
	client *mongo.Client
	uri    string
}

func New(uri string) *Mongo {
	return &Mongo{uri: uri}
}

func (d *Mongo) Connect(tunnel *ssh.Client) error {
	var err error
	d.client, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI(d.uri).
		SetConnectTimeout(10*time.Second).
		SetSocketTimeout(10*time.Second).
		SetTimeout(10*time.Second).
		SetDialer(&Dialer{tunnel}),
	)
	return err
}

func (d *Mongo) Use() *mongo.Client {
	return d.client
}

func (d *Mongo) Close() error {
	if d.client != nil {
		return d.client.Disconnect(context.TODO())
	}
	return nil
}
