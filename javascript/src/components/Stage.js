var Stage = function() {
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
