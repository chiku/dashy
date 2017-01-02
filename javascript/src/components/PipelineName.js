// javascript/src/components/PipelineName.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var pipelineNameProps = {
    class: "pipeline-name"
};

var PipelineName = function PipelineName() {
    var render = function (name) {
        return ["div", pipelineNameProps, name];
    };

    return {
        render: render
    };
};

module.exports = PipelineName;
