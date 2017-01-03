// javascript/test/components/PipelineGroupList.test.js
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

var expect = require("chai").expect;

var PipelineGroup = require("../../src/components/PipelineGroup");
var PipelineGroupLister = require("../../src/components/PipelineGroupList");

describe("PipelineGroupList", function () {
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

        describe("when group-size is less than total pipelines", function () {
            var pipelineList = new PipelineGroupLister(1)().render([pipelineOne, pipelineTwo]);

            it("has more than one pipeline group", function () {
                expect(pipelineList).to.have.length(2);
            });

            it("has pipelines as its DOM children", function () {
                expect(pipelineList[0][0]).to.equal(PipelineGroup);
                expect(pipelineList[1][0]).to.equal(PipelineGroup);
            });

            it("has pipelines DOM constructed from pipelines data", function () {
                expect(pipelineList[0][1]).to.deep.equal([pipelineOne]);
                expect(pipelineList[1][1]).to.deep.equal([pipelineTwo]);
            });
        });

        describe("when group-size equals the number of pipelines", function () {
            var pipelineList = new PipelineGroupLister(2)().render([pipelineOne, pipelineTwo]);

            it("has a single pipeline group with all pipelines", function () {
                expect(pipelineList).to.have.length(1);
            });

            it("has pipelines as its DOM children", function () {
                expect(pipelineList[0][0]).to.equal(PipelineGroup);
            });

            it("has pipelines DOM constructed from pipelines data", function () {
                expect(pipelineList[0][1]).to.deep.equal([pipelineOne, pipelineTwo]);
            });
        });

        describe("when group-size is 3", function () {
            var pipelineList = new PipelineGroupLister(3)().render([pipelineOne, pipelineTwo]);

            it("does not fill out the first pipeline group", function () {
                expect(pipelineList).to.have.length(1);
            });

            it("has pipelines as its DOM children", function () {
                expect(pipelineList[0][0]).to.equal(PipelineGroup);
            });

            it("has pipelines DOM constructed from pipelines data", function () {
                expect(pipelineList[0][1]).to.deep.equal([pipelineOne, pipelineTwo]);
            });
        });
    });
});
