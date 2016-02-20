var pipelineNameProps = {
    "class": "pipeline-name"
};

var PipelineName = function() {
    var render = function(name) {
        return ["div", pipelineNameProps, name];
    };

    return {
        render: render
    };
};

module.exports = PipelineName;
