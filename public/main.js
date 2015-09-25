(function(window, document, $, undefined){
  "use strict";

  var buildGrid = function(pipelines) {
    var dashboard = $("<div>", {id: "dashboard"});
    pipelines.forEach(function(pipeline) {
      var pipelineElem = $("<div>").addClass("pipeline");
      var stageContainerElem = $("<div>").addClass("stage-container").appendTo(pipelineElem);
      var pipelineNameElem = $("<div>").addClass("pipeline-name").text(pipeline.name).appendTo(pipelineElem);

      pipeline.stages.forEach(function(stage) {
        $("<div>").addClass("stage").addClass(stage.status.toLowerCase()).text(stage.name).appendTo(stageContainerElem);
      });

      pipelineElem.appendTo(dashboard);
    });

    $("#dashboard").replaceWith(dashboard);
    $("#dashboard-error").html("");
  };

  var errorGrid = function(pipelines) {
    $("#dashboard").html("");
    $("#dashboard-error").html($("<div>").addClass("error").text("Error!"));
  };

  var dash = function() {
    $.ajax({
      url: "/dashy",
      type: "POST",
      data: JSON.stringify({
        url: "http://localhost:4567/dashboard.json",
        interests: []
      })
    })
    .success(function(data) {
      buildGrid(JSON.parse(data));
    })
    .fail(function() {
      errorGrid();
    })
    .done(function() {
      console.log("tick!");
    });
  }

  dash();
  setInterval(dash, 30000);
}(window, document, jQuery));
