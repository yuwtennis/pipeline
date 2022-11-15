package main

import (
	"flag"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

// simple_rw is an example that runs job on dataflow.
// It simply writes file from specified input to specified output.
//
// To run the program, execute below command.
//
// go run cmd/simple_rw/main.go  \
//   --project PROJECT_NAME \
//   --runner dataflow \
//   --staging_location STAGING_LOCATION_DIRECTORY  \
//   --zone ZONE_NAME  \
//   --subnetwork NETWORK_URL \
//   --output  OUTPUT_FILE_PATH

import (
	"context"
	"log"
)

var (
	input  = flag.String("input", "gs://apache-beam-samples/shakespeare/kinglear.txt", "File(s) to read.")
	output = flag.String("output", "", "Output file required.")
)

func main() {
	flag.Parse()

	beam.Init()
	p := beam.NewPipeline()
	s := p.Root()

	// Simply transfer data left to right
	pCol := textio.Read(s, *input)
	textio.Write(s, *output, pCol)

	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Failed to execute job : %v ", err)
	}
}
