package pulse

type Source struct {
	Index uint32

	Name        string
	Description string

	//sample_spec

	ChannelMap        ChannelMap
	OwnerModule       uint32
	Volume            CVolume
	Mute              bool
	MonitorOfSink     uint32
	MonitorOfSinkName string

	//latency pa_usec_t

	Driver string

	//flags pa_source_flags_t

	Proplist map[string]string

	BaseVolume Volume

	//state

	NVolumeSteps uint32
	Card         uint32
	Ports        []PortInfo
	ActivePort   PortInfo

	//n_formats
	//formats
}

func (*Source) SetPort(port uint32) {
}
