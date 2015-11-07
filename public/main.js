var virtualDom = require("virtual-dom");
var nanoajax = require("nanoajax");

var h = virtualDom.h;
var diff = virtualDom.diff;
var createElement = virtualDom.create;
var patch = virtualDom.patch;

var config = window.config;

var render = function(pipelines) {
      var pipelinesNodes = pipelines.map(function(pipeline){
          var stageNodes = pipeline.stages.map(function(stage){
              return h("div", {className: "stage " + stage.status.toLowerCase()}, [stage.name]);
          });

          return h("div", {className: "pipeline"}, [
              h("div", {className: "stage-container"}, stageNodes),
              h("div", {className: "pipeline-name"}, [pipeline.name])
          ]);
      });

      return h("div", {id: "dashboard"}, pipelinesNodes);
  },

  dash = function() {
      nanoajax.ajax({
          url: "/dashy",
          type: "POST",
          body: JSON.stringify({url: config.url, interests: config.interests})
    }, function(code, responseText, request) {
      if (code >= 200 && code <= 299) {
          newTree = render(JSON.parse(responseText));
      } else {
          newTree = h("div", {id: "dashboard", className: "error"}, [responseText || "Error!"]);
      }
      patches = diff(tree, newTree);
      rootNode = patch(rootNode, patches);
      tree = newTree;
      console.log("tick!");
    });
  },

  tree = h("div", {id: "dashboard"}),
  newTree,
  patches,
  rootNode = createElement(tree);

document.body.appendChild(rootNode);
dash();
setInterval(dash, config.interval);
