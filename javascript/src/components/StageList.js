// javascript/src/components/StageList.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2016. All rights reserved
// License::   MIT

var Stage = require("./Stage");

var stageContainerProps = {
    "class": "stage-container"
};

var StageList = function() {
    var render = function(stages) {
        return ["div", stageContainerProps, stages.map(function(stage) {
            return [Stage, stage];
        })];
    };

    return {
        render: render
    };
};

module.exports = StageList;