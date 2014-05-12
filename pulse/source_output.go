package pulse

/*
#include "dde-pulse.h"
*/
import "C"

type SourceOutput struct{}

func toSourceOutputInfo(info *C.pa_source_output_info) *SourceOutput {
	return nil
}
