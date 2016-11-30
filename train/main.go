package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/unixpickle/poeturn/model"
	"github.com/unixpickle/sgd"
	"github.com/unixpickle/weakai/neuralnet"
	"github.com/unixpickle/weakai/rnn"
	"github.com/unixpickle/weakai/rnn/seqtoseq"
)

const TrainingFraction = 0.9

func main() {
	var sampleFile string
	var outputFile string
	var stepSize float64
	var batchSize int
	var logInterval int
	var maxLen int

	flag.StringVar(&sampleFile, "samples", "", "path to the samples")
	flag.StringVar(&outputFile, "output", "out_net", "network file")
	flag.Float64Var(&stepSize, "step", 0.001, "SGD step size")
	flag.IntVar(&batchSize, "batch", 1, "SGD mini-batch size")
	flag.IntVar(&logInterval, "logint", 4, "log interval")
	flag.IntVar(&maxLen, "maxlen", 1000, "max poem length")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	samples, err := LoadSampleSet(sampleFile, maxLen)
	if err != nil {
		die("Error loading samples:", err)
	}
	net, err := model.LoadModel(outputFile)
	if os.IsNotExist(err) {
		log.Println("Creating model...")
		net = model.NewModel()
	} else if err != nil {
		die("Error loading model:", err)
	}

	costFunc := neuralnet.DotCost{}
	grad := &sgd.Adam{
		Gradienter: &seqtoseq.Gradienter{
			SeqFunc:       &rnn.BlockSeqFunc{B: net.Block},
			Learner:       net.Block.(sgd.Learner),
			CostFunc:      costFunc,
			MaxGoroutines: 1,
			MaxLanes:      batchSize,
		},
	}

	training, validation := sgd.HashSplit(samples, TrainingFraction)
	log.Println("Have", training.Len(), "training and", validation.Len(), "validation.")

	var last sgd.SampleSet
	var iter int
	sgd.SGDMini(grad, training, stepSize, batchSize, func(batch sgd.SampleSet) bool {
		getCost := func(b sgd.SampleSet) float64 {
			return seqtoseq.TotalCostBlock(net.Block, batchSize, b, costFunc)
		}
		if iter%logInterval == 0 {
			var lastCost float64
			if last != nil {
				lastCost = getCost(last)
			}
			last = batch.Copy()
			newCost := getCost(batch)
			sgd.ShuffleSampleSet(validation)
			vc := getCost(validation.Subset(0, batchSize))
			log.Printf("iter %d: validation=%f cost=%f last=%f", iter, vc, newCost, lastCost)
		}
		iter++
		return true
	})

	if err := net.Save(outputFile); err != nil {
		die("Error saving model:", err)
	}
}

func die(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}
