package pulse

type Audio struct {
	//Cards
	Sinks   []*Sink
	Sources []*Source

	SinkInputs    []*SinkInput
	SourceOutputs []*SourceOutput

	DefaultSink   *Sink
	DefaultSource *Sink
}
