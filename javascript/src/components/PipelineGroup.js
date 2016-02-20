var Pipeline = require("./Pipeline");

var PipelineGroup = function() {
    var render = function(pipelines) {
        var groupProps = {
            "class": "pipeline-group pipeline-group-" + pipelines.length
        };
        return ["div", groupProps, pipelines.map(function(pipeline) {
            return [Pipeline, pipeline];
        })];
    };

    return {
        render: render
    };
};

module.exports = PipelineGroup;
