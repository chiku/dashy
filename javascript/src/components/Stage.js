// javascript/src/components/Stage.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2016. All rights reserved
// License::   MIT

var Stage = function Stage() {
    var render = function(stage) {
        var stageProps = {
            "class": "stage " + stage.status.toLowerCase()
        };
        return ["div", stageProps, stage.name];
    };

    return {
        render: render
    };
};

module.exports = Stage;