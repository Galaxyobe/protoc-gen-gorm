// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/galaxyobe/protoc-gen-gorm/plugin"
	"strconv"
	"flag"
)

func main() {

	generatePath := "."
	useGogoImport := false
	ppid := 0

	// parse flag
	inject := flag.Bool("inject", false, "inject *.pb.go for gorm")
	flag.StringVar(&generatePath, "inject-path", ".", "inject path of *.pb.go")
	flag.IntVar(&ppid, "ppid", -1, "inject parent pid")
	
	flag.Parse()

	myPlugin := plugin.NewPlugin(useGogoImport, generatePath)

	if *inject {
		myPlugin.InjectIgnore(ppid)
		return
	}

	// generator
	gen := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		gen.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, gen.Request); err != nil {
		gen.Error(err, "parsing input proto")
	}

	if len(gen.Request.FileToGenerate) == 0 {
		gen.Fail("no files to generate")
	}

	// Match parsing algorithm from Generator.CommandLineParameters
	for _, parameter := range strings.Split(gen.Request.GetParameter(), ",") {
		kvp := strings.SplitN(parameter, "=", 2)
		if len(kvp) != 2 {
			continue
		}
		switch kvp[0] {
		case "gogoimport":
			useGogoImport, err = strconv.ParseBool(kvp[1])
			if err != nil {
				gen.Error(err, "parsing gogoimport option")
			}
			myPlugin.UseGogoImport = useGogoImport
		case "generate_path":
			myPlugin.GeneratePath = kvp[1]
		}

	}

	gen.CommandLineParameters(gen.Request.GetParameter())

	gen.WrapTypes()
	gen.SetPackageNames()
	gen.BuildTypeNameMap()

	gen.GeneratePlugin(myPlugin)

	for i := 0; i < len(gen.Response.File); i++ {
		gen.Response.File[i].Name = proto.String(strings.Replace(*gen.Response.File[i].Name, ".pb.go", ".gorm.go", -1))
	}

	// Send back the results.
	data, err = proto.Marshal(gen.Response)
	if err != nil {
		gen.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		gen.Error(err, "failed to write output proto")
	}
	err = myPlugin.InjectIgnoreFork()
	if err != nil {
		gen.Error(err, "failed to start inject proto fork")
	}
}
