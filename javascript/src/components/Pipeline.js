// javascript/src/components/Pipeline.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2016. All rights reserved
// License::   MIT

var StageList = require("./StageList");
var PipelineName = require("./PipelineName");

var Pipeline = function() {
    var pipelineProps = {
        "class": "pipeline"
    };
    var render = function(pipeline) {
        return ["div", pipelineProps, [
            [StageList, pipeline.stages],
            [PipelineName, pipeline.name]
        ]];
    };

    return {
        render: render
    };
};

module.exports = Pipeline;
