// javascript/test/components/PipelineGroup.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var Pipeline = require("../../src/components/Pipeline");
var PipelineGroup = require("../../src/components/PipelineGroup");

describe("PipelineGroup", function () {
    describe("#render", function () {
        var pipelineOneStages = [{
            name: "DashyCompile",
            status: "Passed"
        }, {
            name: "DashyTest",
            status: "Building"
        }];
        var pipelineTwoStages = [{
            name: "FlashyCompile",
            status: "Passed"
        }, {
            name: "FlashyTest",
            status: "Failing"
        }];

        var pipelineOne = {
            name: "Dashy",
            stages: pipelineOneStages
        };
        var pipelineTwo = {
            name: "Dashy",
            stages: pipelineTwoStages
        };

        var pipelineGroup = new PipelineGroup().render([pipelineOne, pipelineTwo]);

        it("has creates a DOM representation", function () {
            expect(pipelineGroup[0]).toEqual("div");
        });

        it("has CSS class", function () {
            expect(pipelineGroup[1]).toEqual({
                class: "pipeline-group pipeline-group-2"
            });
        });

        it("has pipelines as children", function () {
            var children = pipelineGroup[2];
            var firstChild = children[0];
            var secondChild = children[1];

            expect(firstChild[0]).toEqual(Pipeline);
            expect(firstChild[1]).toEqual(pipelineOne);

            expect(secondChild[0]).toEqual(Pipeline);
            expect(secondChild[1]).toEqual(pipelineTwo);
        });
    });
});
