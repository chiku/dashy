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
