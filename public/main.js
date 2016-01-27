var virtualDom = require("virtual-dom");
var nanoajax = require("nanoajax");

var config = window.config;

var h = virtualDom.h;
var diff = virtualDom.diff;
var createElement = virtualDom.create;
var patch = virtualDom.patch;

var requestBody = JSON.stringify({
    url: config.url,
    interests: config.interests
});
var interval = config.interval;

var buildPipelineTree = function(pipeline) {
    var stageNodes = pipeline.stages.map(function(stage) {
        return h("div", {
            className: "stage " + stage.status.toLowerCase()
        }, [stage.name]);
    });

    return h("div", {
        className: "pipeline"
    }, [
        h("div", {
            className: "stage-container"
        }, stageNodes),
        h("div", {
            className: "pipeline-name"
        }, [pipeline.name])
    ]);
};

var renderPipeline = function(pipelines) {
    return h("div", {
        id: "dashboard"
    }, pipelines.map(buildPipelineTree));
};

var renderError = function(message) {
    return h("div", {
        id: "dashboard",
        className: "error"
    }, message);
};


var isSuccess = function(code) {
    return code >= 200 && code <= 299;
};

var dash = function() {
        nanoajax.ajax({
            url: "/dashy",
            type: "POST",
            body: requestBody
        }, function(code, responseText, request) {
            if (isSuccess(code)) {
                newTree = renderPipeline(JSON.parse(responseText));
            } else {
                newTree = renderError(responseText || "Error!");
            }
            patches = diff(tree, newTree);
            rootNode = patch(rootNode, patches);
            tree = newTree;
            console.log("tick!");
        });
    },

    tree = h("div", {
        id: "dashboard"
    }),
    newTree,
    patches,
    rootNode = createElement(tree);

document.body.appendChild(rootNode);
dash();
setInterval(dash, interval);
