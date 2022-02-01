package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/common"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
)

// This is a sample application showing how to connect to a plumber server's gRPC service
// Please see https://docs.batch.sh/plumber/server-mode for more information on how to run a plumber server

const (
	// By default, plumber's gRPC server runs on port 9090. You can change it via the
	// PLUMBER_SERVER_GRPC_LISTEN_ADDRESS environment variable.
	PlumberInstanceAddress = "localhost:9090"

	// Plumber API auth token. This can be changed by setting the PLUMBER_SERVER_AUTH_TOKEN environment variable
	PlumberAuthToken = "batchcorp"

	// Which topic to read from your local kafka instance
	KafkaTopic = "test-topic"

	// BatchCollectionToken is the token used to authenticate with the Batch Collectors
	// You can obtain one by creating a new collection in https://console.batch.sh
	BatchCollectionToken = "unset"
)

func main() {
	// Dial plumber server
	dialCtx, dialCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer dialCancel()

	conn, err := grpc.DialContext(dialCtx, PlumberInstanceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not dial GRPC server: %s", err)
	}

	log.Println("Connected to Plumber Server gRPC API")

	defer conn.Close()

	client := protos.NewPlumberServerClient(conn)

	// Every request will require a common.Auth message with the authorization token set for plumber
	auth := &common.Auth{Token: PlumberAuthToken}

	//First create a new connection
	connRequest := &protos.CreateConnectionRequest{
		Auth: auth,
		Options: &opts.ConnectionOptions{
			Name: "Local kafka connection",
			Conn: &opts.ConnectionOptions_Kafka{
				Kafka: &args.KafkaConn{
					Address:        []string{"localhost:9092"},
					TimeoutSeconds: 10,
				},
			},
		},
	}

	connResp, err := client.CreateConnection(context.Background(), connRequest)
	if err != nil {
		log.Fatalf("Could not complete CreateConnection request: %s", err)
	}

	log.Printf("Connection created: %s, creating relay. It may take a few seconds to start", connResp.ConnectionId)

	// Now we can create a relay with the newly created connection's ID
	relayResp, err := client.CreateRelay(context.Background(), &protos.CreateRelayRequest{
		Auth: auth,
		Opts: &opts.RelayOptions{
			CollectionToken: BatchCollectionToken,
			ConnectionId:    connResp.ConnectionId,
			Kafka: &opts.RelayGroupKafkaOptions{
				Args: &args.KafkaRelayArgs{
					Topics: []string{KafkaTopic},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Could not complete CreateRelay request: %s", err)
	}

	log.Printf("Relay created: %s", relayResp.RelayId)
	log.Printf("Relaying for 10 seconds")

	// Relay for a few seconds
	time.Sleep(10 * time.Second)

	// We're done relaying, let's stop
	_, err = client.StopRelay(context.Background(), &protos.StopRelayRequest{
		Auth:    auth,
		RelayId: relayResp.RelayId,
	})
	if err != nil {
		log.Fatalf("Could not complete StopRelay request: %s", err)
	}

	log.Println("Relay stopped")

	// You can keep the relay and connection stored in plumber if you need to use it again
	// Relays can be restarted using the ResumeRelay() method
	// For demo purposes, we'll just clean things up

	// Clean up the relay
	_, err = client.DeleteRelay(context.Background(), &protos.DeleteRelayRequest{
		Auth:    auth,
		RelayId: relayResp.RelayId,
	})

	log.Println("Relay deleted")

	// Clean up the connection
	_, err = client.DeleteConnection(context.Background(), &protos.DeleteConnectionRequest{
		Auth:         auth,
		ConnectionId: connResp.ConnectionId,
	})

	log.Println("Connection deleted")

	// All Done! Easy-peasy
}
