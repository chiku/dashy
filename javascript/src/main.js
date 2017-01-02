// javascript/src/main.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var domChanger = require("domchanger");
var nanoajax = require("nanoajax");
var PipelineGroupLister = require("./components/PipelineGroupList");

var config = window.config || {};

var requestBody = JSON.stringify({
    url: config.url,
    interests: config.interests
});
var interval = config.interval || 30000;
var groupSize = config.groupSize || 1;
var baseSize = config.baseSize || 6;

var isSuccess = function (code) {
    return code >= 200 && code <= 299;
};

var initCSS = function () {
    var a1 = baseSize;
    var a2 = baseSize * 0.986111;
    var a3 = baseSize * 0.833334;
    var a4 = baseSize * 0.125;

    var cssContent = ".pipeline-name {height: " + a1 + "vmax; margin-top: " + (-a2) + "vmax; font-size: " + a3 + "vmax;}" +
        " .stage-container {height: " + a1 + "vmax;}" +
        " .stage {font-size: " + a4 + "vmax;}";

    var style = document.createElement("style");
    style.type = "text/css";
    style.innerHTML = cssContent;
    document.getElementsByTagName("head")[0].appendChild(style);
};

var asError = function (message, code) {
    var defaultedMessage = (code === 0) ? "Error - server down" : message;
    return [{
        name: defaultedMessage,
        stages: [{
            name: "Error",
            status: "Failed"
        }]
    }];
};

var Dashy = function (emit, refresh) {
    var pipelines = [];
    var responseHandler = function (code, responseText) {
        if (isSuccess(code)) {
            pipelines = JSON.parse(responseText);
        } else {
            pipelines = asError(responseText, code);
            /* eslint "no-console": 0 */
            console.error(responseText);
        }
        refresh();
    };
    var ajaxOptions = {
        url: "/dashy",
        type: "POST",
        body: requestBody
    };
    var tick = function () {
        nanoajax.ajax(ajaxOptions, responseHandler);
    };
    var render = function () {
        return [PipelineGroupLister(groupSize), pipelines];
    };

    tick();
    setInterval(tick, interval);
    return {
        render: render
    };
};

initCSS();
domChanger(Dashy, document.getElementById("app")).update();
