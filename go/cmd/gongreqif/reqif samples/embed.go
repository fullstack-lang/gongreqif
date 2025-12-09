package samples

import "embed"

//go:embed "reqif samples_ReqIF for Wheeled Tennis Ball Drone-renderingConf.go1"
var SampleRenderingConf embed.FS

//go:embed "ReqIF for Wheeled Tennis Ball Drone.reqifz"
var SampleReqIFz embed.FS
