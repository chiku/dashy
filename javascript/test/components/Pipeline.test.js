// javascript/test/components/Pipeline.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT


var PipelineName = require("../../src/components/PipelineName");
var StageList = require("../../src/components/StageList");
var Pipeline = require("../../src/components/Pipeline");

describe("Pipeline", function() {
    describe("#render", function() {
        var stageOne = {
            name: "Compile",
            status: "Passed"
        };
        var stageTwo = {
            name: "Test",
            status: "Building"
        };
        var pipeline = new Pipeline().render({
            name: "Dashy",
            stages: [stageOne, stageTwo]
        });

        it("creates a DOM representation", function() {
            expect(pipeline[0]).toEqual("div");
        });

        it("has CSS class", function() {
            expect(pipeline[1]).toEqual({
                class: "pipeline"
            });
        });

        it("has DOM children", function() {
            var children = pipeline[2];
            expect(children.length).toEqual(2);
        });

        it("has a list of stages as DOM child", function() {
            var children = pipeline[2];
            var stagesChild = children[0];

            expect(stagesChild[0]).toEqual(StageList);
            expect(stagesChild[1]).toEqual([stageOne, stageTwo]);
        });

        it("has pipeline name as DOM child", function() {
            var children = pipeline[2];
            var nameChild = children[1];

            expect(nameChild[0]).toEqual(PipelineName);
            expect(nameChild[1]).toEqual("Dashy");
        });
    });
});