var domChanger = require("domChanger");
var nanoajax = require("nanoajax");
var config = window.config || {};

var requestBody = JSON.stringify({
    url: config.url,
    interests: config.interests
});
var interval = config.interval || 30000;
var groupSize = config.groupSize || 1;

var isSuccess = function(code) {
    return code >= 200 && code <= 299;
};

var asError = function(message, code) {
    if (code === 0) {
        message = "Error - server down";
    }
    return [{
        name: message,
        stages: [{
            name: "Error",
            status: "Failed"
        }]
    }];
};

var Stage = function() {
    var render = function(stage) {
        var stageProps = {
            class: "stage " + stage.status.toLowerCase()
        };
        return ["div", stageProps, stage.name];
    };

    return {
        render: render
    };
};

var StageList = function() {
    var stageContainerProps = {
        class: "stage-container"
    };
    var render = function(stages) {
        return ["div", stageContainerProps, stages.map(function(stage) {
            return [Stage, stage];
        })];
    };

    return {
        render: render
    };
};

var PipelineName = function() {
    var pipelineNameProps = {
        class: "pipeline-name"
    };
    var render = function(name) {
        return ["div", pipelineNameProps, name];
    };

    return {
        render: render
    };
};

var Pipeline = function() {
    var pipelineProps = {
        class: "pipeline"
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

var PipelineList = function() {
    var pipelineListProps = {
        id: "dashboard"
    };
    var mapper = function(pipeline) {
        return [Pipeline, pipeline];
    };
    var groupProps = {
        class: "pipeline-group pipeline-group-" + groupSize
    };
    var render = function(pipelines) {
        var len = pipelines.length,
            groups = [],
            i,
            group;
        for (i = 0; i < len; i += groupSize) {
            group = pipelines.slice(i, i + groupSize).map(mapper);
            groups.push(["div", groupProps, group]);
        }
        return ["div", pipelineListProps, groups];
    };

    return {
        render: render
    };
};

var Dashy = function(emit, refresh) {
    var pipelines = [];
    var responseHandler = function(code, responseText, request) {
        if (isSuccess(code)) {
            pipelines = JSON.parse(responseText);
        } else {
            pipelines = asError(responseText, code);
        }
        refresh();
        console.log("tick!");
    };
    var ajaxOptions = {
        url: "/dashy",
        type: "POST",
        body: requestBody
    };
    var tick = function() {
        nanoajax.ajax(ajaxOptions, responseHandler);
    };
    var render = function() {
        return [PipelineList, pipelines];
    };

    tick();
    setInterval(tick, interval);
    return {
        render: render
    };
};

var instance = domChanger(Dashy, document.getElementById("app"));
instance.update();
