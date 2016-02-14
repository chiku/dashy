var domChanger = require("domChanger");
var nanoajax = require("nanoajax");
var config = window.config || {};

var requestBody = JSON.stringify({
    url: config.url,
    interests: config.interests
});
var interval = config.interval;

var isSuccess = function(code) {
    return code >= 200 && code <= 299;
};

var DashyError = function() {
    var errorProps = {
        id: "dashboard",
        class: "error"
    };
    var render = function(message) {
        return ["div", errorProps, message];
    };

    return {
        render: render
    };
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
    var render = function(pipelines) {
        return ["div", pipelineListProps, pipelines.map(function(pipeline) {
            return [Pipeline, pipeline];
        })];
    };

    return {
        render: render
    };
};

var Dashy = function(emit, refresh) {
    var renderedItem = ["div", "Initializing..."];
    var responseHandler = function(code, responseText, request) {
        if (isSuccess(code)) {
            renderedItem = [PipelineList, JSON.parse(responseText)];
        } else {
            renderedItem = [DashyError, (responseText || "Error!")];
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
        return renderedItem;
    };

    tick();
    setInterval(tick, interval);
    return {
        render: render
    };
};

var instance = domChanger(Dashy, document.getElementById("app"));
instance.update();
