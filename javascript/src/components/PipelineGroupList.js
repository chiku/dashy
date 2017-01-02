// javascript/src/components/PipelineGroupList.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var PipelineGroup = require("./PipelineGroup");

var PipelineGroupList = function PipelineGroupList(groupSize) {
    return function () {
        var render = function (pipelines) {
            var len = pipelines.length;
            var groups = [];
            var i;
            for (i = 0; i < len; i += groupSize) {
                groups.push(pipelines.slice(i, i + groupSize));
            }
            return groups.map(function (group) {
                return [PipelineGroup, group];
            });
        };

        return {
            render: render
        };
    };
};

module.exports = PipelineGroupList;
