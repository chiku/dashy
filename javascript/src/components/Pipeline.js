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
