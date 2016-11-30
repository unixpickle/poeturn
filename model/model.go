package model

import (
	"io/ioutil"

	"github.com/unixpickle/serializer"
	"github.com/unixpickle/weakai/neuralnet"
	"github.com/unixpickle/weakai/rnn"
)

const CharCount = 256

func init() {
	var m Model
	serializer.RegisterTypedDeserializer(m.SerializerType(), DeserializeModel)
}

// A Model uses an RNN to write poems.
type Model struct {
	Block rnn.Block
}

// DeserializeModel deserializes a Model from data.
func DeserializeModel(d []byte) (*Model, error) {
	var m Model
	if err := serializer.DeserializeAny(d, &m.Block); err != nil {
		return nil, err
	}
	return &m, nil
}

// LoadModel loads a Model from a file.
func LoadModel(f string) (*Model, error) {
	contents, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return DeserializeModel(contents)
}

// NewModel creates a fresh, untrained Model.
func NewModel() *Model {
	outNet := neuralnet.Network{
		&neuralnet.DenseLayer{
			InputCount:  512,
			OutputCount: CharCount,
		},
		&neuralnet.LogSoftmaxLayer{},
	}
	outNet.Randomize()
	outBlock := rnn.NewNetworkBlock(outNet, 0)
	block := rnn.StackedBlock{
		rnn.NewLSTM(CharCount, 512),
		rnn.NewLSTM(512, 512),
		outBlock,
	}
	return &Model{Block: block}
}

// Save saves the model to a file.
func (m *Model) Save(f string) error {
	data, err := m.Serialize()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f, data, 0755)
}

// SerializerType returns the unique ID used to serialize
// a Model with the serializer package.
func (m *Model) SerializerType() string {
	return "github.com/unixpickle/poeturn/model.Model"
}

// Serialize serializes the model.
func (m *Model) Serialize() ([]byte, error) {
	return serializer.SerializeAny(m.Block)
}
