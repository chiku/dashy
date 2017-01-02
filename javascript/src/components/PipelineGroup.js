// javascript/src/components/PipelineGroup.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var Pipeline = require("./Pipeline");

var PipelineGroup = function PipelineGroup() {
    var render = function (pipelines) {
        var groupProps = {
            class: "pipeline-group pipeline-group-" + pipelines.length
        };
        return ["div", groupProps, pipelines.map(function (pipeline) {
            return [Pipeline, pipeline];
        })];
    };

    return {
        render: render
    };
};

module.exports = PipelineGroup;
